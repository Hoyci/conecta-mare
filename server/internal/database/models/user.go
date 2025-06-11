package models

import (
	"conecta-mare-server/pkg/valueobjects"
	"time"
)

type User struct {
	ID           string            `db:"id"`
	Name         string            `db:"name"`
	Email        string            `db:"email"`
	Role         valueobjects.Role `db:"role"`
	PasswordHash string            `db:"password_hash"`
	CreatedAt    time.Time         `db:"created_at"`
	UpdatedAt    *time.Time        `db:"updated_at"`
	DeletedAt    *time.Time        `db:"deleted_at"`
}
