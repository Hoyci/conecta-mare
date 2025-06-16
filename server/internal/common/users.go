package common

import (
	"conecta-mare-server/pkg/valueobjects"
	"encoding/json"
	"fmt"
)

type JSONMap map[string]string

func (j *JSONMap) Scan(src any) error {
	bytes, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("expected []byte for JSONMap, got %T", src)
	}
	return json.Unmarshal(bytes, j)
}

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

	ProfessionalResponse struct {
		UserID          string  `json:"user_id" db:"user_id"`
		Email           string  `json:"email" db:"email"`
		Role            string  `json:"role" db:"role"`
		FullName        string  `json:"full_name" db:"full_name"`
		ProfileImage    string  `json:"profile_image" db:"profile_image"`
		JobDescription  string  `json:"job_description" db:"job_description"`
		Phone           string  `json:"phone" db:"phone"`
		SocialLinks     JSONMap `json:"social_links" db:"social_links"`
		CategoryName    string  `json:"category_name" db:"category_name"`
		SubcategoryName string  `json:"subcategory_name" db:"subcategory_name"`
		Rating          int     `json:"rating" db:"rating"`
		Location        string  `json:"location" db:"location"`
	}
)
