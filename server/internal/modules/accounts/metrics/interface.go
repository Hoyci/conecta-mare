package metrics

import (
	"context"
	"log/slog"
	"time"

	"github.com/jmoiron/sqlx"
)

type (
	MetricsRepository interface {
		UserProfileViews(ctx context.Context, userID string, startDate, endDate time.Time) (any, error)
		UserTopPerformingServices(ctx context.Context, userID string) (any, error)
		UserProfileViewsComparisonBySubcategory(ctx context.Context, userID, subcategoryID string, startDate, endDate time.Time) (any, error)
		UserProfileViewsComparisonByCategory(ctx context.Context, userID, categoryID string, startDate, endDate time.Time) (any, error)
	}

	MetricsService interface {
		GetUserProfileViews(ctx context.Context, userID, startDate, endDate string) (any, error)
	}

	metricsRepository struct {
		db *sqlx.DB
	}

	metricsService struct {
		repository MetricsRepository
		logger     *slog.Logger
	}
	metricsHandler struct {
		metricsService MetricsService
		accessKey      string
	}
)
