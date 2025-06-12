package onboardings

import (
	"conecta-mare-server/internal/common"
	"conecta-mare-server/internal/database/models"
	"conecta-mare-server/pkg/storage"
	"context"
	"log/slog"
	"net/http"
)

type (
	OnboardingsRepository interface {
		CreateUserProfile(
			ctx context.Context,
			profile *UserProfile,
			certifications []*Certification,
			services []*Service,
			serviceImages []*ServiceImage,
		) error
		GetUserProfileByUserID(ctx context.Context, userID string) (*models.UserProfile, error)
	}
	OnboardingsService interface {
		CompleteOnboarding(ctx context.Context, req *common.OnboardingRequest, r *http.Request) error
	}
	onboardingService struct {
		repository OnboardingsRepository
		storage    *storage.StorageClient
		logger     *slog.Logger
	}
	onboardingHandler struct {
		service   OnboardingsService
		accessKey string
	}
)
