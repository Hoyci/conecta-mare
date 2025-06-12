package onboardings

import (
	"conecta-mare-server/internal/database/models"
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
)

type onboardingRepository struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) OnboardingsRepository {
	return &onboardingRepository{db: db}
}

func (o *onboardingRepository) CreateUserProfile(
	ctx context.Context,
	profile *UserProfile,
	certifications []*Certification,
	services []*Service,
	serviceImages []*ServiceImage,
) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	tx, err := o.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	profileModel := profile.UserProfileToModel()
	query := `
		INSERT INTO user_profiles (
			id, user_id, full_name, profile_image, job_description, phone, social_links, created_at
		) VALUES (
			:id, :user_id, :full_name, :profile_image, :job_description, :phone, :social_links, :created_at
		)
	`
	if _, err := tx.NamedExecContext(ctx, query, profileModel); err != nil {
		return err
	}

	for _, cert := range certifications {
		certModel := cert.ToModel()
		query := `
			INSERT INTO certifications (
				id, user_profile_id, institution, course_name, start_date, end_date, created_at
			) VALUES (
				:id, :user_profile_id, :institution, :course_name, :start_date, :end_date, :created_at
			)
		`
		if _, err := tx.NamedExecContext(ctx, query, certModel); err != nil {
			return err
		}
	}

	for _, svc := range services {
		svcModel := svc.ToModel()
		query := `
			INSERT INTO services (
				id, user_profile_id, name, description, created_at
			) VALUES (
				:id, :user_profile_id, :name, :description, :created_at
			)
		`
		if _, err := tx.NamedExecContext(ctx, query, svcModel); err != nil {
			return err
		}
	}

	for _, img := range serviceImages {
		imgModel := img.ToModel()
		query := `
			INSERT INTO service_images (
				id, service_id, url, ordering, created_at
			) VALUES (
				:id, :service_id, :url, :ordering, :created_at
			)
		`
		if _, err := tx.NamedExecContext(ctx, query, imgModel); err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (o *onboardingRepository) GetUserProfileByUserID(ctx context.Context, userID string) (*models.UserProfile, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	var userProfile models.UserProfile
	err := o.db.GetContext(ctx, &userProfile, "SELECT * FROM user_profiles up WHERE up.user_id = $1", userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &userProfile, nil
}
