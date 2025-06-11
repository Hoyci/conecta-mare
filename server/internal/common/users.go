package common

import (
	"conecta-mare-server/pkg/valueobjects"
)

type (
	User struct {
		ID            string            `json:"id"`
		Name          string            `json:"name"`
		Email         string            `json:"email"`
		Role          valueobjects.Role `json:"role"`
		AvatarURL     string            `json:"avatar_url"`
		SubcategoryID *string           `json:"subcategory_id,omitempty"`
	}

	RegisterUserRequest struct {
		Name            string            `json:"name"`
		Email           string            `json:"email"`
		Role            valueobjects.Role `json:"role"`
		Password        string            `json:"password"`
		ConfirmPassword string            `json:"confirm_password"`
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
