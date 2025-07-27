package communities

import (
	"conecta-mare-server/internal/common"
	"conecta-mare-server/pkg/exceptions"
	"context"
	"log/slog"
)

func NewService(
	repository CommunitiesRepository,
	logger *slog.Logger,
) CommunitiesService {
	return &communitiesService{
		repository: repository,
		logger:     logger,
	}
}

func (s *communitiesService) GetCommunities(ctx context.Context) ([]common.Community, *exceptions.ApiError[string]) {
	s.logger.InfoContext(ctx, "attempting to get communities")

	communities, err := s.repository.GetCommunities(ctx)
	if err != nil {
		s.logger.ErrorContext(ctx, "error while attempting to get communities", "err", err)
		return nil, exceptions.MakeGenericApiError()
	}

	response := make([]common.Community, len(communities))
	for i, com := range communities {
		response[i] = common.Community{
			ID:      com.ID,
			Name:    com.Name,
			CensoID: com.CensoID,
		}
	}

	return response, nil
}
