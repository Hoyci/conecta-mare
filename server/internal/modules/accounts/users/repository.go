package users

import (
	"conecta-mare-server/internal/common"
	"conecta-mare-server/internal/database/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type usersRepository struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) UsersRepository {
	return &usersRepository{db: db}
}

func (ur *usersRepository) Register(ctx context.Context, tx *sqlx.Tx, user *User) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	modelUser := user.ToModel()

	query := `
		INSERT INTO users (
			id, email, password_hash, role, created_at, updated_at, deleted_at
		) VALUES (
			:id, :email, :password_hash, :role, :created_at, :updated_at, :deleted_at
		)`

	_, err := tx.NamedExecContext(ctx, query, modelUser)
	return err
}

func (ur *usersRepository) DeleteByID(ctx context.Context, ID string) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	_, err := ur.db.ExecContext(ctx, "DELETE FROM users WHERE id = $1", ID)
	return err
}

func (ur *usersRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	var user models.User
	err := ur.db.GetContext(ctx, &user, "SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (ur *usersRepository) GetByID(ctx context.Context, ID string) (*common.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	query := `
		SELECT 
			u.id,
			u.email,
			u.role,
			up.full_name,
			up.profile_image,
			up.job_description,
			subc."name"
		FROM users u
		INNER JOIN user_profiles up ON up.user_id = u.id
		INNER JOIN subcategories subc ON subc.id = up.subcategory_id
		WHERE u.id = $1 
	`

	var user common.User
	err := ur.db.GetContext(ctx, &user, query, ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (ur *usersRepository) GetByRole(ctx context.Context, role string) ([]*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	var users []*models.User
	err := ur.db.SelectContext(ctx, &users, "SELECT * FROM users WHERE role = $1", role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return users, nil
}

func (r *usersRepository) CountBySubcategoryIDs(ctx context.Context, subcategoryIDs []string) (map[string]int, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	if len(subcategoryIDs) == 0 {
		return make(map[string]int), nil
	}

	query, args, err := sqlx.In(`
		SELECT 
			s.id,
			count(*)
		FROM users u 
		INNER JOIN user_profiles up ON up.user_id = u.id
		INNER JOIN subcategories s ON s.id = up.subcategory_id AND s.id IN (?)
		GROUP BY s.id`, subcategoryIDs)
	if err != nil {
		return nil, err
	}

	query = r.db.Rebind(query)
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	counts := make(map[string]int)
	for rows.Next() {
		var subcategoryID string
		var count int
		if err := rows.Scan(&subcategoryID, &count); err != nil {
			return nil, err
		}
		counts[subcategoryID] = count
	}
	return counts, nil
}

func (ur *usersRepository) GetProfessionalUsers(ctx context.Context) ([]*common.GetProfessionalsResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	var professionals []*common.GetProfessionalsResponse
	err := ur.db.SelectContext(
		ctx,
		&professionals,
		`
			SELECT 
					u.id as user_id,
					up.full_name,
					up.profile_image,
					up.job_description,
					5 as rating,
					'vila do pinheiro' as location
			FROM users u
			INNER JOIN user_profiles up ON up.user_id = u.id
			LEFT JOIN categories c ON c.id = up.category_id 
			LEFT JOIN subcategories s ON s.id = up.subcategory_id 
			WHERE u."role" = 'professional' AND u.deleted_at IS NULL;
		`)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return professionals, nil
}

func (ur *usersRepository) GetProfessionalByID(ctx context.Context, ID string) (*common.GetProfessionalByIDResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	var raw common.GetProfessionalByIDRaw

	err := ur.db.GetContext(
		ctx,
		&raw,
		`
		SELECT 
				u.id as user_id,
				u.email,
				up.full_name,
				up.profile_image,
				up.job_description,
				up.phone,
				up.social_links,
				ca.name as category_name,
				sc.name as subcategory_name,
				5 as rating,
				'vila do pinheiro' as location,
				(
						SELECT json_agg(
								json_build_object(
										'name', p.name,
										'description', p.description,
										'images', COALESCE(images.images, '[]'::json)
								)
						)
						FROM projects p
						LEFT JOIN LATERAL (
								SELECT json_agg(
										json_build_object(
												'url', pi.url,
												'ordering', pi.ordering
										)
								) AS images
								FROM project_images pi
								WHERE pi.project_id = p.id
						) images ON true
						WHERE p.user_profile_id = up.id
				) AS projects,
				(
						SELECT json_agg(
								json_build_object(
										'institution', ce.institution,
										'course_name', ce.course_name,
										'start_date', TO_CHAR(ce.start_date, 'YYYY-MM-DD"T"HH24:MI:SS.MS"Z"'),
										'end_date', CASE 
												WHEN ce.end_date IS NULL THEN NULL 
												ELSE TO_CHAR(ce.end_date, 'YYYY-MM-DD"T"HH24:MI:SS.MS"Z"') 
										END
								)
						)
						FROM certifications ce
						WHERE ce.user_profile_id = up.id
				) AS certifications
		FROM users u
		INNER JOIN user_profiles up ON up.user_id = u.id
		LEFT JOIN categories ca ON ca.id = up.category_id 
		LEFT JOIN subcategories sc ON sc.id = up.subcategory_id 
		WHERE 
				u.role = 'professional' 
				AND u.id = $1
				AND u.deleted_at IS NULL;
		`,
		ID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	var projects []common.Project
	var certifications []common.Certification

	if err := raw.ProjectsJSON.Unmarshal(&projects); err != nil {
		return nil, fmt.Errorf("error parsing services json: %w", err)
	}
	if err := raw.CertificationsJSON.Unmarshal(&certifications); err != nil {
		return nil, fmt.Errorf("error parsing certifications json: %w", err)
	}

	professional := &common.GetProfessionalByIDResponse{
		UserID:          raw.UserID,
		Email:           raw.Email,
		FullName:        raw.FullName,
		ProfileImage:    raw.ProfileImage,
		JobDescription:  raw.JobDescription,
		Phone:           raw.Phone,
		SocialLinks:     raw.SocialLinks,
		Location:        raw.Location,
		Rating:          raw.Rating,
		CategoryName:    raw.CategoryName,
		SubcategoryName: raw.SubcategoryName,
		Projects:        projects,
		Certifications:  certifications,
	}

	return professional, nil
}
