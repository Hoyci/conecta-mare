package users

import (
	"conecta-mare-server/internal/common"
	"conecta-mare-server/pkg/exceptions"
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

func NewService(repository UsersRepository, storageClient *storage.StorageClient, logger *slog.Logger) UsersService {
	return &userService{
		repository:    repository,
		storageClient: storageClient,
		logger:        logger,
	}
}

// DeleteByID implements UsersService.
func (us *userService) DeleteByID(ctx context.Context, ID string) error {
	panic("unimplemented")
}

// GetByEmail implements UsersService.
func (us *userService) GetByEmail(ctx context.Context, email string) (*User, error) {
	panic("unimplemented")
}

// GetByID implements UsersService.
func (us *userService) GetByID(ctx context.Context, ID string) (*User, error) {
	panic("unimplemented")
}

// Register implements UsersService.
func (us *userService) Register(ctx context.Context, input common.RegisterUserRequest) error {
	// TODO: adicionar verificação de subcategoryID para validar se realmente existingUser
	us.logger.InfoContext(ctx, "atempting to create user", "email", input.Email)

	existingUser, err := us.repository.GetByEmail(ctx, input.Email)
	if err != nil {
		us.logger.ErrorContext(ctx, "error while attempting check for existing user", "err", err, "email", input.Email)
		return exceptions.MakeApiError(err)
	}

	if existingUser != nil {
		us.logger.InfoContext(ctx, "email is taken", "email", input.Email)
		return exceptions.MakeApiErrorWithStatus(http.StatusConflict, exceptions.ErrEmailTaken)
	}

	passwordHash, err := valueobjects.NewPassword(input.Password)
	if err != nil {
		us.logger.ErrorContext(ctx, "error while attempting to hash user password", "err", err, "email", input.Email)
		return exceptions.MakeGenericApiError()
	}

	us.logger.InfoContext(ctx, "password successfully hashed, creating user", "email", input.Email)
	user, err := New(
		input.Name,
		input.Email,
		passwordHash.Hash,
		input.Role,
		input.SubcategoryID,
	)
	if err != nil {
		us.logger.ErrorContext(ctx, "error while attempting to process user entity", "err", err, "email", input.Email)
		return exceptions.MakeApiErrorWithStatus(http.StatusUnprocessableEntity, err)
	}

	avatarUrl, err := us.UploadUserPicture(ctx, user.ID(), input.Avatar)
	if err != nil {
		us.logger.ErrorContext(ctx, "error while attempting to upload user avatar", "err", err, "email", input.Email)
		return exceptions.MakeGenericApiError()
	}

	user.avatarURL = avatarUrl
	us.logger.InfoContext(ctx, "user avatar successfully created", "email", input.Email)

	if err := us.repository.Register(ctx, user); err != nil {
		us.logger.ErrorContext(ctx, "error while attempting create user", "err", err, "user", user)
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			return exceptions.MakeApiErrorWithStatus(http.StatusConflict, fmt.Errorf("%s already taken", pqErr.Detail))
		}

		return exceptions.MakeGenericApiError()
	}

	us.logger.InfoContext(ctx, "user successfully created", "user_id", user.ID())
	return nil
}

func (us *userService) UploadUserPicture(ctx context.Context, userID string, fileHeader *multipart.FileHeader) (string, error) {
	if fileHeader == nil {
		return "", fmt.Errorf("file header is nil")
	}

	objectName := fmt.Sprintf("%s-%s", "avatar", userID)
	avatarURL, err := us.storageClient.UploadFile(objectName, fileHeader)
	if err != nil {
		return "", err
	}

	return avatarURL, nil
}
