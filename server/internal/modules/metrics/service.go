package metrics

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
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
	key := eventType + ":" + userID

	if err := s.rdbClient.Incr(ctx, key).Err(); err != nil {
		s.logger.Error("Failed to increment metric", "key", key, "error", err)
		return false
	}

	s.logger.Info("Metric incremented", "key", key)
	return true
}

func (s *metricsService) StartAggregationWorker() {
	s.logger.Info("Starting metrics aggregation worker...")

	ticker := time.NewTicker(5 * time.Minute)

	go func() {
		for range ticker.C {
			s.logger.Info("Running metrics aggregation job")
			ctx := context.Background()

			dateToProcess := time.Now().AddDate(0, 0, -1)
			if err := s.processMetricsForDate(ctx, dateToProcess); err != nil {
				s.logger.Error("Failed to process metrics", "error", err)
			}
		}
	}()
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
		profileID := parts[2]

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
		if err := s.repository.UpsertDailyMetrics(ctx, profileID, date, views, clicks); err != nil {
			s.logger.Error("Failed to upsert daily metrics to postgres", "profile_id", profileID, "error", err)
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
