package userprofiles

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

func New(
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

func NewFromModel(m models.UserProfile) *UserProfile {
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
