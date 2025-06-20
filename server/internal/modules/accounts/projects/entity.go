package projects

import (
	"conecta-mare-server/internal/database/models"
	"conecta-mare-server/pkg/exceptions"
	"time"
)

type Project struct {
	id            string
	userProfileID string
	name          string
	description   string
	createdAt     time.Time
}

func New(
	porfolioID string,
	userProfileID string,
	name string,
	description string,
) (*Project, error) {
	project := Project{
		id:            porfolioID,
		userProfileID: userProfileID,
		name:          name,
		description:   description,
		createdAt:     time.Now(),
	}

	if err := project.validate(); err != nil {
		return nil, exceptions.MakeApiError(err)
	}

	return &project, nil
}

func NewFromModel(m models.Project) *Project {
	return &Project{
		id:            m.ID,
		userProfileID: m.UserProfileID,
		name:          m.Name,
		description:   m.Description,
		createdAt:     m.CreatedAt,
	}
}

func (s *Project) ToModel() models.Project {
	return models.Project{
		ID:            s.id,
		UserProfileID: s.userProfileID,
		Name:          s.name,
		Description:   s.description,
		CreatedAt:     s.createdAt,
	}
}

func (s *Project) validate() error {
	return nil
}

func (s *Project) ID() string            { return s.id }
func (s *Project) UserProfileID() string { return s.userProfileID }
func (s *Project) Name() string          { return s.name }
func (s *Project) Description() string   { return s.description }
func (s *Project) CreatedAt() time.Time  { return s.createdAt }
