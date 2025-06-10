package users

import (
	"conecta-mare-server/internal/common"
	"conecta-mare-server/internal/modules/session"
	"conecta-mare-server/pkg/exceptions"
	"conecta-mare-server/pkg/jwt"
	"conecta-mare-server/pkg/security"
	"conecta-mare-server/pkg/storage"
	"conecta-mare-server/pkg/valueobjects"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"mime/multipart"
	"net/http"

	"github.com/lib/pq"
)

func NewService(
	repository UsersRepository,
	sessionsService session.SessionsService,
	storageClient *storage.StorageClient,
	tokenProvider jwt.JWTProvider,
	logger *slog.Logger,
) UsersService {
	return &userService{
		repository:     repository,
		sessionService: sessionsService,
		storageClient:  storageClient,
		logger:         logger,
	}
}

func (s *userService) DeleteByID(ctx context.Context, ID string) error {
	panic("unimplemented")
}

func (s *userService) GetByEmail(ctx context.Context, email string) (*User, error) {
	panic("unimplemented")
}

func (s *userService) GetByID(ctx context.Context, ID string) (*User, error) {
	panic("unimplemented")
}

func (s *userService) Register(ctx context.Context, input common.RegisterUserRequest) error {
	// TODO: adicionar verificação de subcategoryID para validar se realmente existingser
	s.logger.InfoContext(ctx, "atempting to create user", "email", input.Email)

	existingser, err := s.repository.GetByEmail(ctx, input.Email)
	if err != nil {
		s.logger.ErrorContext(ctx, "error while attempting check for existing user", "err", err, "email", input.Email)
		return exceptions.MakeApiError(err)
	}

	if existingser != nil {
		s.logger.InfoContext(ctx, "email is taken", "email", input.Email)
		return exceptions.MakeApiErrorWithStatus(http.StatusConflict, exceptions.ErrEmailTaken)
	}

	passwordHash, err := valueobjects.NewPassword(input.Password)
	if err != nil {
		s.logger.ErrorContext(ctx, "error while attempting to hash user password", "err", err, "email", input.Email)
		return exceptions.MakeGenericApiError()
	}

	s.logger.InfoContext(ctx, "password successfully hashed, creating user", "email", input.Email)
	user, err := New(
		input.Name,
		input.Email,
		passwordHash.Hash,
		input.Role,
		input.SubcategoryID,
	)
	if err != nil {
		s.logger.ErrorContext(ctx, "error while attempting to process user entity", "err", err, "email", input.Email)
		return exceptions.MakeApiErrorWithStatus(http.StatusUnprocessableEntity, err)
	}

	avatarUrl, err := s.UploadUserPicture(ctx, user.ID(), input.Avatar)
	if err != nil {
		s.logger.ErrorContext(ctx, "error while attempting to upload user avatar", "err", err, "email", input.Email)
		return exceptions.MakeGenericApiError()
	}

	user.avatarURL = avatarUrl
	s.logger.InfoContext(ctx, "user avatar successfully created", "email", input.Email)

	if err := s.repository.Register(ctx, user); err != nil {
		s.logger.ErrorContext(ctx, "error while attempting create user", "err", err, "user", user)
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			return exceptions.MakeApiErrorWithStatus(http.StatusConflict, fmt.Errorf("%s already taken", pqErr.Detail))
		}

		return exceptions.MakeGenericApiError()
	}

	s.logger.InfoContext(ctx, "user successfully created", "user_id", user.ID())
	return nil
}

func (s *userService) Login(ctx context.Context, input common.LoginUserRequest) (*common.LoginUserResponse, error) {
	s.logger.InfoContext(ctx, "attempting to login user, checking for existing user", "email", input.Email)

	existingUser, err := s.repository.GetByEmail(ctx, input.Email)
	if err != nil {
		s.logger.ErrorContext(ctx, "error while attempting to query for existing users", "err", err)
		return nil, exceptions.MakeGenericApiError()
	}
	if existingUser == nil {
		s.logger.InfoContext(ctx, "user was not found", "email", input.Email)
		return nil, exceptions.MakeApiErrorWithStatus(http.StatusNotFound, exceptions.ErrUserNotFound)
	}

	if existingUser.DeletedAt != nil {
		s.logger.InfoContext(ctx, "user must be active to login", "email", input.Email)
		return nil, exceptions.MakeApiErrorWithStatus(http.StatusUnauthorized, exceptions.ErrUserDisabled)
	}

	user := NewFromModel(*existingUser)

	s.logger.InfoContext(ctx, "user found, attempting to verify password", "email", user.Email())
	if !security.PasswordMatches(input.Password, user.PasswordHash()) {
		s.logger.ErrorContext(ctx, "unauthorized attempt to login", "email", input.Email)
		return nil, exceptions.MakeApiErrorWithStatus(http.StatusBadRequest, exceptions.ErrInvalidLoginAttempt)
	}

	err = s.sessionService.DeactivateAllSessions(ctx, user.ID())
	if err != nil {
		s.logger.ErrorContext(ctx, "error while attempting to deactivate all user sessions", "user_id", existingUser.ID, "err", err)
		return nil, exceptions.MakeGenericApiError()
	}

	accessToken, _, err := s.tokenProvider.GenerateAccessToken(user)
	if err != nil {
		s.logger.ErrorContext(ctx, "error while attempting to create user access token", "user", user, "err", err)
		return nil, exceptions.MakeGenericApiError()
	}

	refreshToken, claims, err := s.tokenProvider.GenerateAccessToken(user)
	if err != nil {
		s.logger.ErrorContext(ctx, "error while attempting to create user refresh token", "user", user, "err", err)
		return nil, exceptions.MakeGenericApiError()
	}

	_, err = s.sessionService.CreateSession(
		ctx, common.CreateSessionRequest{UserID: user.ID(), JTI: claims.ID},
	)
	if err != nil {
		s.logger.ErrorContext(ctx, "error while attempting to create session", "err", err)
		return nil, exceptions.MakeGenericApiError()
	}

	s.logger.InfoContext(ctx, "access token, refresh token and session created", "user_id", user.ID())

	return &common.LoginUserResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *userService) UploadUserPicture(ctx context.Context, userID string, fileHeader *multipart.FileHeader) (string, error) {
	if fileHeader == nil {
		return "", fmt.Errorf("file header is nil")
	}

	objectName := fmt.Sprintf("%s-%s", "avatar", userID)
	avatarURL, err := s.storageClient.UploadFile(objectName, fileHeader)
	if err != nil {
		return "", err
	}

	return avatarURL, nil
}
