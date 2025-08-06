package categories

import (
	"conecta-mare-server/internal/common"
	"conecta-mare-server/internal/databases/postgres/models"
	"conecta-mare-server/internal/modules/accounts/subcategories"
	"conecta-mare-server/internal/modules/accounts/users"
	"conecta-mare-server/pkg/exceptions"
	"context"
	"log/slog"
)

type (
	CategoriesRepository interface {
		GetByID(ctx context.Context, ID string) (*Category, error)
		GetCategories(ctx context.Context) ([]*models.Category, error)
		GetCategoriesWithSubcats(ctx context.Context) ([]*common.CategoryWithSubcats, error)
	}
	CategoriesService interface {
		GetCategories(ctx context.Context, includeSubcats bool) (any, *exceptions.ApiError[string])
	}

	categoryService struct {
		repository    CategoriesRepository
		subcatService subcategories.SubcategoriesService
		usersService  users.UsersService
		logger        *slog.Logger
	}
	categoryHandler struct {
		categoriesService CategoriesService
	}
)
