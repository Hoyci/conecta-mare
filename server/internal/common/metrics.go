package common

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type DailyVisitData []DailyVisit

type (
	UserProfileView struct {
		PreviousWeekVisits int64          `db:"visits_previous_week" json:"-"`
		PreviousWeekData   DailyVisitData `db:"daily_data_previous"  json:"-"`
		CurrentWeekVisits  int64          `db:"visits_current_week"  json:"current_week_visits"`
		CurrentWeekData    DailyVisitData `db:"daily_data_current"   json:"current_week_data"`
		PercentageChange   float64        `db:"percentage_change"    json:"percentage_change"`
	}
	DailyVisit struct {
		Date   time.Time `json:"date"`
		Visits int64     `json:"visits"`
	}
)

func (d *DailyVisitData) Scan(value interface{}) error {
	rows, ok := value.([][]interface{})
	if !ok {
		return fmt.Errorf("unsupported scan, expected [][]interface{}, got %T", value)
	}

	for _, row := range rows {
		if len(row) != 2 {
			return fmt.Errorf("expected 2 columns in daily visit data, got %d", len(row))
		}

		date, ok := row[0].(time.Time)
		if !ok {
			return fmt.Errorf("expected first column to be time.Time, got %T", row[0])
		}

		visitsUint, ok := row[1].(uint64)
		if !ok {
			if visitsInt, ok := row[1].(int64); ok {
				visitsUint = uint64(visitsInt)
			} else {
				return fmt.Errorf("expected second column to be uint64 or int64, got %T", row[1])
			}
		}

		visits := int64(visitsUint)

		*d = append(*d, DailyVisit{Date: date, Visits: visits})
	}

	return nil
}

func (d DailyVisitData) Value() (driver.Value, error) {
	return nil, fmt.Errorf("valuer not implemented for DailyVisitData")
}
