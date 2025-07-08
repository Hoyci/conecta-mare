package userprofiles

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type UserProfilesRepository interface {
	CreateInitialProfileTx(ctx context.Context, tx *sqlx.Tx, userProfile *UserProfile) error
	FindByUserID(ctx context.Context, userID string) (*UserProfile, error)
	UpdateTx(ctx context.Context, tx *sqlx.Tx, userProfile *UserProfile) error
}
