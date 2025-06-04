package users

import (
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
			name,
			email,
			password_hash,
			avatar_url,
			role,
			subcategory_id,
			created_at,
			updated_at,
			deleted_at
		) VALUES (
			:id,
			:name,
			:email,
			:password_hash,
			:avatar_url,
			:role,
			:subcategory_id,
			:created_at,
			:updated_at,
			:deleted_at
		)
	`

	_, err := ur.db.NamedExecContext(ctx, query, modelUser)
	return err
}
