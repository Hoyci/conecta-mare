package users

import (
	"conecta-mare-server/internal/common"
	"conecta-mare-server/internal/database/models"
	"conecta-mare-server/pkg/storage"
	"context"
	"log/slog"
)

type (
	UsersRepository interface {
		Register(ctx context.Context, user *User) error
		GetByID(ctx context.Context, ID string) (*models.User, error)
		GetByEmail(ctx context.Context, email string) (*models.User, error)
		GetByRole(ctx context.Context, role string) ([]*models.User, error)
		// Update(ctx context.Context, user *User) (*User, error)
		DeleteByID(ctx context.Context, ID string) error
	}
	UsersService interface {
		Register(ctx context.Context, input common.RegisterUserRequest) error
		GetByID(ctx context.Context, ID string) (*User, error)
		GetByEmail(ctx context.Context, email string) (*User, error)
		// Updated(ctx context.Context common.)
		DeleteByID(ctx context.Context, ID string) error
	}
	userService struct {
		repository    UsersRepository
		storageClient *storage.StorageClient
		logger        *slog.Logger
	}
	userHandler struct {
		usersService UsersService
	}
)
