package metrics

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type (
	MetricsRepository interface {
		UserProfileViews(ctx context.Context, userID string) (any, error)
		UserTopPerformingServices(ctx context.Context, userID string) (any, error)
		UserProfileViewsComparisonBySubcategory(ctx context.Context, userID, subcategoryID string, startDate, endDate time.Time) (any, error)
		UserProfileViewsComparisonCategory(ctx context.Context, userID, categoryID string, startDate, endDate time.Time) (any, error)
	}

	// CommunitiesService interface {
	// 	GetCommunities(ctx context.Context) ([]common.Community, *exceptions.ApiError[string])
	// }

	metricsRepository struct {
		db *sqlx.DB
	}

	// communitiesService struct {
	// 	repository CommunitiesRepository
	// 	logger     *slog.Logger
	// }
	// communitiesHandler struct {
	// 	categoriesService CommunitiesService
	// }
)
