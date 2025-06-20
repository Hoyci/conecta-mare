package common

type (
	OnboardingRequest struct {
		UserID         string            `json:"user_id"`
		FullName       string            `json:"full_name"`
		CategoryID     string            `json:"category_id"`
		SubcategoryID  string            `json:"subcategory_id"`
		JobDescription string            `json:"job_description"`
		Phone          string            `json:"phone"`
		SocialLinks    map[string]string `json:"social_links"`
		Certifications []Certification   `json:"certifications"`
		Projects       []Project         `json:"projects"`
	}

	// Certification struct {
	// 	Institution string    `json:"institution"`
	// 	CourseName  string    `json:"course_name"`
	// 	StartDate   time.Time `json:"start_date"`
	// 	EndDate     time.Time `json:"end_date"`
	// }
)
