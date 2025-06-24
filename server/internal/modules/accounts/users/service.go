package users

import (
	"conecta-mare-server/internal/common"
	"conecta-mare-server/internal/modules/accounts/session"
	"conecta-mare-server/internal/server/middlewares"
	"conecta-mare-server/pkg/exceptions"
	"conecta-mare-server/pkg/jwt"
	"conecta-mare-server/pkg/security"
	"conecta-mare-server/pkg/storage"
	"conecta-mare-server/pkg/valueobjects"
	"context"
	"errors"
	"fmt"
	"log/slog"
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
	s.logger.InfoContext(ctx, "attempting to create user", "email", input.Email)

	existingser, err := s.repository.GetByEmail(ctx, input.Email)
	if err != nil {
		s.logger.ErrorContext(ctx, "error while attempting check for existing user", "err", err, "email", input.Email)
		return exceptions.MakeGenericApiError()
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
		input.Email,
		passwordHash.Hash,
		input.Role,
	)
	if err != nil {
		s.logger.ErrorContext(ctx, "error while attempting to process user entity", "err", err, "email", input.Email)
		return exceptions.MakeApiErrorWithStatus(http.StatusUnprocessableEntity, err)
	}

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

func (s *userService) Login(ctx context.Context, input common.LoginUserRequest) (*common.LoginUserResponse, *exceptions.ApiError[string]) {
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
	if ok := security.PasswordMatches(input.Password, user.PasswordHash()); !ok {
		s.logger.ErrorContext(ctx, "unauthorized attempt to login", "email", input.Email)
		return nil, exceptions.MakeApiErrorWithStatus(http.StatusUnauthorized, exceptions.ErrInvalidLoginAttempt)
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

func (s *userService) Logout(ctx context.Context) *exceptions.ApiError[string] {
	s.logger.InfoContext(ctx, "attempting to logout user")

	c, ok := ctx.Value(middlewares.AuthKey{}).(*jwt.Claims)
	if !ok {
		s.logger.ErrorContext(ctx, "error while attempting to get auth values from context")
		return exceptions.MakeApiErrorWithStatus(http.StatusUnauthorized, exceptions.ErrUnauthorized)
	}

	s.logger.InfoContext(ctx, "destructured data from context, logging out user", "user_id", c.UserID)

	activeSession, err := s.sessionService.GetActiveSessionByUserID(ctx, c.UserID)
	if err != nil {
		s.logger.ErrorContext(ctx, "error while attempting to get user active sessions", "user_id", c.UserID, "err", err)
		return exceptions.MakeGenericApiError()
	}

	if activeSession == nil {
		s.logger.WarnContext(ctx, "active session not found for user", "user_id", c.UserID)
		return exceptions.MakeApiErrorWithStatus(http.StatusBadRequest, exceptions.ErrActiveSessionNotFound)
	}

	activeSession.Deactivate()

	sess, err := s.sessionService.UpdateSession(ctx, activeSession)
	if err != nil {
		s.logger.ErrorContext(ctx, "error while attempting to update user session", "user_id", c.UserID, "err", err)
		return exceptions.MakeGenericApiError()
	}

	s.logger.InfoContext(ctx, "user logged out with success", "user_id", c.UserID, "session_id", sess.ToModel().ID)
	return nil
}

func (s *userService) GetSigned(ctx context.Context) (*User, *exceptions.ApiError[string]) {
	s.logger.InfoContext(ctx, "attempting to logout user")

	c, ok := ctx.Value(middlewares.AuthKey{}).(*jwt.Claims)
	if !ok {
		s.logger.ErrorContext(ctx, "error while attempting to get auth values from context")
		return nil, exceptions.MakeApiErrorWithStatus(http.StatusUnauthorized, exceptions.ErrUnauthorized)
	}

	user, err := s.repository.GetByID(ctx, c.UserID)
	if err != nil {
		s.logger.ErrorContext(ctx, "error while attempting to get user by id", "user_id", c.UserID, "err", err)
		return nil, exceptions.MakeGenericApiError()
	}

	if user == nil {
		s.logger.WarnContext(ctx, "user not found", "user_id", c.UserID)
		return nil, exceptions.MakeApiErrorWithStatus(http.StatusNotFound, exceptions.ErrUserNotFound)
	}

	userSession, err := s.sessionService.GetActiveSessionByUserID(ctx, c.UserID)
	if err != nil {
		s.logger.ErrorContext(ctx, "error whilhe attempting to get user session", "user_id", user.ID, "err", err)
		return nil, exceptions.MakeGenericApiError()
	}
	if userSession == nil {
		s.logger.WarnContext(ctx, "session not found", "user_id", user.ID)
		return nil, exceptions.MakeApiErrorWithStatus(http.StatusUnauthorized, exceptions.ErrActiveSessionNotFound)
	}

	return NewFromModel(*user), nil
}

func (s *userService) CountUsersBySubcategoryIDs(ctx context.Context, subcategoryIDs []string) (map[string]int, error) {
	s.logger.InfoContext(ctx, "attempting to count user by subcategory IDs")

	count, err := s.repository.CountBySubcategoryIDs(ctx, subcategoryIDs)
	if err != nil {
		s.logger.ErrorContext(ctx, "error while attempting to count user by subcategory IDs", "err", err)
		return nil, err
	}

	return count, nil
}

func (s *userService) GetProfessionals(ctx context.Context) ([]*common.GetProfessionalsResponse, *exceptions.ApiError[string]) {
	s.logger.InfoContext(ctx, "attemping to get professional users")

	professionals, err := s.repository.GetProfessionalUsers(ctx)
	if err != nil {
		s.logger.ErrorContext(ctx, "error while attempting to get professional users", "err", err)
		return nil, exceptions.MakeGenericApiError()
	}

	if professionals == nil {
		s.logger.WarnContext(ctx, "any professional user found")
		return nil, nil
	}

	return professionals, nil
}

func (s *userService) GetProfessionalByID(ctx context.Context, ID string) (*common.GetProfessionalByIDResponse, *exceptions.ApiError[string]) {
	s.logger.InfoContext(ctx, "attemping to get professional user by ID", "id", ID)

	professional, err := s.repository.GetProfessionalByID(ctx, ID)
	if err != nil {
		s.logger.ErrorContext(ctx, "error while attempting to get professional user", "err", err)
		return nil, exceptions.MakeGenericApiError()
	}

	if professional == nil {
		s.logger.WarnContext(ctx, "professional user not found")
		return nil, nil
	}

	return professional, nil
}
