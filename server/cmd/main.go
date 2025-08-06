package main

import (
	"conecta-mare-server/internal/config"
	"conecta-mare-server/internal/databases/clickhouse"
	"conecta-mare-server/internal/databases/postgres"
	"conecta-mare-server/internal/modules/accounts/categories"
	"conecta-mare-server/internal/modules/accounts/certifications"
	"conecta-mare-server/internal/modules/accounts/communities"
	"conecta-mare-server/internal/modules/accounts/locations"
	"conecta-mare-server/internal/modules/accounts/metrics"
	"conecta-mare-server/internal/modules/accounts/onboardings"
	"conecta-mare-server/internal/modules/accounts/projectimages"
	"conecta-mare-server/internal/modules/accounts/projects"
	"conecta-mare-server/internal/modules/accounts/serviceimages"
	"conecta-mare-server/internal/modules/accounts/services"
	"conecta-mare-server/internal/modules/accounts/session"
	"conecta-mare-server/internal/modules/accounts/subcategories"
	"conecta-mare-server/internal/modules/accounts/userprofiles"
	"conecta-mare-server/internal/modules/accounts/users"
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

	router := server.NewRouter()
	server := server.NewServer(cfg.Port, router)

	logger.Info("Starting postgres connection")
	pg := postgres.New(cfg.PGUsername, cfg.PGPassword, cfg.PGHost, cfg.PGPort, cfg.PGDatabase)
	defer pg.Close()

	logger.Info("Starting clickHouse connection")
	ch := clickhouse.New(cfg.CHUsername, cfg.CHPassword, cfg.CHHost, cfg.CHPort, cfg.CHDatabase)
	defer ch.Close()

	logger.Info("Starting storage connection")
	storageClient := storage.NewStorageClient(
		cfg.StorageURL,
		cfg.StorageAccessKey,
		cfg.StorageSecretKey,
		cfg.StorageBucketName,
		cfg.Environment,
	)

	logger.Info(fmt.Sprintf("Launching %s with the following settings:", cfg.AppName),
		"port", cfg.Port,
	)

	tokenProvider := jwt.NewProvider(cfg.JWTAccessKey, cfg.JWTRefreshKey)

	subcategoriesRepo := subcategories.NewRepository(pg.DB())
	categoriesRepo := categories.NewRepository(pg.DB())
	sessionsRepo := session.NewRepository(pg.DB())
	usersRepo := users.NewRepository(pg.DB())
	userProfilesRepo := userprofiles.NewRepository(pg.DB())
	certificationsRepo := certifications.NewRepository(pg.DB())
	projectsRepo := projects.NewRepository(pg.DB())
	projectImagesRepo := projectimages.NewRepository(pg.DB())
	servicesRepo := services.NewRepository(pg.DB())
	serviceImagesRepo := serviceimages.NewRepository(pg.DB())
	locationsRepo := locations.NewRepository(pg.DB())
	communitiesRepo := communities.NewRepository(pg.DB())
	metricsRepo := metrics.NewRepository(ch.DB())

	sessionsService := session.NewService(sessionsRepo, logger)
	subcategoriesService := subcategories.NewService(subcategoriesRepo, logger)
	usersService := users.NewService(
		pg.DB(),
		usersRepo,
		userProfilesRepo,
		sessionsService,
		storageClient,
		*tokenProvider,
		logger,
	)
	categoriesService := categories.NewService(categoriesRepo, subcategoriesService, usersService, logger)
	onboardingsService := onboardings.NewService(
		pg.DB(),
		usersRepo,
		userProfilesRepo,
		projectsRepo,
		projectImagesRepo,
		certificationsRepo,
		subcategoriesRepo,
		servicesRepo,
		serviceImagesRepo,
		locationsRepo,
		storageClient,
		logger,
	)
	communitiesService := communities.NewService(communitiesRepo, logger)
	metricsService := metrics.NewService(metricsRepo, logger)

	categoriesHandler := categories.NewHandler(categoriesService)
	categoriesHandler.RegisterRoutes(router)

	usersHandler := users.NewHandler(usersService, cfg.JWTAccessKey)
	usersHandler.RegisterRoutes(router)

	onboardingsHandler := onboardings.NewHandler(onboardingsService, cfg.JWTAccessKey)
	onboardingsHandler.RegisterRoutes(router)

	communitiesHandler := communities.NewHandler(communitiesService)
	communitiesHandler.RegisterRoutes(router)

	metricsHandler := metrics.NewHandler(metricsService, cfg.JWTAccessKey)
	metricsHandler.RegisterRoutes(router)

	done := make(chan bool, 1)

	go gracefulShutdown(server, done, logger)

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %s", err))
	}

	<-done
	logger.Info("Graceful shutdown complete.")
}
