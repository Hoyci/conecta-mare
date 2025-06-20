package onboardings

import (
	"conecta-mare-server/internal/common"
	"conecta-mare-server/internal/modules/accounts/categories"
	"conecta-mare-server/internal/modules/accounts/certifications"
	"conecta-mare-server/internal/modules/accounts/projectimages"
	"conecta-mare-server/internal/modules/accounts/projects"
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
	categoriesRepository categories.CategoriesRepository,
	subcategoriesRepository subcategories.SubcategoriesRepository,
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
		categoriesRepository:     categoriesRepository,
		subcategoriesRepository:  subcategoriesRepository,
		storage:                  storage,
		logger:                   logger,
	}
}

func (s *onboardingsService) MakeOnboarding(ctx context.Context, r *http.Request, req *common.OnboardingRequest) error {
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
	if userProfile != nil {
		s.logger.WarnContext(ctx, "onboarding already exists", "user_id", req.UserID)
		return exceptions.MakeApiErrorWithStatus(http.StatusConflict, fmt.Errorf("onboarding already done"))
	}

	category, err := s.categoriesRepository.GetByID(ctx, req.CategoryID)
	if err != nil {
		s.logger.ErrorContext(ctx, "failed to verify if category exists", "category_id", req.CategoryID, "err", err)
		return exceptions.MakeGenericApiError()
	}
	if category == nil {
		s.logger.WarnContext(ctx, "category does not exists", "category_id", req.CategoryID)
		return exceptions.MakeApiErrorWithStatus(http.StatusBadRequest, fmt.Errorf("category with category_id %s does not exists", req.CategoryID))
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

	if err := s.createProjectsTx(ctx, tx, r, userProfile.UserID(), userProfile.ID(), req.Projects); err != nil {
		s.logger.ErrorContext(ctx, "error while creating user projects", "err", err)
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
		input.CategoryID,
		input.SubcategoryID,
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
	prjs []common.Project,
) error {
	for i := range prjs {
		svc := &prjs[i]
		svc.ID = uid.New("project")

		imageWithIDs, err := s.uploadServiceImages(ctx, r, userID, svc.ID, i, svc.Name)
		if err != nil {
			return err
		}

		for _, img := range imageWithIDs {
			svc.Images = append(svc.Images, common.ProjectImageWithID{ID: img.ID, URL: img.URL, Ordering: img.Ordering})
		}
		svc.Images = imageWithIDs

		if err := s.createServiceAndImages(ctx, tx, svc, userProfileID); err != nil {
			return err
		}
	}

	return nil
}

func (s *onboardingsService) uploadServiceImages(
	ctx context.Context,
	r *http.Request,
	userID, projectID string,
	index int,
	projectName string,
) ([]common.ProjectImageWithID, error) {
	formField := fmt.Sprintf("projects[%d].images", index)
	files, ok := r.MultipartForm.File[formField]
	if !ok || len(files) == 0 {
		s.logger.InfoContext(ctx, "no project images found", "project", projectName, "field", formField)
		return nil, nil
	}

	var result []common.ProjectImageWithID
	for _, file := range files {
		imageID := uid.New("projectimg")
		objectName := fmt.Sprintf("projects/%s/%s/%s", userID, projectID, imageID)

		url, err := s.storage.UploadFile(objectName, file)
		if err != nil {
			s.logger.ErrorContext(ctx, "failed to upload project image", "project", projectName, "err", err)
			return nil, exceptions.MakeGenericApiError()
		}

		result = append(result, common.ProjectImageWithID{
			ID:  imageID,
			URL: url,
		})
	}

	return result, nil
}

func (s *onboardingsService) createServiceAndImages(
	ctx context.Context,
	tx *sqlx.Tx,
	prj *common.Project,
	userProfileID string,
) error {
	project, err := projects.New(prj.ID, userProfileID, prj.Name, prj.Description)
	if err != nil {
		s.logger.ErrorContext(ctx, "error creating service entity", "err", err)
		return exceptions.MakeApiErrorWithStatus(http.StatusUnprocessableEntity, fmt.Errorf("error creating service entity"))
	}

	if err := s.projectsRepository.CreateTx(tx, project); err != nil {
		s.logger.ErrorContext(ctx, "failed to insert service", "service_id", prj.ID)
		return exceptions.MakeGenericApiError()
	}

	for i, img := range prj.Images {
		serviceImg, err := projectimages.New(img.ID, prj.ID, img.URL, i)
		if err != nil {
			s.logger.ErrorContext(ctx, "error creating project image entity", "err", err)
			return exceptions.MakeApiErrorWithStatus(http.StatusUnprocessableEntity, fmt.Errorf("error creating project image entity"))
		}
		if err := s.projectImagesRepository.CreateTx(tx, serviceImg); err != nil {
			s.logger.ErrorContext(ctx, "failed to insert project image", "err", err)
			return exceptions.MakeGenericApiError()
		}
	}

	return nil
}
