package communities

import (
	"conecta-mare-server/internal/common"
	"conecta-mare-server/internal/databases/postgres/models"
	"conecta-mare-server/pkg/exceptions"
	"context"
	"log/slog"
)

type (
	CommunitiesRepository interface {
		GetByID(ctx context.Context, ID string) (*Community, error)
		GetCommunities(ctx context.Context) ([]*models.Community, error)
	}
	CommunitiesService interface {
		GetCommunities(ctx context.Context) ([]common.Community, *exceptions.ApiError[string])
	}

	communitiesService struct {
		repository CommunitiesRepository
		logger     *slog.Logger
	}
	communitiesHandler struct {
		communitiesService CommunitiesService
	}
)
