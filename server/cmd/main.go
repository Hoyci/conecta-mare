package main

import (
	"conecta-mare-server/internal/config"
	"conecta-mare-server/internal/database"
	"conecta-mare-server/internal/modules/accounts/categories"
	"conecta-mare-server/internal/modules/accounts/certifications"
	"conecta-mare-server/internal/modules/accounts/locations"
	"conecta-mare-server/internal/modules/accounts/onboardings"
	"conecta-mare-server/internal/modules/accounts/projectimages"
	"conecta-mare-server/internal/modules/accounts/projects"
	"conecta-mare-server/internal/modules/accounts/serviceimages"
	"conecta-mare-server/internal/modules/accounts/services"
	"conecta-mare-server/internal/modules/accounts/session"
	"conecta-mare-server/internal/modules/accounts/subcategories"
	"conecta-mare-server/internal/modules/accounts/userprofiles"
	"conecta-mare-server/internal/modules/accounts/users"
	"conecta-mare-server/internal/modules/metrics"
	"conecta-mare-server/internal/redis"
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

	logger.Info("Starting database connection")
	db := database.New(cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBDatabase)
	defer db.Close()

	logger.Info("Starting Redis connection")
	rdbClient := redis.NewClient(cfg.RedisHost, cfg.RedisPort)
	defer rdbClient.Close()

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

	subcategoriesRepo := subcategories.NewRepo(db.DB())
	categoriesRepo := categories.NewRepo(db.DB())
	sessionsRepo := session.NewRepo(db.DB())
	usersRepo := users.NewRepo(db.DB())
	userProfilesRepo := userprofiles.NewRepository(db.DB())
	certificationsRepo := certifications.NewRepository(db.DB())
	projectsRepo := projects.NewRepository(db.DB())
	projectImagesRepo := projectimages.NewRepository(db.DB())
	servicesRepo := services.NewRepository(db.DB())
	serviceImagesRepo := serviceimages.NewRepository(db.DB())
	locationsRepo := locations.NewRepository(db.DB())
	metricsRepo := metrics.NewRepository(db.DB())

	sessionsService := session.NewService(sessionsRepo, logger)
	subcategoriesService := subcategories.NewService(subcategoriesRepo, logger)
	usersService := users.NewService(
		db.DB(),
		usersRepo,
		userProfilesRepo,
		sessionsService,
		storageClient,
		*tokenProvider,
		logger,
	)
	categoriesService := categories.NewService(categoriesRepo, subcategoriesService, usersService, logger)
	onboardingsService := onboardings.NewService(
		db.DB(),
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

	metricsService := metrics.NewService(metricsRepo, rdbClient, logger)

	categoriesHandler := categories.NewHandler(categoriesService)
	categoriesHandler.RegisterRoutes(router)

	usersHandler := users.NewHandler(usersService, cfg.JWTAccessKey)
	usersHandler.RegisterRoutes(router)

	onboardingsHandler := onboardings.NewHandler(onboardingsService, cfg.JWTAccessKey)
	onboardingsHandler.RegisterRoutes(router)

	metricsHandler := metrics.NewHandler(
		cfg.JWTAccessKey,
		metricsService,
		logger,
	)
	metricsHandler.RegisterRoutes(router)

	done := make(chan bool, 1)

	go metricsService.StartAggregationWorker()
	go gracefulShutdown(server, done, logger)

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %s", err))
	}

	<-done
	logger.Info("Graceful shutdown complete.")
}
