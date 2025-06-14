package services

import (
	"conecta-mare-server/internal/database/models"
	"conecta-mare-server/pkg/exceptions"
	"time"
)

type Service struct {
	id            string
	userProfileID string
	name          string
	description   string
	createdAt     time.Time
}

func New(
	serviceID string,
	userProfileID string,
	name string,
	description string,
) (*Service, error) {
	service := Service{
		id:            serviceID,
		userProfileID: userProfileID,
		name:          name,
		description:   description,
		createdAt:     time.Now(),
	}

	if err := service.validate(); err != nil {
		return nil, exceptions.MakeApiError(err)
	}

	return &service, nil
}

func NewFromModel(m models.Service) *Service {
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
	return nil
}

func (s *Service) ID() string            { return s.id }
func (s *Service) UserProfileID() string { return s.userProfileID }
func (s *Service) Name() string          { return s.name }
func (s *Service) Description() string   { return s.description }
func (s *Service) CreatedAt() time.Time  { return s.createdAt }
