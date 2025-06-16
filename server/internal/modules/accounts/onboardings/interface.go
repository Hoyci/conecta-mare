package onboardings

import (
	"conecta-mare-server/internal/common"
	"conecta-mare-server/internal/modules/accounts/categories"
	"conecta-mare-server/internal/modules/accounts/certifications"
	"conecta-mare-server/internal/modules/accounts/serviceimages"
	"conecta-mare-server/internal/modules/accounts/services"
	"conecta-mare-server/internal/modules/accounts/subcategories"
	"conecta-mare-server/internal/modules/accounts/userprofiles"
	"conecta-mare-server/internal/modules/accounts/users"
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
		categoriesRepository     categories.CategoriesRepository
		subcategoriesRepository  subcategories.SubcategoriesRepository
		storage                  *storage.StorageClient
		logger                   *slog.Logger
	}
	onboardingsHandler struct {
		onboardingsService OnboardingsService
		accessKey          string
	}
)
