package users

import (
	"conecta-mare-server/internal/databases/postgres/models"
	"conecta-mare-server/pkg/exceptions"
	"conecta-mare-server/pkg/uid"
	"conecta-mare-server/pkg/valueobjects"
	"time"
)

type User struct {
	id           string
	email        string
	role         valueobjects.Role
	passwordHash string
	createdAt    time.Time
	updatedAt    *time.Time
	deletedAt    *time.Time
}

func New(
	email,
	passwordHash string,
	role valueobjects.Role,
) (*User, error) {
	user := User{
		id:           uid.New("user"),
		email:        email,
		role:         role,
		passwordHash: passwordHash,
		createdAt:    time.Now(),
		updatedAt:    nil,
		deletedAt:    nil,
	}

	if err := user.validate(); err != nil {
		return nil, exceptions.MakeApiError(err)
	}

	return &user, nil
}

func NewFromModel(m models.User) *User {
	return &User{
		id:           m.ID,
		email:        m.Email,
		passwordHash: m.PasswordHash,
		role:         m.Role,
		createdAt:    m.CreatedAt,
		updatedAt:    m.UpdatedAt,
		deletedAt:    m.DeletedAt,
	}
}

func (u *User) ToModel() models.User {
	return models.User{
		ID:           u.id,
		Email:        u.email,
		PasswordHash: u.passwordHash,
		Role:         u.role,
		CreatedAt:    u.createdAt,
		UpdatedAt:    u.updatedAt,
		DeletedAt:    u.deletedAt,
	}
}

func (u *User) validate() error {
	if _, err := valueobjects.NewEmail(u.email); err != nil {
		return err
	}
	if !u.role.IsValid() {
		return exceptions.ErrInvalidRole
	}
	return nil
}

func (u *User) ID() string {
	return u.id
}

func (u *User) Email() string {
	return u.email
}

func (u *User) PasswordHash() string {
	return u.passwordHash
}

func (u *User) Role() valueobjects.Role {
	return u.role
}

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

func (u *User) UpdatedAt() *time.Time {
	return u.updatedAt
}

func (u *User) DeletedAt() *time.Time {
	return u.deletedAt
}
