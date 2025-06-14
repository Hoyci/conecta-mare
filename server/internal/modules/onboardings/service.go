package onboardings

import (
	"conecta-mare-server/internal/common"
	"conecta-mare-server/internal/modules/certifications"
	"conecta-mare-server/internal/modules/serviceimages"
	"conecta-mare-server/internal/modules/services"
	"conecta-mare-server/internal/modules/userprofiles"
	"conecta-mare-server/internal/modules/users"
	"conecta-mare-server/pkg/exceptions"
	"conecta-mare-server/pkg/storage"
	"conecta-mare-server/pkg/uid"
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/jmoiron/sqlx"
)

func NewService(
	db *sqlx.DB,
	usersRepository users.UsersRepository,
	userProfilesRepository userprofiles.UserProfilesRepository,
	servicesRepository services.ServicesRepository,
	serviceImagesRepository serviceimages.ServiceImagesRepository,
	certificationsRepository certifications.CertificationsRepository,
	storage *storage.StorageClient,
	logger *slog.Logger,
) OnboardingsService {
	return &onboardingsService{
		db:                       db,
		usersRepository:          usersRepository,
		userProfilesRepository:   userProfilesRepository,
		servicesRepository:       servicesRepository,
		serviceImagesRepository:  serviceImagesRepository,
		certificationsRepository: certificationsRepository,
		storage:                  storage,
		logger:                   logger,
	}
}

func (s *onboardingsService) MakeOnboarding(ctx context.Context, r *http.Request, req *common.OnboardingRequest) error {
	s.logger.InfoContext(ctx, "starting onboarding process", "user_id", req.UserID)

	userProfile, err := s.userProfilesRepository.FindByUserID(ctx, req.UserID)
	if err != nil {
		s.logger.ErrorContext(ctx, "failed to verifiy if user profile already exists", "err", err)
		return exceptions.MakeGenericApiError()
	}
	if userProfile != nil {
		s.logger.WarnContext(ctx, "onboarding already exists", "user_id", req.UserID)
		return exceptions.MakeApiErrorWithStatus(http.StatusConflict, fmt.Errorf("onboarding already done"))
	}

	s.logger.InfoContext(ctx, "any user_profile found, making onboarding", "user_id", req.UserID)

	tx, err := s.db.Beginx()
	if err != nil {
		s.logger.ErrorContext(ctx, "error while starting transaction", "err", err)
		return exceptions.MakeGenericApiError()
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		}
	}()

	userProfile, err = s.createUserProfileTx(ctx, tx, r, req)
	if err != nil {
		s.logger.ErrorContext(ctx, "error while creating user_profile", "err", err)
		return exceptions.MakeGenericApiError()
	}

	if err := s.createCertificationsTx(ctx, tx, userProfile.ID(), req.Certifications); err != nil {
		s.logger.ErrorContext(ctx, "error while creating user certifications", "err", err)
		return exceptions.MakeGenericApiError()
	}

	if err := s.createServicesTx(ctx, tx, r, userProfile.UserID(), userProfile.ID(), req.Services); err != nil {
		s.logger.ErrorContext(ctx, "error while creating user services", "err", err)
		return exceptions.MakeGenericApiError()
	}

	if err := tx.Commit(); err != nil {
		s.logger.ErrorContext(ctx, "error while commiting onboarding transaction", "err", err)
		return exceptions.MakeGenericApiError()
	}

	s.logger.InfoContext(ctx, "onboarding completed successfully", "user_id", req.UserID)
	return nil
}

func (s *onboardingsService) createUserProfileTx(
	ctx context.Context,
	tx *sqlx.Tx,
	r *http.Request,
	input *common.OnboardingRequest,
) (*userprofiles.UserProfile, error) {
	var profileImageURL string
	profileImageFile, profileImageHeader, _ := r.FormFile("profile_image")
	if profileImageFile != nil {
		defer profileImageFile.Close()
		objectName := fmt.Sprintf("profiles/profile_%s", input.UserID)
		url, err := s.storage.UploadFile(objectName, profileImageHeader)
		if err != nil {
			s.logger.ErrorContext(ctx, "failed to upload profile image", "err", err)
			return nil, exceptions.MakeGenericApiError()
		}
		profileImageURL = url
	}

	profile, err := userprofiles.New(
		input.UserID,
		input.FullName,
		profileImageURL,
		input.JobDescription,
		input.Phone,
		input.SocialLinks,
	)
	if err != nil {
		s.logger.ErrorContext(ctx, "error while creating user profile entity", "err", err)
		return nil, exceptions.MakeApiErrorWithStatus(http.StatusUnprocessableEntity, fmt.Errorf("error creating user_profile entity"))
	}

	if err := s.userProfilesRepository.CreateTx(tx, profile); err != nil {
		s.logger.ErrorContext(ctx, "error while making create user profile transaction", "err", err)
		return nil, exceptions.MakeGenericApiError()
	}

	return profile, nil
}

func (s *onboardingsService) createCertificationsTx(ctx context.Context, tx *sqlx.Tx, profileID string, certs []common.Certification) error {
	for _, c := range certs {
		cert, err := certifications.New(
			profileID,
			c.Institution,
			c.CourseName,
			c.StartDate,
			&c.EndDate,
		)
		if err != nil {
			s.logger.ErrorContext(ctx, "error while creating certification entity", "err", err)
			return exceptions.MakeApiErrorWithStatus(http.StatusUnprocessableEntity, fmt.Errorf("error creating certification entity"))
		}

		if err := s.certificationsRepository.CreateTx(tx, cert); err != nil {
			s.logger.ErrorContext(ctx, "error while making create certification transaction", "err", err)
			return exceptions.MakeGenericApiError()
		}
	}
	return nil
}

func (s *onboardingsService) createServicesTx(
	ctx context.Context,
	tx *sqlx.Tx,
	r *http.Request,
	userID string,
	userProfileID string,
	svcs []common.ServiceWithImages,
) error {
	for i := range svcs {
		svc := &svcs[i]
		svc.ID = uid.New("service")

		imageURLs, err := s.uploadServiceImages(ctx, r, userID, svc.ID, i, svc.Name)
		if err != nil {
			return err
		}
		svc.Images = imageURLs

		if err := s.createServiceAndImages(ctx, tx, svc, userProfileID); err != nil {
			return err
		}
	}

	return nil
}

func (s *onboardingsService) uploadServiceImages(
	ctx context.Context,
	r *http.Request,
	userID, serviceID string,
	index int,
	serviceName string,
) ([]string, error) {
	formField := fmt.Sprintf("services[%d].images", index)
	files, ok := r.MultipartForm.File[formField]
	if !ok || len(files) == 0 {
		s.logger.InfoContext(ctx, "no service images found", "service", serviceName, "field", formField)
		return nil, nil
	}

	var urls []string
	for _, file := range files {
		objectName := fmt.Sprintf("services/%s/%s", userID, serviceID)
		url, err := s.storage.UploadFile(objectName, file)
		if err != nil {
			s.logger.ErrorContext(ctx, "failed to upload service image", "service", serviceName, "err", err)
			return nil, exceptions.MakeGenericApiError()
		}
		urls = append(urls, url)
	}

	return urls, nil
}

func (s *onboardingsService) createServiceAndImages(
	ctx context.Context,
	tx *sqlx.Tx,
	svc *common.ServiceWithImages,
	userProfileID string,
) error {
	service, err := services.New(svc.ID, userProfileID, svc.Name, svc.Description)
	if err != nil {
		s.logger.ErrorContext(ctx, "error creating service entity", "err", err)
		return exceptions.MakeApiErrorWithStatus(http.StatusUnprocessableEntity, fmt.Errorf("error creating service entity"))
	}

	if err := s.servicesRepository.CreateTx(tx, service); err != nil {
		s.logger.ErrorContext(ctx, "failed to insert service", "service_id", svc.ID)
		return exceptions.MakeGenericApiError()
	}

	for i, url := range svc.Images {
		img, err := serviceimages.New(svc.ID, url, i)
		if err != nil {
			s.logger.ErrorContext(ctx, "error creating service image entity", "err", err)
			return exceptions.MakeApiErrorWithStatus(http.StatusUnprocessableEntity, fmt.Errorf("error creating service image entity"))
		}
		if err := s.serviceImagesRepository.CreateTx(tx, img); err != nil {
			s.logger.ErrorContext(ctx, "failed to insert service image", "err", err)
			return exceptions.MakeGenericApiError()
		}
	}

	return nil
}
