package categories

import (
	"conecta-mare-server/internal/common"
	"conecta-mare-server/internal/modules/subcategories"
	"conecta-mare-server/internal/modules/users"
	"conecta-mare-server/pkg/exceptions"
	"context"
	"log/slog"
)

func NewService(
	repository CategoriesRepository,
	subcatService subcategories.SubcategoriesService,
	usersService users.UsersService,
	logger *slog.Logger,
) CategoriesService {
	return &categoryService{
		repository:    repository,
		subcatService: subcatService,
		usersService:  usersService,
		logger:        logger,
	}
}

func (s *categoryService) GetCategories(ctx context.Context, includeSubcats bool) (any, *exceptions.ApiError[string]) {
	s.logger.InfoContext(ctx, "attempting to get categories with user count")

	if includeSubcats {
		records, err := s.repository.GetCategoriesWithSubcats(ctx)
		if err != nil {
			s.logger.ErrorContext(ctx, "error while attempting to get categories with subcategories", "err", err)
			return nil, exceptions.MakeGenericApiError()
		}
		return records, nil
	}

	categories, err := s.repository.GetCategories(ctx)
	if err != nil {
		s.logger.ErrorContext(ctx, "error while attempting to get categories", "err", err)
		return nil, exceptions.MakeGenericApiError()
	}
	if len(categories) == 0 {
		return []common.CategoryWithUserCount{}, nil
	}

	categoryIDs := make([]string, len(categories))
	for i, cat := range categories {
		categoryIDs[i] = cat.ID
	}

	subcats, err := s.subcatService.GetByCategoriesID(ctx, categoryIDs)
	if err != nil {
		s.logger.ErrorContext(ctx, "failed to get subcategories by category ids", "err", err)
		return nil, exceptions.MakeGenericApiError()
	}

	subcatsByCatID := make(map[string][]string)
	allSubcatIDs := make([]string, 0, len(subcats))
	for _, sub := range subcats {
		subID := sub.ID()
		allSubcatIDs = append(allSubcatIDs, subID)
		catID := sub.CategoryID()
		subcatsByCatID[catID] = append(subcatsByCatID[catID], subID)
	}

	userCountsBySubcatID, err := s.usersService.CountUsersBySubcategoryIDs(ctx, allSubcatIDs)
	if err != nil {
		s.logger.ErrorContext(ctx, "failed to count users by subcategory ids", "err", err)
		return nil, exceptions.MakeGenericApiError()
	}

	response := make([]common.CategoryWithUserCount, len(categories))
	for i, cat := range categories {
		totalUsers := 0
		if subcatIDs, ok := subcatsByCatID[cat.ID]; ok {
			for _, subID := range subcatIDs {
				totalUsers += userCountsBySubcatID[subID]
			}
		}
		response[i] = common.CategoryWithUserCount{
			Category: common.Category{
				ID:   cat.ID,
				Name: cat.Name,
				Icon: cat.Icon,
			},
			UserCount: totalUsers,
		}
	}

	return response, nil
}
