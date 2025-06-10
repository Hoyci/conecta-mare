package subcategories

import (
	"conecta-mare-server/pkg/exceptions"
	"context"
	"log/slog"
	"net/http"
)

func NewService(
	repository SubcategoriesRepository,
	logger *slog.Logger,
) SubcategoriesService {
	return &subcategoryService{
		repository: repository,
		logger:     logger,
	}
}

func (s *subcategoryService) GetByCategoriesID(ctx context.Context, categoriesID []string) ([]*Subcategory, error) {
	s.logger.InfoContext(ctx, "attempting to get subcategories by categories_id", "categories_id", categoriesID)

	subcategories, err := s.repository.GetByCategoriesID(ctx, categoriesID)
	if err != nil {
		s.logger.ErrorContext(ctx, "error while attempting to get subcategories by categories_id", "err", err)
		return nil, exceptions.MakeGenericApiError()
	}

	if subcategories == nil {
		s.logger.InfoContext(ctx, "any subcategory found")
		return nil, exceptions.MakeApiErrorWithStatus(http.StatusNotFound, exceptions.ErrSubcategoriesNotFound)
	}

	result := []*Subcategory{}
	for _, subcategory := range subcategories {
		result = append(result, NewFromModel(*subcategory))
	}

	return result, nil
}
