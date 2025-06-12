package common

import "time"

type (
	OnboardingRequest struct {
		UserID         string              `json:"user_id"`
		FullName       string              `json:"full_name"`
		JobDescription string              `json:"job_description"`
		Phone          string              `json:"phone"`
		SocialLinks    map[string]string   `json:"social_links"`
		Certifications []Certification     `json:"certifications"`
		Services       []ServiceWithImages `json:"services"`
	}

	Certification struct {
		Institution string    `json:"institution"`
		CourseName  string    `json:"course_name"`
		StartDate   time.Time `json:"start_date"`
		EndDate     time.Time `json:"end_date"`
	}

	ServiceWithImages struct {
		ID          string   `json:"-"`
		Name        string   `json:"name"`
		Description string   `json:"description"`
		Images      []string `json:"images"`
	}
)
