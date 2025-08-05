package onboardings

import (
	"conecta-mare-server/internal/common"
	"conecta-mare-server/internal/modules/accounts/certifications"
	"conecta-mare-server/internal/modules/accounts/locations"
	"conecta-mare-server/internal/modules/accounts/projectimages"
	"conecta-mare-server/internal/modules/accounts/projects"
	"conecta-mare-server/internal/modules/accounts/serviceimages"
	"conecta-mare-server/internal/modules/accounts/services"
	"conecta-mare-server/internal/modules/accounts/subcategories"
	"conecta-mare-server/internal/modules/accounts/userprofiles"
	"conecta-mare-server/internal/modules/accounts/users"
	"conecta-mare-server/pkg/exceptions"
	"conecta-mare-server/pkg/storage"
	"conecta-mare-server/pkg/uid"
	"conecta-mare-server/pkg/valueobjects"
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
	projectsRepository projects.ProjectsRepository,
	projectImagesRepository projectimages.ProjectImagesRepository,
	certificationsRepository certifications.CertificationsRepository,
	subcategoriesRepository subcategories.SubcategoriesRepository,
	servicesRepository services.ServicesRepository,
	serviceImagesRepository serviceimages.ServiceImagesRepository,
	locationsRepository locations.LocationsRepository,
	storage *storage.StorageClient,
	logger *slog.Logger,
) OnboardingsService {
	return &onboardingsService{
		db:                       db,
		usersRepository:          usersRepository,
		userProfilesRepository:   userProfilesRepository,
		projectsRepository:       projectsRepository,
		projectImagesRepository:  projectImagesRepository,
		certificationsRepository: certificationsRepository,
		subcategoriesRepository:  subcategoriesRepository,
		serviceRepository:        servicesRepository,
		serviceImagesRepository:  serviceImagesRepository,
		locationRepository:       locationsRepository,
		storage:                  storage,
		logger:                   logger,
	}
}

func (s *onboardingsService) MakeOnboarding(
	ctx context.Context,
	r *http.Request,
	req *common.OnboardingRequest,
) error {
	s.logger.InfoContext(ctx, "starting onboarding process", "user_id", req.UserID)

	user, err := s.usersRepository.GetByID(ctx, req.UserID)
	if err != nil {
		s.logger.ErrorContext(ctx, "failed to verify if user already exists", "err", err)
		return exceptions.MakeGenericApiError()
	}
	if user == nil {
		s.logger.ErrorContext(ctx, "user does not exists", "user_id", req.UserID)
		return exceptions.MakeApiErrorWithStatus(http.StatusBadRequest, fmt.Errorf("user with id %s does not exists", req.UserID))
	}

	if user.Role != valueobjects.Professional {
		s.logger.WarnContext(ctx, "user trying to do onboarding as a client user", "user_id", req.UserID)
		return exceptions.MakeApiErrorWithStatus(http.StatusBadRequest, fmt.Errorf("only professional can do onboarding"))
	}

	userProfile, err := s.userProfilesRepository.FindByUserID(ctx, req.UserID)
	if err != nil {
		s.logger.ErrorContext(ctx, "failed to verifiy if user profile already exists", "err", err)
		return exceptions.MakeGenericApiError()
	}
	if userProfile == nil {
		s.logger.ErrorContext(ctx, "user profile for the given user was not found", "user_id", req.UserID)
		return exceptions.MakeApiErrorWithStatus(http.StatusInternalServerError, fmt.Errorf("consistency error: user profile for user %s not found", req.UserID))
	}
	if userProfile.JobDescription() != nil {
		s.logger.WarnContext(ctx, "onboarding already exists", "user_id", req.UserID)
		return exceptions.MakeApiErrorWithStatus(http.StatusConflict, fmt.Errorf("onboarding already done"))
	}

	subcategory, err := s.subcategoriesRepository.GetByID(ctx, req.SubcategoryID)
	if err != nil {
		s.logger.ErrorContext(ctx, "failed to verify if category exists", "category_id", req.SubcategoryID, "err", err)
		return exceptions.MakeGenericApiError()
	}
	if subcategory == nil {
		s.logger.WarnContext(ctx, "subcategory does not exists", "subcategory_id", req.SubcategoryID)
		return exceptions.MakeApiErrorWithStatus(http.StatusBadRequest, fmt.Errorf("subcategory with subcategory_id %s does not exists", req.SubcategoryID))
	}

	s.logger.InfoContext(ctx, "found user profile, starting onboarding update", "user_id", req.UserID)

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

	var profileImageURL string
	profileImageFile, profileImageHeader, _ := r.FormFile("profile_image")
	if profileImageFile != nil {
		defer profileImageFile.Close()
		objectName := fmt.Sprintf("profiles/profile_%s", req.UserID)
		url, uploadErr := s.storage.UploadFile(objectName, profileImageHeader)
		if uploadErr != nil {
			s.logger.ErrorContext(ctx, "failed to upload profile image", "err", uploadErr)
			err = exceptions.MakeGenericApiError()
			return err
		}
		profileImageURL = url
	}

	if err = userProfile.Update(
		req.SubcategoryID,
		profileImageURL,
		req.JobDescription,
		req.Phone,
		req.SocialLinks,
	); err != nil {
		s.logger.ErrorContext(ctx, "error validating user profile update", "err", err)
		return err
	}

	if err = s.userProfilesRepository.UpdateTx(ctx, tx, userProfile); err != nil {
		s.logger.ErrorContext(ctx, "error while making update user profile transaction", "err", err)
		return exceptions.MakeGenericApiError()
	}

	if err = s.createCertificationsTx(ctx, tx, userProfile.ID(), req.Certifications); err != nil {
		s.logger.ErrorContext(ctx, "error while creating user certifications", "err", err)
		return exceptions.MakeGenericApiError()
	}

	if err = s.createProjectsTx(ctx, tx, r, userProfile.UserID(), userProfile.ID(), req.Projects); err != nil {
		s.logger.ErrorContext(ctx, "error while creating user projects", "err", err)
		return exceptions.MakeGenericApiError()
	}

	if err = s.createServicesTx(
		ctx,
		tx,
		r,
		userProfile.UserID(),
		userProfile.ID(),
		req.Services,
	); err != nil {
		s.logger.ErrorContext(ctx, "error while creating user services", "err", err)
		return exceptions.MakeGenericApiError()
	}

	if err = s.createLocationTx(ctx, tx, userProfile.ID(), &req.Location); err != nil {
		s.logger.ErrorContext(ctx, "error while creating user location", "err", err)
		return exceptions.MakeGenericApiError()
	}

	if err = tx.Commit(); err != nil {
		s.logger.ErrorContext(ctx, "error while commiting onboarding transaction", "err", err)
		return exceptions.MakeGenericApiError()
	}

	s.logger.InfoContext(ctx, "onboarding completed successfully", "user_id", req.UserID)
	return nil
}

func (s *onboardingsService) createCertificationsTx(ctx context.Context, tx *sqlx.Tx, profileID string, certs []common.Certification) error {
	for _, c := range certs {
		cert, err := certifications.New(
			profileID,
			c.Institution,
			c.CourseName,
			c.StartDate,
			c.EndDate,
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

func (s *onboardingsService) createProjectsTx(
	ctx context.Context,
	tx *sqlx.Tx,
	r *http.Request,
	userID string,
	userProfileID string,
	pjts []common.Project,
) error {
	for i := range pjts {
		pjt := &pjts[i]
		pjt.ID = uid.New("project")

		imageWithIDs, err := s.uploadProjectImages(ctx, r, userID, pjt.ID, i, pjt.Name)
		if err != nil {
			return err
		}

		for _, img := range imageWithIDs {
			pjt.Images = append(pjt.Images, common.ProjectImage{ID: img.ID, URL: img.URL, Ordering: img.Ordering})
		}
		pjt.Images = imageWithIDs

		if err := s.createProjectAndImages(ctx, tx, pjt, userProfileID); err != nil {
			return err
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
	services []common.OnboardingService,
) error {
	for i := range services {
		service := &services[i]
		service.ID = uid.New("service")

		imageWithIDs, err := s.uploadServiceImages(ctx, r, userID, service.ID, i, service.Name)
		if err != nil {
			return err
		}

		for _, img := range imageWithIDs {
			service.Images = append(service.Images, common.ServiceImage{ID: img.ID, URL: img.URL, Ordering: img.Ordering})
		}
		service.Images = imageWithIDs

		if err := s.createServiceAndImages(ctx, tx, service, userProfileID); err != nil {
			return err
		}
	}

	return nil
}

func (s *onboardingsService) uploadProjectImages(
	ctx context.Context,
	r *http.Request,
	userID, projectID string,
	index int,
	projectName string,
) ([]common.ProjectImage, error) {
	formField := fmt.Sprintf("projects[%d].images", index)
	files, ok := r.MultipartForm.File[formField]
	if !ok || len(files) == 0 {
		s.logger.InfoContext(ctx, "no project images found", "project", projectName, "field", formField)
		return nil, nil
	}

	var result []common.ProjectImage
	for _, file := range files {
		imageID := uid.New("projectimg")
		objectName := fmt.Sprintf("projects/%s/%s/%s", userID, projectID, imageID)

		url, err := s.storage.UploadFile(objectName, file)
		if err != nil {
			s.logger.ErrorContext(ctx, "failed to upload project image", "project", projectName, "err", err)
			return nil, exceptions.MakeGenericApiError()
		}

		result = append(result, common.ProjectImage{
			ID:  imageID,
			URL: url,
		})
	}

	return result, nil
}

func (s *onboardingsService) uploadServiceImages(
	ctx context.Context,
	r *http.Request,
	userID, serviceID string,
	index int,
	serviceName string,
) ([]common.ServiceImage, error) {
	formField := fmt.Sprintf("services[%d].images", index)
	files, ok := r.MultipartForm.File[formField]
	if !ok || len(files) == 0 {
		s.logger.InfoContext(ctx, "no service images found", "service", serviceName, "field", formField)
		return nil, nil
	}

	var result []common.ServiceImage
	for _, file := range files {
		imageID := uid.New("service_img")
		objectName := fmt.Sprintf("services/%s/%s/%s", userID, serviceID, imageID)

		url, err := s.storage.UploadFile(objectName, file)
		if err != nil {
			s.logger.ErrorContext(ctx, "failed to upload service image", "project", serviceName, "err", err)
			return nil, exceptions.MakeGenericApiError()
		}

		result = append(result, common.ServiceImage{
			ID:  imageID,
			URL: url,
		})
	}

	return result, nil
}

func (s *onboardingsService) createProjectAndImages(
	ctx context.Context,
	tx *sqlx.Tx,
	prj *common.Project,
	userProfileID string,
) error {
	project, err := projects.New(prj.ID, userProfileID, prj.Name, prj.Description)
	if err != nil {
		s.logger.ErrorContext(ctx, "error creating project entity", "err", err)
		return exceptions.MakeApiErrorWithStatus(http.StatusUnprocessableEntity, fmt.Errorf("error creating service entity"))
	}

	if err := s.projectsRepository.CreateTx(tx, project); err != nil {
		s.logger.ErrorContext(ctx, "failed to insert project", "project_id", prj.ID)
		return exceptions.MakeGenericApiError()
	}

	for i, img := range prj.Images {
		projectImg, err := projectimages.New(img.ID, prj.ID, img.URL, i)
		if err != nil {
			s.logger.ErrorContext(ctx, "error creating project image entity", "err", err)
			return exceptions.MakeApiErrorWithStatus(http.StatusUnprocessableEntity, fmt.Errorf("error creating project image entity"))
		}
		if err := s.projectImagesRepository.CreateTx(tx, projectImg); err != nil {
			s.logger.ErrorContext(ctx, "failed to insert project image", "err", err)
			return exceptions.MakeGenericApiError()
		}
	}

	return nil
}

func (s *onboardingsService) createServiceAndImages(
	ctx context.Context,
	tx *sqlx.Tx,
	svc *common.OnboardingService,
	userProfileID string,
) error {
	service, err := services.New(svc.ID, userProfileID, svc.Name, svc.Description, svc.Price, svc.OwnLocationPrice)
	if err != nil {
		s.logger.ErrorContext(ctx, "error creating service entity", "err", err)
		return exceptions.MakeApiErrorWithStatus(http.StatusUnprocessableEntity, fmt.Errorf("error creating service entity"))
	}

	if err := s.serviceRepository.CreateTx(tx, service); err != nil {
		s.logger.ErrorContext(ctx, "failed to insert service", "service_id", svc.ID)
		return exceptions.MakeGenericApiError()
	}

	for i, img := range svc.Images {
		serviceImg, err := serviceimages.New(img.ID, svc.ID, img.URL, i)
		if err != nil {
			s.logger.ErrorContext(ctx, "error creating service image entity", "err", err)
			return exceptions.MakeApiErrorWithStatus(http.StatusUnprocessableEntity, fmt.Errorf("error creating project image entity"))
		}
		if err := s.serviceImagesRepository.CreateTx(tx, serviceImg); err != nil {
			s.logger.ErrorContext(ctx, "failed to insert service image", "err", err)
			return exceptions.MakeGenericApiError()
		}
	}

	return nil
}

func (s *onboardingsService) createLocationTx(
	ctx context.Context,
	tx *sqlx.Tx,
	userProfileID string,
	loc *common.OnboardingLocation,
) error {
	fmt.Println(loc)
	location, err := locations.New(
		userProfileID,
		loc.Street,
		loc.Number,
		loc.Complement,
		loc.CommunityID,
	)
	if err != nil {
		s.logger.ErrorContext(ctx, "error creating location entity", "err", err)
		return exceptions.MakeApiErrorWithStatus(http.StatusUnprocessableEntity, fmt.Errorf("error creating location entity"))
	}

	if err := s.locationRepository.CreateTx(tx, location); err != nil {
		fmt.Println(err)
		s.logger.ErrorContext(ctx, "failed to insert location", "user_profile_id", userProfileID)
		return exceptions.MakeGenericApiError()
	}

	return nil
}
