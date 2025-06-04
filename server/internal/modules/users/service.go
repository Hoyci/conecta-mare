package users

import (
	"conecta-mare-server/internal/common"
	"conecta-mare-server/pkg/exceptions"
	"conecta-mare-server/pkg/storage"
	"conecta-mare-server/pkg/valueobjects"
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/lib/pq"
)

func NewService(repository UsersRepository, storageClient *storage.StorageClient) UsersService {
	return &userService{
		repository:    repository,
		storageClient: storageClient,
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
func (us *userService) Register(ctx context.Context, input common.RegisterUserRequest) (*User, error) {
	existingUser, err := us.repository.GetByEmail(ctx, input.Email)
	if err != nil {
		return nil, exceptions.MakeApiError(err)
	}

	if existingUser != nil {
		return nil, exceptions.MakeApiErrorWithStatus(http.StatusConflict, exceptions.ErrEmailTaken)
	}

	passwordHash, err := valueobjects.NewPassword(input.Password)
	if err != nil {
		return nil, exceptions.MakeGenericApiError()
	}

	user, err := New(
		input.Name,
		input.Email,
		passwordHash.Hash,
		input.Role,
		input.SubcategoryID,
	)
	if err != nil {
		return nil, exceptions.MakeApiErrorWithStatus(http.StatusUnprocessableEntity, err)
	}

	avatarUrl, err := us.UploadUserPicture(ctx, user.ID(), input.Avatar)
	if err != nil {
		return nil, exceptions.MakeGenericApiError()
	}

	user.avatarURL = avatarUrl

	if err := us.repository.Register(ctx, user); err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			return nil, exceptions.MakeApiErrorWithStatus(http.StatusConflict, fmt.Errorf("%s already taken", pqErr.Detail))
		}

		return nil, exceptions.MakeGenericApiError()
	}

	return user, nil
}

func (us *userService) UploadUserPicture(ctx context.Context, userID string, fileHeader *multipart.FileHeader) (string, error) {
	bucketName := "user-avatar"
	objectName := fmt.Sprintf("%s-%s", bucketName, userID)
	avatarURL, err := us.storageClient.UploadFile(bucketName, objectName, fileHeader)
	if err != nil {
		return "", err
	}

	return avatarURL, nil
}
