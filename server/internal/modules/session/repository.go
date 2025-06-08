package session

import (
	"conecta-mare-server/internal/database/models"
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
)

type sessionsRepository struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) SessionsRepository {
	return &sessionsRepository{db: db}
}

func (r *sessionsRepository) GetAllByUserID(ctx context.Context, userID string) ([]*Session, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	dbSessions := make([]models.Session, 0)
	err := r.db.SelectContext(
		ctx,
		&dbSessions,
		"SELECT * FROM sessions WHERE user_id = $1 ORDER BY created_at DESC",
		userID,
	)
	if err != nil {
		return nil, err
	}

	result := make([]*Session, len(dbSessions))
	for i, ms := range dbSessions {
		result[i] = NewFromModel(ms)
	}

	return result, nil
}

func (r *sessionsRepository) GetActiveByUserID(ctx context.Context, userID string) (*Session, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	var sessionModel models.Session
	err := r.db.GetContext(
		ctx,
		&sessionModel,
		"SELECT * FROM sessions WHERE user_id = $1 AND active = true LIMIT 1",
		userID,
	)
	if err != nil {
		return nil, err
	}

	return NewFromModel(sessionModel), nil
}

func (r *sessionsRepository) GetByJTI(ctx context.Context, JTI string) (*Session, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	var sessionModel models.Session
	err := r.db.GetContext(
		ctx,
		&sessionModel,
		"SELECT * FROM sessions WHERE jti = $1",
		JTI,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return NewFromModel(sessionModel), nil
}

func (r *sessionsRepository) Create(ctx context.Context, session *Session) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	modelSession := session.ToModel()

	query := `
		INSERT INTO sessions (
			id,
			user_id,
			jti,
			active,
			created_at,
			updated_at,
			expires_at
		)	VALUES (
			:id,
			:user_id,
			:jti,
			:active,
			:created_at,
			:updated_at,
			:expires_at
		)
	`

	_, err := r.db.NamedExecContext(ctx, query, modelSession)
	if err != nil {
		return err
	}

	return nil
}

func (r *sessionsRepository) Update(ctx context.Context, session *Session) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	modelSession := session.ToModel()
	query := `
		UPDATE sessions
		SET
			active = :active,
			jti = :jti,
			updated_at = :updated_at
		WHERE id = :id
	`

	_, err := r.db.NamedExecContext(ctx, query, modelSession)
	if err != nil {
		return err
	}

	return nil
}

func (r sessionsRepository) DeactivateAll(ctx context.Context, userId string) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, "UPDATE sessions SET active = false WHERE user_id = $1", userId)
	if err != nil {
		return err
	}

	return nil
}
