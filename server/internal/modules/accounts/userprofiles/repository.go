package userprofiles

import (
	"conecta-mare-server/internal/database/models"
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) UserProfilesRepository {
	return &repository{db}
}

func (r *repository) CreateTx(tx *sqlx.Tx, profile *UserProfile) error {
	model := profile.UserProfileToModel()
	query := `
	INSERT INTO user_profiles (
		id,
		user_id,
		full_name,
		profile_image,
		job_description,
		phone,
		social_links
	) VALUES (
		:id,
		:user_id,
		:full_name,
		:profile_image,
		:job_description,
		:phone,
		:social_links
	)
	`
	_, err := tx.NamedExec(query, model)
	return err
}

func (r *repository) FindByUserID(ctx context.Context, userID string) (*UserProfile, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	var userProfile models.UserProfile
	err := r.db.GetContext(ctx, &userProfile, "SELECT * FROM user_profiles up WHERE up.user_id = $1", userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return NewFromModel(userProfile), nil
}
