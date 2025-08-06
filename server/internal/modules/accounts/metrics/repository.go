package metrics

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

func NewRepository(db *sqlx.DB) MetricsRepository {
	return &metricsRepository{db: db}
}

func (r *metricsRepository) UserProfileViews(ctx context.Context, userID string) (any, error) {
	// ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	// defer cancel()

	fmt.Println("not implemented")
	return "", nil
}

func (r *metricsRepository) UserTopPerformingServices(ctx context.Context, userID string) (any, error) {
	// ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	// defer cancel()

	fmt.Println("not implemented")
	return "", nil
}

func (r *metricsRepository) UserProfileViewsComparisonBySubcategory(
	ctx context.Context,
	userID, subcategoryID string,
	startDate, endDate time.Time,
) (any, error) {
	// ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	// defer cancel()

	fmt.Println("not implemented")
	return "", nil
}

func (r *metricsRepository) UserProfileViewsComparisonByCategory(
	ctx context.Context,
	userID, categoryID string,
	startDate, endDate time.Time,
) (any, error) {
	// ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	// defer cancel()

	fmt.Println("not implemented")
	return "", nil
}
