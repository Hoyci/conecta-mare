package onboardings

import (
	"conecta-mare-server/internal/database/models"
	"conecta-mare-server/pkg/exceptions"
	"conecta-mare-server/pkg/uid"
	"conecta-mare-server/pkg/valueobjects"
	"encoding/json"
	"fmt"
	"time"
)

type UserProfile struct {
	id             string
	userID         string
	fullName       string
	profileImage   string
	jobDescription string
	phone          string
	socialLinks    map[string]string
	createdAt      time.Time
	updatedAt      *time.Time
}

type Certification struct {
	id            string
	userProfileID string
	institution   string
	courseName    string
	startDate     time.Time
	endDate       *time.Time
	createdAt     time.Time
}

type Service struct {
	id            string
	userProfileID string
	name          string
	description   string
	createdAt     time.Time
}
type ServiceImage struct {
	id        string
	serviceID string
	url       string
	ordering  int
	createdAt time.Time
}

func NewUserProfile(
	userID,
	fullName,
	profileImage,
	jobDescription,
	phone string,
	socialLinks map[string]string,
) (*UserProfile, error) {
	userProfile := UserProfile{
		id:             uid.New("userprofile"),
		userID:         userID,
		fullName:       fullName,
		profileImage:   profileImage,
		jobDescription: jobDescription,
		phone:          phone,
		socialLinks:    socialLinks,
		createdAt:      time.Now(),
		updatedAt:      nil,
	}

	if err := userProfile.validate(); err != nil {
		return nil, exceptions.MakeApiError(err)
	}

	return &userProfile, nil
}

func NewUserProfileFromModel(m models.UserProfile) *UserProfile {
	var socialLinks map[string]string
	_ = json.Unmarshal(m.SocialLinks, &socialLinks)

	return &UserProfile{
		id:             m.ID,
		userID:         m.UserID,
		fullName:       m.FullName,
		profileImage:   m.ProfileImage,
		jobDescription: m.JobDescription,
		phone:          m.Phone,
		socialLinks:    socialLinks,
		createdAt:      m.CreatedAt,
		updatedAt:      m.UpdatedAt,
	}
}

func (up *UserProfile) UserProfileToModel() models.UserProfile {
	socialLinksBytes, _ := json.Marshal(up.socialLinks)

	return models.UserProfile{
		ID:             up.id,
		UserID:         up.userID,
		FullName:       up.fullName,
		ProfileImage:   up.profileImage,
		JobDescription: up.jobDescription,
		Phone:          up.phone,
		SocialLinks:    socialLinksBytes,
		CreatedAt:      up.createdAt,
		UpdatedAt:      up.updatedAt,
	}
}

func (up *UserProfile) validate() error {
	if up.userID == "" {
		return fmt.Errorf("user_id is required")
	}

	if up.fullName == "" {
		return fmt.Errorf("full_name is required")
	}

	if up.jobDescription == "" {
		return fmt.Errorf("job_description is required")
	}

	if up.phone == "" {
		return fmt.Errorf("phone is required")
	}

	if _, ok := valueobjects.SanitizePhoneNumber(up.phone); !ok {
		return fmt.Errorf("phone is invalid. use the 219887654321 format")
	}

	return nil
}

func (up *UserProfile) ID() string                     { return up.id }
func (up *UserProfile) UserID() string                 { return up.userID }
func (up *UserProfile) FullName() string               { return up.fullName }
func (up *UserProfile) ProfileImage() string           { return up.profileImage }
func (up *UserProfile) JobDescription() string         { return up.jobDescription }
func (up *UserProfile) Phone() string                  { return up.phone }
func (up *UserProfile) SocialLinks() map[string]string { return up.socialLinks }
func (up *UserProfile) CreatedAt() time.Time           { return up.createdAt }
func (up *UserProfile) UpdatedAt() *time.Time          { return up.updatedAt }

func NewCertification(
	userProfileID,
	institution,
	courseName string,
	startDate time.Time,
	endDate *time.Time,
) (*Certification, error) {
	c := &Certification{
		id:            uid.New("cert"),
		userProfileID: userProfileID,
		institution:   institution,
		courseName:    courseName,
		startDate:     startDate,
		endDate:       endDate,
		createdAt:     time.Now(),
	}

	if err := c.validate(); err != nil {
		return nil, exceptions.MakeApiError(err)
	}

	return c, nil
}

func NewCertificationFromModel(m models.Certification) *Certification {
	return &Certification{
		id:            m.ID,
		userProfileID: m.UserProfileID,
		institution:   m.Institution,
		courseName:    m.CourseName,
		startDate:     m.StartDate,
		endDate:       m.EndDate,
		createdAt:     m.CreatedAt,
	}
}

func (c *Certification) ToModel() models.Certification {
	return models.Certification{
		ID:            c.id,
		UserProfileID: c.userProfileID,
		Institution:   c.institution,
		CourseName:    c.courseName,
		StartDate:     c.startDate,
		EndDate:       c.endDate,
		CreatedAt:     c.createdAt,
	}
}

func (c *Certification) validate() error {
	if c.userProfileID == "" {
		return fmt.Errorf("user_profile_id is required")
	}
	if c.institution == "" {
		return fmt.Errorf("institution is required")
	}
	if c.courseName == "" {
		return fmt.Errorf("course_name is required")
	}
	return nil
}

func (c *Certification) ID() string            { return c.id }
func (c *Certification) UserProfileID() string { return c.userProfileID }
func (c *Certification) Institution() string   { return c.institution }
func (c *Certification) CourseName() string    { return c.courseName }
func (c *Certification) StartDate() time.Time  { return c.startDate }
func (c *Certification) EndDate() *time.Time   { return c.endDate }
func (c *Certification) CreatedAt() time.Time  { return c.createdAt }

func NewServices(
	userProfileID, id, name, description string,
) (*Service, error) {
	s := &Service{
		id:            id,
		userProfileID: userProfileID,
		name:          name,
		description:   description,
		createdAt:     time.Now(),
	}

	if err := s.validate(); err != nil {
		return nil, exceptions.MakeApiError(err)
	}

	return s, nil
}

func NewServiceFromModel(m models.Service) *Service {
	return &Service{
		id:            m.ID,
		userProfileID: m.UserProfileID,
		name:          m.Name,
		description:   m.Description,
		createdAt:     m.CreatedAt,
	}
}

func (s *Service) ToModel() models.Service {
	return models.Service{
		ID:            s.id,
		UserProfileID: s.userProfileID,
		Name:          s.name,
		Description:   s.description,
		CreatedAt:     s.createdAt,
	}
}

func (s *Service) validate() error {
	if s.userProfileID == "" {
		return fmt.Errorf("user_profile_id is required")
	}
	if s.name == "" {
		return fmt.Errorf("name is required")
	}
	if s.description == "" {
		return fmt.Errorf("description is required")
	}
	return nil
}

func (s *Service) ID() string            { return s.id }
func (s *Service) UserProfileID() string { return s.userProfileID }
func (s *Service) Name() string          { return s.name }
func (s *Service) Description() string   { return s.description }
func (s *Service) CreatedAt() time.Time  { return s.createdAt }

func NewServiceImage(
	serviceID, url string,
	ordering int,
) (*ServiceImage, error) {
	img := &ServiceImage{
		id:        uid.New("serviceimg"),
		serviceID: serviceID,
		url:       url,
		ordering:  ordering,
		createdAt: time.Now(),
	}

	if err := img.validate(); err != nil {
		return nil, exceptions.MakeApiError(err)
	}

	return img, nil
}

func NewServiceImageFromModel(m models.ServiceImage) *ServiceImage {
	return &ServiceImage{
		id:        m.ID,
		serviceID: m.ServiceID,
		url:       m.URL,
		ordering:  m.Ordering,
		createdAt: m.CreatedAt,
	}
}

func (img *ServiceImage) ToModel() models.ServiceImage {
	return models.ServiceImage{
		ID:        img.id,
		ServiceID: img.serviceID,
		URL:       img.url,
		Ordering:  img.ordering,
		CreatedAt: img.createdAt,
	}
}

func (img *ServiceImage) validate() error {
	if img.serviceID == "" {
		return fmt.Errorf("service_id is required")
	}
	if img.url == "" {
		return fmt.Errorf("url is required")
	}
	return nil
}

func (si *ServiceImage) ID() string           { return si.id }
func (si *ServiceImage) ServiceID() string    { return si.serviceID }
func (si *ServiceImage) URL() string          { return si.url }
func (si *ServiceImage) Ordering() int        { return si.ordering }
func (si *ServiceImage) CreatedAt() time.Time { return si.createdAt }
