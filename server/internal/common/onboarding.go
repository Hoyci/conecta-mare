package common

type (
	OnboardingRequest struct {
		UserID         string            `json:"user_id"`
		FullName       string            `json:"full_name"`
		SubcategoryID  string            `json:"subcategory_id"`
		JobDescription string            `json:"job_description"`
		Phone          string            `json:"phone"`
		SocialLinks    map[string]string `json:"social_links"`
		Certifications []Certification   `json:"certifications"`
		Projects       []Project         `json:"projects"`
		Services       []Service         `json:"services"`
		Location       Location          `json:"location"`
	}
)
