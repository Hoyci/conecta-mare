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

func (r *repository) CreateInitialProfileTx(ctx context.Context, tx *sqlx.Tx, userProfile *UserProfile) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	modelUserProfile := userProfile.ToModel()

	query := `
		INSERT INTO user_profiles (
 			id, 
			user_id, 
			full_name, 
			subcategory_id,
			profile_image,
			job_description,
		 	phone,
			social_links,
			created_at,
			updated_at
		) VALUES (
			:id, 
			:user_id, 
			:full_name,
			:subcategory_id,
		 	:profile_image,
			:job_description,
			:phone,
			:social_links,
			:created_at,
			:updated_at
		)`

	_, err := tx.NamedExecContext(ctx, query, modelUserProfile)
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

func (r *repository) UpdateTx(ctx context.Context, tx *sqlx.Tx, userProfile *UserProfile) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	modelUserProfile := userProfile.ToModel()

	query := `
		UPDATE user_profiles SET
			full_name = :full_name,
			subcategory_id = :subcategory_id,
			profile_image = :profile_image,
			job_description = :job_description,
			phone = :phone,
			social_links = :social_links,
			updated_at = :updated_at
		WHERE user_id = :user_id
	`

	_, err := tx.NamedExecContext(ctx, query, modelUserProfile)
	return err
}
