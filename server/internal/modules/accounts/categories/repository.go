package categories

import (
	"conecta-mare-server/internal/common"
	"conecta-mare-server/internal/databases/postgres/models"
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
)

type categoriesRepository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) CategoriesRepository {
	return &categoriesRepository{db: db}
}

func (r *categoriesRepository) GetByID(ctx context.Context, ID string) (*Category, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	var category models.Category
	err := r.db.GetContext(ctx, &category, "SELECT * FROM categories WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return NewFromModel(category), nil
}

func (r *categoriesRepository) GetCategories(ctx context.Context) ([]*models.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	var categories []*models.Category
	err := r.db.SelectContext(ctx, &categories, "SELECT * FROM categories c ORDER BY c.name DESC")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return categories, nil
}

func (r *categoriesRepository) GetCategoriesWithSubcats(ctx context.Context) ([]*common.CategoryWithSubcats, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	query := `
		SELECT 
			c.id AS cat_id,
			c.name AS cat_name,
			c.icon AS cat_icon,
			s.id AS sub_id,
			s.name AS sub_name
    FROM categories c
    LEFT JOIN subcategories s 
			ON s.category_id = c.id 
			AND s.deleted_at IS NULL
    WHERE c.deleted_at IS NULL
    ORDER BY c.id, s.name
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	categoriesMap := make(map[string]*common.CategoryWithSubcats)
	var result []*common.CategoryWithSubcats

	for rows.Next() {
		var (
			catID   string
			catName string
			catIcon string
			subID   sql.NullString
			subName sql.NullString
		)

		if err := rows.Scan(
			&catID,
			&catName,
			&catIcon,
			&subID,
			&subName,
		); err != nil {
			return nil, err
		}

		if _, exists := categoriesMap[catID]; !exists {
			category := &common.CategoryWithSubcats{
				Category: common.Category{
					ID:   catID,
					Name: catName,
					Icon: catIcon,
				},
				Subcategories: []common.Subcategory{},
			}
			categoriesMap[catID] = category
			result = append(result, category)
		}

		if subID.Valid && subName.Valid {
			subcategory := common.Subcategory{
				ID:   subID.String,
				Name: subName.String,
			}
			categoriesMap[catID].Subcategories = append(
				categoriesMap[catID].Subcategories,
				subcategory,
			)
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
