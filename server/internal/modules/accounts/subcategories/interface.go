package subcategories

import (
	"conecta-mare-server/internal/database/models"
	"context"
	"log/slog"
)

type (
	SubcategoriesRepository interface {
		GetByID(ctx context.Context, ID string) (*Subcategory, error)
		GetByCategoriesID(ctx context.Context, categoriesID []string) ([]*models.Subcategory, error)
	}
	SubcategoriesService interface {
		GetByCategoriesID(ctx context.Context, categoriesID []string) ([]*Subcategory, error)
	}
	subcategoryService struct {
		repository SubcategoriesRepository
		logger     *slog.Logger
	}
)
