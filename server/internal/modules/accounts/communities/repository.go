package communities

import (
	"conecta-mare-server/internal/database/models"
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
)

type communitiesRepository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) CommunitiesRepository {
	return &communitiesRepository{db: db}
}

func (r *communitiesRepository) GetByID(ctx context.Context, ID string) (*Community, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	var community models.Community
	err := r.db.GetContext(ctx, &community, "SELECT * FROM communities WHERE id = $1", ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return NewFromModel(community), nil
}

func (r *communitiesRepository) GetCommunities(ctx context.Context) ([]*models.Community, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	var communities []*models.Community
	err := r.db.SelectContext(ctx, &communities, "SELECT * FROM communities cm ORDER BY cm.name ASC")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return communities, nil
}
