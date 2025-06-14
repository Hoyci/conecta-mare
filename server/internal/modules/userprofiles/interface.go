package userprofiles

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type UserProfilesRepository interface {
	CreateTx(tx *sqlx.Tx, profile *UserProfile) error
	FindByUserID(ctx context.Context, userID string) (*UserProfile, error)
}
