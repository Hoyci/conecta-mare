package main

import (
	"conecta-mare-server/internal/config"
	"conecta-mare-server/internal/database"
	"conecta-mare-server/internal/modules/session"
	"conecta-mare-server/internal/modules/users"
	"conecta-mare-server/internal/server"
	"conecta-mare-server/pkg/jwt"
	"conecta-mare-server/pkg/storage"
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	slogformatter "github.com/samber/slog-formatter"
)

func gracefulShutdown(apiServer *http.Server, done chan bool, logger *slog.Logger) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	logger.InfoContext(ctx, "shutting down gracefully, press Ctrl+C again to force")
	stop()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := apiServer.Shutdown(ctx); err != nil {
		logger.InfoContext(ctx, fmt.Sprintf("Server forced to shutdown with error: %v", err))
	}

	logger.InfoContext(ctx, "Server exiting")

	done <- true
}

func main() {
	cfg := config.GetConfig()

	logLevel := slog.LevelInfo
	if cfg.Environment == "development" {
		logLevel = slog.LevelDebug
	}

	var handler slog.Handler
	if cfg.Environment == "production" {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel, AddSource: true})
	} else {
		handler = slogformatter.NewFormatterHandler(
			slogformatter.TimezoneConverter(time.UTC),
			slogformatter.TimeFormatter(time.RFC3339, nil),
		)(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel, AddSource: true}),
		)
	}

	logger := slog.New(handler)

	logger = logger.With("env", cfg.Environment)

	logger.Info(fmt.Sprintf("Launching %s with the following settings:", cfg.AppName),
		"port", cfg.Port,
	)

	router := server.NewRouter()
	server := server.NewServer(cfg.Port, router)

	db := database.New(cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBDatabase)
	defer db.Close()

	storageClient := storage.NewStorageClient(
		cfg.StorageURL,
		cfg.StorageAccessKey,
		cfg.StorageSecretKey,
		cfg.StorageBucketName,
	)

	tokenProvider := jwt.NewProvider(cfg.JWTAccessKey, cfg.JWTRefreshKey)

	sessionsRepo := session.NewRepo(db.DB())
	usersRepo := users.NewRepo(db.DB())

	sessionsService := session.NewService(sessionsRepo, logger)

	usersService := users.NewService(usersRepo, sessionsService, storageClient, *tokenProvider, logger)

	usersHandler := users.NewHandler(usersService)
	usersHandler.RegisterRoutes(router)

	done := make(chan bool, 1)

	go gracefulShutdown(server, done, logger)

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %s", err))
	}

	<-done
	logger.Info("Graceful shutdown complete.")
}
