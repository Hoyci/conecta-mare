package users

import (
	"conecta-mare-server/internal/common"
	"conecta-mare-server/internal/database/models"
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
)

type usersRepository struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) UsersRepository {
	return &usersRepository{db: db}
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

func (ur *usersRepository) GetByID(ctx context.Context, ID string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	var user models.User
	err := ur.db.GetContext(ctx, &user, "SELECT * FROM users WHERE id = $1", ID)
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

func (ur *usersRepository) Register(ctx context.Context, user *User) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	modelUser := user.ToModel()

	query := `
		INSERT INTO users (
			id,
			email,
			password_hash,
			role,
			created_at,
			updated_at,
			deleted_at
		) VALUES (
			:id,
			:email,
			:password_hash,
			:role,
			:created_at,
			:updated_at,
			:deleted_at
		)
	`

	_, err := ur.db.NamedExecContext(ctx, query, modelUser)
	return err
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
		LEFT JOIN subcategories s ON s.id = up.subcategory_id AND s.id IN (?)
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

func (ur *usersRepository) GetProfessionalUsers(ctx context.Context) ([]*common.ProfessionalResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	var professionals []*common.ProfessionalResponse
	err := ur.db.SelectContext(
		ctx,
		&professionals,
		`
			SELECT 
					u.id as user_id,
					u.email,
					u."role",
					up.full_name,
					up.profile_image,
					up.job_description,
					up.phone,
					up.social_links,
					c.name AS category_name,
					s."name" AS subcategory_name,
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
