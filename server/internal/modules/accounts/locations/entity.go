package locations

import (
	"conecta-mare-server/internal/database/models"
	"conecta-mare-server/pkg/uid"
	"time"
)

type Location struct {
	id            string
	userProfileID string
	street        string
	number        string
	complement    string
	communityID   string
	createdAt     time.Time
	updatedAt     *time.Time
	deletedAt     *time.Time
}

func New(
	userProfileID,
	street,
	number,
	complement,
	communityID string,
) (*Location, error) {
	service := Location{
		id:            uid.New("location"),
		userProfileID: userProfileID,
		street:        street,
		number:        number,
		complement:    complement,
		communityID:   communityID,
		createdAt:     time.Now(),
		updatedAt:     nil,
		deletedAt:     nil,
	}

	return &service, nil
}

func NewFromModel(m models.Location) *Location {
	return &Location{
		id:            m.ID,
		userProfileID: m.UserProfileID,
		street:        m.Street,
		number:        m.Number,
		complement:    m.Complement,
		communityID:   m.CommunityID,
		createdAt:     m.CreatedAt,
		updatedAt:     m.UpdatedAt,
		deletedAt:     m.DeletedAt,
	}
}

func (s *Location) ToModel() models.Location {
	return models.Location{
		ID:            s.id,
		UserProfileID: s.userProfileID,
		Street:        s.street,
		Number:        s.number,
		Complement:    s.complement,
		CommunityID:   s.communityID,
		CreatedAt:     s.createdAt,
		UpdatedAt:     s.updatedAt,
		DeletedAt:     s.deletedAt,
	}
}

func (s *Location) ID() string            { return s.id }
func (s *Location) UserProfileID() string { return s.userProfileID }
func (s *Location) Street() string        { return s.street }
func (s *Location) Number() string        { return s.number }
func (s *Location) Complement() string    { return s.complement }
func (s *Location) CommunityID() string   { return s.communityID }
