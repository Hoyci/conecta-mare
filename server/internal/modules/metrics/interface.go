package metrics

import (
	"context"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
)

type (
	MetricsRepository interface {
		UpsertDailyMetrics(ctx context.Context, profileID string, date time.Time, views, clicks int) error
	}
	MetricsService interface {
		TrackEvent(ctx context.Context, eventType, profileID string) bool
		StartAggregationWorker()
	}
	metricsService struct {
		repository MetricsRepository
		rdbClient  *redis.Client
		logger     *slog.Logger
	}
	metricsHandler struct {
		accessKey      string
		metricsService MetricsService
		logger         *slog.Logger
	}
)
