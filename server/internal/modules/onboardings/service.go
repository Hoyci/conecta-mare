package onboardings

import (
	"conecta-mare-server/internal/common"
	"conecta-mare-server/pkg/exceptions"
	"conecta-mare-server/pkg/storage"
	"conecta-mare-server/pkg/uid"
	"context"
	"fmt"
	"log/slog"
	"net/http"
)

func NewService(repository OnboardingsRepository, storage *storage.StorageClient, logger *slog.Logger) OnboardingsService {
	return &onboardingService{
		repository: repository,
		storage:    storage,
		logger:     logger,
	}
}

func (s *onboardingService) CompleteOnboarding(ctx context.Context, req *common.OnboardingRequest, r *http.Request) error {
	s.logger.InfoContext(ctx, "starting onboarding process", "user_id", req.UserID)

	onboarding, err := s.repository.GetUserProfileByUserID(ctx, req.UserID)
	if err != nil {
		s.logger.ErrorContext(ctx, "failed to verifiy if user profile already exists", "err", err)
		return exceptions.MakeGenericApiError()
	}
	if onboarding != nil {
		s.logger.WarnContext(ctx, "onboarding already exists", "user_id", req.UserID)
		return exceptions.MakeApiErrorWithStatus(http.StatusConflict, fmt.Errorf("onboarding already done"))
	}

	var profileImageURL string
	profileImageFile, profileImageHeader, _ := r.FormFile("profile_image")
	if profileImageFile != nil {
		defer profileImageFile.Close()
		objectName := fmt.Sprintf("profiles/profile_%s", req.UserID)
		url, err := s.storage.UploadFile(objectName, profileImageHeader)
		if err != nil {
			s.logger.ErrorContext(ctx, "failed to upload profile image", "err", err)
			return exceptions.MakeGenericApiError()
		}
		profileImageURL = url
	}

	for i := range req.Services {
		req.Services[i].Images = []string{}

		formFieldName := fmt.Sprintf("services[%d].images", i)
		serviceImageHeaders, ok := r.MultipartForm.File[formFieldName]
		if !ok {
			continue
		}

		for _, fileHeader := range serviceImageHeaders {
			serviceID := uid.New("service")
			objectName := fmt.Sprintf("services/%s/%s", req.UserID, serviceID)
			url, err := s.storage.UploadFile(objectName, fileHeader)
			if err != nil {
				s.logger.ErrorContext(ctx, "failed to upload service image", "err", err, "service_name", req.Services[i].Name)
				return exceptions.MakeGenericApiError()
			}
			req.Services[i].ID = serviceID
			req.Services[i].Images = append(req.Services[i].Images, url)
		}
	}

	userProfile, err := NewUserProfile(
		req.UserID,
		req.FullName,
		profileImageURL,
		req.JobDescription,
		req.Phone,
		req.SocialLinks,
	)
	if err != nil {
		return err
	}

	var certifications []*Certification
	for _, c := range req.Certifications {
		cert, err := NewCertification(userProfile.ID(), c.Institution, c.CourseName, c.StartDate, &c.EndDate)
		if err != nil {
			return err
		}
		certifications = append(certifications, cert)
	}

	var services []*Service
	var serviceImages []*ServiceImage
	for _, svc := range req.Services {
		service, err := NewServices(userProfile.ID(), svc.ID, svc.Name, svc.Description)
		if err != nil {
			return err
		}
		services = append(services, service)

		for i, imgURL := range svc.Images {
			serviceImg, err := NewServiceImage(service.ID(), imgURL, i)
			if err != nil {
				return err
			}
			serviceImages = append(serviceImages, serviceImg)
		}
	}

	if err := s.repository.CreateUserProfile(ctx, userProfile, certifications, services, serviceImages); err != nil {
		s.logger.ErrorContext(ctx, "failed to save onboarding data", "err", err)
		return exceptions.MakeGenericApiError()
	}

	s.logger.InfoContext(ctx, "onboarding completed successfully", "user_id", req.UserID)
	return nil
}
