package models

import (
	"time"
)

type UserProfile struct {
	ID             string     `db:"id"`
	UserID         string     `db:"user_id"`
	FullName       string     `db:"full_name"`
	SubcategoryID  *string    `db:"subcategory_id"`
	ProfileImage   *string    `db:"profile_image"`
	JobDescription *string    `db:"job_description"`
	Phone          *string    `db:"phone"`
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

type Project struct {
	ID            string    `db:"id"`
	UserProfileID string    `db:"user_profile_id"`
	Name          string    `db:"name"`
	Description   string    `db:"description"`
	CreatedAt     time.Time `db:"created_at"`
}

type ProjectImage struct {
	ID        string    `db:"id"`
	ProjectID string    `db:"project_id"`
	URL       string    `db:"url"`
	Ordering  int       `db:"ordering"`
	CreatedAt time.Time `db:"created_at"`
}

type Service struct {
	ID               string     `db:"id"`
	UserProfileID    string     `db:"user_profile_id"`
	Name             string     `db:"name"`
	Description      string     `db:"description"`
	Price            int        `db:"price"`
	OwnLocationPrice *int       `db:"own_location_price"`
	CreatedAt        time.Time  `db:"created_at"`
	UpdatedAt        *time.Time `db:"updated_at"`
	DeletedAt        *time.Time `db:"deleted_at"`
}

type ServiceImage struct {
	ID        string    `db:"id"`
	ServiceID string    `db:"service_id"`
	URL       string    `db:"url"`
	Ordering  int       `db:"ordering"`
	CreatedAt time.Time `db:"created_at"`
}

type Location struct {
	ID            string     `db:"id"`
	UserProfileID string     `db:"user_profile_id"`
	Street        string     `db:"street"`
	Number        string     `db:"number"`
	Complement    string     `db:"complement"`
	Neighborhood  string     `db:"neighborhood"`
	CreatedAt     time.Time  `db:"created_at"`
	UpdatedAt     *time.Time `db:"updated_at"`
	DeletedAt     *time.Time `db:"deleted_at"`
}
