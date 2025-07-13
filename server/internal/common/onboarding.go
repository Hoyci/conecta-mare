package common

import "time"

type (
	OnboardingRequest struct {
		UserID         string              `json:"user_id"`
		FullName       string              `json:"full_name"`
		SubcategoryID  string              `json:"subcategory_id"`
		JobDescription string              `json:"job_description"`
		Phone          string              `json:"phone"`
		SocialLinks    map[string]string   `json:"social_links"`
		Certifications []Certification     `json:"certifications"`
		Projects       []Project           `json:"projects"`
		Services       []OnboardingService `json:"services"`
		Location       OnboardingLocation  `json:"location"`
	}

	OnboardingService struct {
		ID               string               `json:"id" db:"id"`
		UserProfileID    string               `json:"user_profile_id" db:"user_profile_id"`
		Name             string               `json:"name" db:"name"`
		Description      string               `json:"description" db:"description"`
		Price            int                  `json:"price" db:"price"`
		OwnLocationPrice *int                 `json:"own_location_price" db:"own_location_price"`
		CreatedAt        time.Time            `json:"created_at" db:"created_at"`
		UpdatedAt        *time.Time           `json:"updated_at" db:"updated_at"`
		DeletedAt        *time.Time           `json:"deleted_at" db:"deleted_at"`
		Images           []ServiceImageWithID `json:"images" db:"images"`
	}

	OnboardingLocation struct {
		ID            string     `json:"id" db:"id"`
		UserProfileID string     `json:"user_profile_id" db:"user_profile_id"`
		Street        string     `json:"street" db:"street"`
		Number        string     `json:"number" db:"number"`
		Complement    string     `json:"complement" db:"complement"`
		Neighborhood  string     `json:"neighborhood" db:"neighborhood"`
		CreatedAt     time.Time  `json:"created_at" db:"created_at"`
		UpdatedAt     *time.Time `json:"updated_at,omitempty" db:"updated_at"`
		DeletedAt     *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
	}
)
