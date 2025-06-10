package common

import (
	"conecta-mare-server/pkg/valueobjects"
	"mime/multipart"
)

type (
	User struct {
		ID            string            `json:"id"`
		Name          string            `json:"name"`
		Email         string            `json:"email"`
		Role          valueobjects.Role `json:"role"`
		AvatarURL     string            `json:"avatar_url"`
		SubcategoryID *string           `json:"subcategory_id"`
	}

	RegisterUserRequest struct {
		Name            string                `form:"name"`
		Email           string                `form:"email"`
		Role            valueobjects.Role     `form:"role"`
		Password        string                `form:"password"`
		ConfirmPassword string                `form:"confirm_password"`
		Avatar          *multipart.FileHeader `form:"-"`
		SubcategoryID   *string               `form:"subcategory_id,omitempty"`
	}

	RegisterUserResponse struct {
		Message string `json:"message"`
	}

	LoginUserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	LoginUserResponse struct {
		AccessToken  *string `json:"access_token"`
		RefreshToken *string `json:"refresh_token"`
	}

	UserResponse struct {
		User *User `json:"user"`
	}
)
