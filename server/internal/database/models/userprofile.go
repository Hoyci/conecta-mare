package models

import (
	"time"
)

type UserProfile struct {
	ID             string     `db:"id"`
	UserID         string     `db:"user_id"`
	FullName       string     `db:"full_name"`
	CategoryID     string     `db:"category_id"`
	SubcategoryID  string     `db:"subcategory_id"`
	ProfileImage   string     `db:"profile_image"`
	JobDescription string     `db:"job_description"`
	Phone          string     `db:"phone"`
	SocialLinks    []byte     `db:"social_links"`
	CreatedAt      time.Time  `db:"created_at"`
	UpdatedAt      *time.Time `db:"updated_at"`
}

type Certification struct {
	ID            string     `db:"id"`
	UserProfileID string     `db:"user_profile_id"`
	Institution   string     `db:"institution"`
	CourseName    string     `db:"course_name"`
	StartDate     time.Time  `db:"start_date"`
	EndDate       *time.Time `db:"end_date"`
	CreatedAt     time.Time  `db:"created_at"`
}

type Service struct {
	ID            string    `db:"id"`
	UserProfileID string    `db:"user_profile_id"`
	Name          string    `db:"name"`
	Description   string    `db:"description"`
	CreatedAt     time.Time `db:"created_at"`
}

type ServiceImage struct {
	ID        string    `db:"id"`
	ServiceID string    `db:"service_id"`
	URL       string    `db:"url"`
	Ordering  int       `db:"ordering"`
	CreatedAt time.Time `db:"created_at"`
}
