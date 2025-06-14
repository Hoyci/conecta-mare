package onboardings

import (
	"conecta-mare-server/internal/common"
	"conecta-mare-server/internal/modules/certifications"
	"conecta-mare-server/internal/modules/serviceimages"
	"conecta-mare-server/internal/modules/services"
	"conecta-mare-server/internal/modules/userprofiles"
	"conecta-mare-server/internal/modules/users"
	"conecta-mare-server/pkg/storage"
	"context"
	"log/slog"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type (
	OnboardingsService interface {
		MakeOnboarding(ctx context.Context, r *http.Request, req *common.OnboardingRequest) error
	}
	onboardingsService struct {
		db                       *sqlx.DB
		usersRepository          users.UsersRepository
		userProfilesRepository   userprofiles.UserProfilesRepository
		servicesRepository       services.ServicesRepository
		serviceImagesRepository  serviceimages.ServiceImagesRepository
		certificationsRepository certifications.CertificationsRepository
		storage                  *storage.StorageClient
		logger                   *slog.Logger
	}
	onboardingsHandler struct {
		onboardingsService OnboardingsService
		accessKey          string
	}
)
