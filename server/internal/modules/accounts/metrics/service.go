package metrics

import (
	"conecta-mare-server/pkg/utils"
	"context"
	"fmt"
	"log/slog"
	"time"
)

func NewService(
	repository MetricsRepository,
	logger *slog.Logger,
) MetricsService {
	return &metricsService{
		repository: repository,
		logger:     logger,
	}
}

func (s *metricsService) GetUserProfileViews(ctx context.Context, userID, startDateStr, endDateStr string) (any, error) {
	s.logger.InfoContext(ctx, "attempting to get user profile views")

	var startDate, endDate time.Time
	var err error

	now := time.Now().UTC()
	loc := time.UTC

	if startDateStr == "" {
		startDate = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc).AddDate(0, 0, -6)
	} else {
		startDate, err = utils.ParseDateTimeFlexible(startDateStr)
		if err != nil {
			return nil, fmt.Errorf("Invalid startDate, accepted formats: RFC3339, '2006-01-02 15:04:05' or '2006-01-02'")
		}
	}
	if endDateStr == "" {
		endDate = time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, loc)
	} else {
		endDate, err = utils.ParseDateTimeFlexible(endDateStr)
		if err != nil {
			return nil, fmt.Errorf("Invalid endDate, accepted formats: RFC3339, '2006-01-02 15:04:05' or '2006-01-02'")
		}
		if len(endDateStr) == len("2006-01-02") {
			endDate = time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 23, 59, 59, 0, loc)
		}
	}

	userProfileViews, err := s.repository.UserProfileViews(ctx, userID, startDate, endDate)
	if err != nil {
		s.logger.ErrorContext(ctx, "error while attempting to get user profile views", "err", err)
		return nil, err
	}

	fmt.Println(userProfileViews)

	return userProfileViews, nil
}
