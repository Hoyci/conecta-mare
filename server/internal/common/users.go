package common

import (
	"conecta-mare-server/pkg/valueobjects"
	"encoding/json"
	"time"
)

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
		FullName        string            `json:"full_name"`
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
		UserID             string          `db:"user_id"`
		Email              string          `db:"email"`
		FullName           string          `db:"full_name"`
		ProfileImage       string          `db:"profile_image"`
		JobDescription     string          `db:"job_description"`
		Phone              string          `db:"phone"`
		SocialLinks        json.RawMessage `db:"social_links"`
		Category           json.RawMessage `db:"category"`
		Subcategory        json.RawMessage `db:"subcategory"`
		ProjectsJSON       json.RawMessage `db:"projects"`
		CertificationsJSON json.RawMessage `db:"certifications"`
		Rating             int             `db:"rating"`
		Location           json.RawMessage `db:"location"`
		ServicesJSON       json.RawMessage `db:"services"`
	}

	GetProfessionalByIDResponse struct {
		UserID         string          `json:"user_id" db:"user_id"`
		Email          string          `json:"email" db:"email"`
		FullName       string          `json:"full_name" db:"full_name"`
		ProfileImage   string          `json:"profile_image" db:"profile_image"`
		JobDescription string          `json:"job_description" db:"job_description"`
		Phone          string          `json:"phone" db:"phone"`
		SocialLinks    json.RawMessage `json:"social_links" db:"social_links"`
		Category       json.RawMessage `json:"category" db:"category"`
		Subcategory    json.RawMessage `json:"subcategory" db:"subcategory"`
		Rating         int             `json:"rating" db:"location"`
		Location       json.RawMessage `json:"location" db:"location"`
		Projects       []Project       `json:"projects" db:"projects"`
		Certifications []Certification `json:"certifications" db:"certifications"`
		Services       []Service       `json:"services" db:"services"`
	}

	Project struct {
		ID          string         `json:"id"`
		Name        string         `json:"name"`
		Description string         `json:"description"`
		Images      []ProjectImage `json:"images"`
	}

	ProjectImage struct {
		ID       string `json:"id"`
		URL      string `json:"url"`
		Ordering int    `json:"ordering"`
	}

	Certification struct {
		ID          string     `json:"id"`
		Institution string     `json:"institution"`
		CourseName  string     `json:"course_name"`
		StartDate   time.Time  `json:"start_date"`
		EndDate     *time.Time `json:"end_date"`
	}

	Service struct {
		ID               string         `json:"id"`
		Name             string         `json:"name"`
		Description      string         `json:"description"`
		Price            int            `json:"price"`
		OwnLocationPrice *int           `json:"own_location_price"`
		Images           []ServiceImage `json:"images"`
	}

	ServiceImage struct {
		ID       string `json:"id"`
		URL      string `json:"url"`
		Ordering int    `json:"ordering"`
	}

	Location struct {
		Street      string `json:"street" db:"street"`
		Number      string `json:"number" db:"number"`
		Complement  string `json:"complement" db:"complement"`
		CommunityID string `json:"community_id" db:"community_id"`
	}
)
