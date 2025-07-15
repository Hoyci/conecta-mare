package metrics

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) MetricsRepository {
	return &repository{db: db}
}

func (r *repository) UpsertDailyMetrics(ctx context.Context, profileID string, date time.Time, views, clicks int) error {
	query := `
        INSERT INTO daily_metrics (user_profile_id, metric_date, profile_views, contact_clicks)
        VALUES ($1, $2, $3, $4)
        ON CONFLICT (user_profile_id, metric_date) DO UPDATE
        SET
            profile_views = daily_metrics.profile_views + EXCLUDED.profile_views,
            contact_clicks = daily_metrics.contact_clicks + EXCLUDED.contact_clicks;
    `
	_, err := r.db.ExecContext(ctx, query, profileID, date, views, clicks)
	return err
}
