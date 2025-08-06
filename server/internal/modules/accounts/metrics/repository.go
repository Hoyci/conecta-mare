package metrics

import (
	"conecta-mare-server/internal/common"
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

func NewRepository(db *sqlx.DB) MetricsRepository {
	return &metricsRepository{db: db}
}

func (r *metricsRepository) UserProfileViews(ctx context.Context, userID string, startDate, endDate time.Time) (any, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	var userProfileView common.UserProfileView

	query := `
		WITH
				toDate($2) AS start_date_current,
				toDate($3) AS end_date_current,
				$1 AS target_professional_id,
				start_date_current - (end_date_current - start_date_current + 1) AS start_date_previous,
				start_date_current - 1 AS end_date_previous
		SELECT
				sumIf(daily_visits.total_visits, visit_date BETWEEN start_date_previous AND end_date_previous) AS visits_previous_week,
				sumIf(daily_visits.total_visits, visit_date BETWEEN start_date_current AND end_date_current) AS visits_current_week,
				if(
						visits_previous_week = 0,
						if(visits_current_week > 0, 100.0, 0.0),
						((toFloat64(visits_current_week) - visits_previous_week) / visits_previous_week) * 100
				) AS percentage_change,


				arraySort(
						x -> x.1,
						groupArrayIf(
								(visit_date, total_visits),
								visit_date BETWEEN start_date_previous AND end_date_previous
						)
				) AS daily_data_previous,

				arraySort(
						x -> x.1,
						groupArrayIf(
								(visit_date, total_visits),
								visit_date BETWEEN start_date_current AND end_date_current
						)
				) AS daily_data_current
		FROM
		(
				SELECT
						toDate(d.dates) AS visit_date,
						count(pv.professional_id) AS total_visits
				FROM
				(
						SELECT arrayJoin(range(toUInt32(start_date_previous), toUInt32(end_date_current) + 1)) AS dates
				) AS d
				LEFT JOIN profile_visited AS pv
						ON toDate(pv."timestamp") = toDate(d.dates)
						AND pv.professional_id = target_professional_id
				GROUP BY
						visit_date
		) AS daily_visits;
	`

	err := r.db.GetContext(ctx, &userProfileView, query, userID, startDate, endDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return &common.UserProfileView{}, nil
		}
		return nil, err
	}

	return &userProfileView, nil
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
