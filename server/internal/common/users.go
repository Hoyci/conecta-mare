package common

import (
	"conecta-mare-server/pkg/valueobjects"
	"encoding/json"
	"fmt"
	"time"
)

type JSONMap map[string]string

func (j *JSONMap) Scan(src any) error {
	bytes, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("expected []byte for JSONMap, got %T", src)
	}
	return json.Unmarshal(bytes, j)
}

type JSONB []byte

func (j *JSONB) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan JSONB: %v", value)
	}
	*j = b
	return nil
}

func (j JSONB) Unmarshal(v interface{}) error {
	return json.Unmarshal(j, v)
}

type (
	User struct {
		ID              string            `json:"id" db:"id"`
		Email           string            `json:"email" db:"email"`
		Role            valueobjects.Role `json:"role" db:"role"`
		FullName        *string           `json:"full_name" db:"full_name"`
		ProfileImage    *string           `json:"profile_image" db:"profile_image"`
		JobDescription  *string           `json:"job_description" db:"job_description"`
		SubcategoryName *string           `json:"subcategory_name" db:"name"`
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

	GetProfessionalsResponse struct {
		UserID         string `json:"user_id" db:"user_id"`
		FullName       string `json:"full_name" db:"full_name"`
		ProfileImage   string `json:"profile_image" db:"profile_image"`
		JobDescription string `json:"job_description" db:"job_description"`
		Rating         int    `json:"rating" db:"rating"`
		Location       string `json:"location" db:"location"`
	}

	GetProfessionalByIDRaw struct {
		UserID             string  `db:"user_id"`
		Email              string  `db:"email"`
		FullName           string  `db:"full_name"`
		ProfileImage       string  `db:"profile_image"`
		JobDescription     string  `db:"job_description"`
		Phone              string  `db:"phone"`
		SocialLinks        JSONMap `db:"social_links"`
		CategoryName       string  `db:"category_name"`
		SubcategoryName    string  `db:"subcategory_name"`
		ProjectsJSON       JSONB   `db:"projects"`
		CertificationsJSON JSONB   `db:"certifications"`
		Rating             int     `db:"rating"`
		Location           string  `db:"location"`
		ServicesJSON       JSONB   `db:"services"`
	}

	GetProfessionalByIDResponse struct {
		UserID          string          `json:"user_id" db:"user_id"`
		Email           string          `json:"email" db:"email"`
		FullName        string          `json:"full_name" db:"full_name"`
		ProfileImage    string          `json:"profile_image" db:"profile_image"`
		JobDescription  string          `json:"job_description" db:"job_description"`
		Phone           string          `json:"phone" db:"phone"`
		SocialLinks     JSONMap         `json:"social_links" db:"social_links"`
		SubcategoryName string          `json:"subcategory_name" db:"subcategory_name"`
		Rating          int             `json:"rating" db:"location"`
		Location        string          `json:"location" db:"location"`
		Projects        []Project       `json:"projects" db:"projects"`
		Certifications  []Certification `json:"certifications" db:"certifications"`
		Services        []Service       `json:"services" db:"services"`
	}

	Project struct {
		ID          string               `json:"id" db:"id"`
		Name        string               `json:"name" db:"name"`
		Description string               `json:"description" db:"description"`
		Images      []ProjectImageWithID `json:"images" db:"images"`
	}

	ProjectImageWithID struct {
		ID       string `json:"id" db:"id"`
		URL      string `json:"url" db:"url"`
		Ordering int    `json:"ordering" db:"ordering"`
	}

	Certification struct {
		Institution string     `json:"institution" db:"institution"`
		CourseName  string     `json:"course_name" db:"course_name"`
		StartDate   time.Time  `json:"start_date" db:"start_date"`
		EndDate     *time.Time `json:"end_date" db:"end_date"`
	}

	Service struct {
		Name             string               `json:"name" db:"name"`
		Description      string               `json:"description" db:"description"`
		Price            int                  `json:"price" db:"price"`
		OwnLocationPrice *int                 `json:"own_location_price" db:"own_location_price"`
		Images           []ServiceImageWithID `json:"images" db:"images"`
	}

	ServiceImageWithID struct {
		ID       string `json:"id" db:"id"`
		URL      string `json:"url" db:"url"`
		Ordering int    `json:"ordering" db:"ordering"`
	}

	Location struct {
		Street       string `json:"street" db:"street"`
		Number       string `json:"number" db:"number"`
		Complement   string `json:"complement" db:"complement"`
		Neighborhood string `json:"neighborhood" db:"neighborhood"`
	}
)
