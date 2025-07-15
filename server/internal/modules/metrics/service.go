package metrics

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/robfig/cron/v3"
)

func NewService(repository MetricsRepository, rdbClient *redis.Client, logger *slog.Logger) MetricsService {
	return &metricsService{
		repository: repository,
		rdbClient:  rdbClient,
		logger:     logger,
	}
}

func (s *metricsService) TrackEvent(ctx context.Context, eventType string, userID string) bool {
	s.logger.InfoContext(ctx, fmt.Sprintf("creating metric of type %s for user %s", eventType, userID))
	dateStr := time.Now().Format("2006-01-02")
	key := fmt.Sprintf("%s:%s:%s", eventType, userID, dateStr)

	if err := s.rdbClient.Incr(ctx, key).Err(); err != nil {
		s.logger.Error("Failed to increment metric", "key", key, "error", err)
		return false
	}
	s.rdbClient.Expire(ctx, key, 48*time.Hour)

	s.logger.Info("Metric incremented", "key", key)
	return true
}

func (s *metricsService) StartAggregationWorker() {
	s.logger.Info("Starting metrics aggregation worker...")

	c := cron.New(cron.WithLocation(time.FixedZone("BRT", -3*60*60)))

	_, err := c.AddFunc("51 19 * * *", func() {
		s.logger.Info("Running daily metrics aggregation job")
		ctx := context.Background()
		dateToProcess := time.Now()

		if err := s.processMetricsForDate(ctx, dateToProcess); err != nil {
			s.logger.Error("Failed to process metrics", "error", err)
		}
	})
	if err != nil {
		s.logger.Error("Failed to schedule aggregation job", "error", err)
		return
	}

	c.Start()
}

func (s *metricsService) processMetricsForDate(ctx context.Context, date time.Time) error {
	dateStr := date.Format("2006-01-02")
	s.logger.Info("Processing metrics for date", "date", dateStr)

	if err := s.aggregateMetric(ctx, "views:profile", dateStr); err != nil {
		return err
	}

	if err := s.aggregateMetric(ctx, "clicks:contact", dateStr); err != nil {
		return err
	}

	return nil
}

func (s *metricsService) aggregateMetric(ctx context.Context, metricPrefix, dateStr string) error {
	pattern := metricPrefix + ":*:" + dateStr
	s.logger.Debug("Scanning for keys with pattern", "pattern", pattern)

	var cursor uint64
	var keys []string

	iter := s.rdbClient.Scan(ctx, cursor, pattern, 0).Iterator()
	for iter.Next(ctx) {
		keys = append(keys, iter.Val())
	}
	if err := iter.Err(); err != nil {
		s.logger.Error("Error scanning redis keys", "pattern", pattern, "error", err)
		return err
	}

	s.logger.Info("Found keys to process", "metric", metricPrefix, "count", len(keys))

	for _, key := range keys {
		parts := strings.Split(key, ":")
		if len(parts) != 4 {
			s.logger.Warn("Invalid key format found, skipping", "key", key)
			continue
		}
		userID := parts[2]

		countStr, err := s.rdbClient.Get(ctx, key).Result()
		if err != nil {
			s.logger.Error("Failed to get key value from redis", "key", key, "error", err)
			continue
		}
		count, _ := strconv.Atoi(countStr)

		var views, clicks int
		if metricPrefix == "views:profile" {
			views = count
		} else if metricPrefix == "clicks:contact" {
			clicks = count
		}

		date, _ := time.Parse("2006-01-02", dateStr)
		if err := s.repository.UpsertDailyMetrics(ctx, userID, date, views, clicks); err != nil {
			s.logger.Error("Failed to upsert daily metrics to postgres", "user_id", userID, "error", err)
			continue
		}

		s.logger.Debug("Successfully upserted metric", "key", key, "count", count)
	}

	if len(keys) > 0 {
		if err := s.rdbClient.Del(ctx, keys...).Err(); err != nil {
			s.logger.Error("Failed to delete processed keys from redis", "error", err)
		}
		s.logger.Info("Cleaned up processed keys from Redis", "metric", metricPrefix, "count", len(keys))
	}

	return nil
}
