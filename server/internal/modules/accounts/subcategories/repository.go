package subcategories

import (
	"conecta-mare-server/internal/databases/postgres/models"
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
)

type subcategoriesRepository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) SubcategoriesRepository {
	return &subcategoriesRepository{db: db}
}

func (r *subcategoriesRepository) GetByID(ctx context.Context, ID string) (*Subcategory, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	var subcategory models.Subcategory
	err := r.db.GetContext(ctx, &subcategory, "SELECT * FROM subcategories WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return NewFromModel(subcategory), nil
}

func (r *subcategoriesRepository) GetByCategoriesID(ctx context.Context, categoriesID []string) ([]*models.Subcategory, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	query, args, err := sqlx.In("SELECT * FROM subcategories s WHERE s.category_id IN (?)", categoriesID)
	if err != nil {
		return nil, err
	}

	query = r.db.Rebind(query)

	var subcategories []*models.Subcategory
	err = r.db.SelectContext(ctx, &subcategories, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return subcategories, nil
}
