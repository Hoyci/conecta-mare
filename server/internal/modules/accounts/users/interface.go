package users

import (
	"conecta-mare-server/internal/common"
	"conecta-mare-server/internal/database/models"
	"conecta-mare-server/internal/modules/accounts/session"
	"conecta-mare-server/internal/modules/accounts/userprofiles"
	"conecta-mare-server/pkg/exceptions"
	"conecta-mare-server/pkg/jwt"
	"conecta-mare-server/pkg/storage"
	"context"
	"log/slog"

	"github.com/jmoiron/sqlx"
)

type (
	UsersRepository interface {
		Register(ctx context.Context, tx *sqlx.Tx, user *User) error
		GetByID(ctx context.Context, ID string) (*models.User, error)
		GetByEmail(ctx context.Context, email string) (*models.User, error)
		GetByRole(ctx context.Context, role string) ([]*models.User, error)
		CountBySubcategoryIDs(ctx context.Context, subcategoryIDs []string) (map[string]int, error)
		// Update(ctx context.Context, user *User) (*User, error)
		DeleteByID(ctx context.Context, ID string) error
		GetProfessionalUsers(ctx context.Context) ([]*common.GetProfessionalsResponse, error)
		GetProfessionalByID(ctx context.Context, ID string) (*common.GetProfessionalByIDResponse, error)
	}
	UsersService interface {
		Login(ctx context.Context, input common.LoginUserRequest) (*common.LoginUserResponse, *exceptions.ApiError[string])
		Logout(ctx context.Context) *exceptions.ApiError[string]
		Register(ctx context.Context, input common.RegisterUserRequest) error
		GetSigned(ctx context.Context) (*User, *exceptions.ApiError[string])
		CountUsersBySubcategoryIDs(ctx context.Context, subcategoryIDs []string) (map[string]int, error)
		GetByID(ctx context.Context, ID string) (*User, error)
		GetByEmail(ctx context.Context, email string) (*User, error)
		// Updated(ctx context.Context common.)
		DeleteByID(ctx context.Context, ID string) error
		GetProfessionals(ctx context.Context) ([]*common.GetProfessionalsResponse, *exceptions.ApiError[string])
		GetProfessionalByID(ctx context.Context, ID string) (*common.GetProfessionalByIDResponse, *exceptions.ApiError[string])
	}
	userService struct {
		db               *sqlx.DB
		repository       UsersRepository
		userProfilesRepo userprofiles.UserProfilesRepository
		sessionService   session.SessionsService
		tokenProvider    jwt.JWTProvider
		storageClient    *storage.StorageClient
		logger           *slog.Logger
	}
	userHandler struct {
		usersService UsersService
		accessKey    string
	}
)
