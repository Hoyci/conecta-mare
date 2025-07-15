package services

import (
	"conecta-mare-server/internal/database/models"
	"time"
)

type Service struct {
	id               string
	userProfileID    string
	name             string
	description      string
	price            int
	ownLocationPrice *int
	createdAt        time.Time
	updatedAt        *time.Time
	deletedAt        *time.Time
}

func New(
	id,
	userProfileID,
	name,
	description string,
	price int,
	ownLocationPrice *int,
) (*Service, error) {
	service := Service{
		id:               id,
		userProfileID:    userProfileID,
		name:             name,
		description:      description,
		price:            price,
		ownLocationPrice: ownLocationPrice,
		createdAt:        time.Now(),
		updatedAt:        nil,
		deletedAt:        nil,
	}

	return &service, nil
}

func NewFromModel(m models.Service) *Service {
	return &Service{
		id:               m.ID,
		userProfileID:    m.UserProfileID,
		name:             m.Name,
		description:      m.Description,
		price:            m.Price,
		ownLocationPrice: m.OwnLocationPrice,
		createdAt:        m.CreatedAt,
		updatedAt:        m.UpdatedAt,
		deletedAt:        m.DeletedAt,
	}
}

func (s *Service) ToModel() models.Service {
	return models.Service{
		ID:               s.id,
		UserProfileID:    s.userProfileID,
		Name:             s.name,
		Description:      s.description,
		Price:            s.price,
		OwnLocationPrice: s.ownLocationPrice,
		CreatedAt:        s.createdAt,
		UpdatedAt:        s.updatedAt,
		DeletedAt:        s.deletedAt,
	}
}

func (s *Service) ID() string             { return s.id }
func (s *Service) UserProfileID() string  { return s.userProfileID }
func (s *Service) Name() string           { return s.name }
func (s *Service) Description() string    { return s.description }
func (s *Service) Price() int             { return s.price }
func (s *Service) OwnLocationPrice() *int { return s.ownLocationPrice }
