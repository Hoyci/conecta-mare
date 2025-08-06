package certifications

import (
	"conecta-mare-server/internal/databases/postgres/models"
	"conecta-mare-server/pkg/exceptions"
	"conecta-mare-server/pkg/uid"
	"time"
)

type Certification struct {
	id            string
	userProfileID string
	institution   string
	courseName    string
	startDate     time.Time
	endDate       *time.Time
	createdAt     time.Time
}

func New(
	userProfileID string,
	institution string,
	courseName string,
	startDate time.Time,
	endDate *time.Time,
) (*Certification, error) {
	certification := Certification{
		id:            uid.New("certification"),
		userProfileID: userProfileID,
		institution:   institution,
		courseName:    courseName,
		startDate:     startDate,
		endDate:       endDate,
		createdAt:     time.Now(),
	}

	if err := certification.validate(); err != nil {
		return nil, exceptions.MakeApiError(err)
	}

	return &certification, nil
}

func NewFromModel(m models.Certification) *Certification {
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
	return nil
}

// Getters
func (c *Certification) ID() string            { return c.id }
func (c *Certification) UserProfileID() string { return c.userProfileID }
func (c *Certification) Institution() string   { return c.institution }
func (c *Certification) CourseName() string    { return c.courseName }
func (c *Certification) StartDate() time.Time  { return c.startDate }
func (c *Certification) EndDate() *time.Time   { return c.endDate }
func (c *Certification) CreatedAt() time.Time  { return c.createdAt }
