package subcategories

import (
	"conecta-mare-server/internal/databases/postgres/models"
	"conecta-mare-server/pkg/exceptions"
	"conecta-mare-server/pkg/uid"
	"fmt"
	"time"
)

type Subcategory struct {
	id          string
	name        string
	category_id string
	createdAt   time.Time
	updatedAt   *time.Time
	deletedAt   *time.Time
}

func New(
	name,
	categoryID string,
) (*Subcategory, error) {
	subcategory := &Subcategory{
		id:          uid.New("subcategory"),
		name:        name,
		category_id: categoryID,
		createdAt:   time.Now(),
		updatedAt:   nil,
		deletedAt:   nil,
	}

	if err := subcategory.validate(); err != nil {
		return nil, exceptions.MakeApiError(err)
	}

	return subcategory, nil
}

func NewFromModel(m models.Subcategory) *Subcategory {
	return &Subcategory{
		id:          m.ID,
		name:        m.Name,
		category_id: m.CategoryID,
		createdAt:   m.CreatedAt,
		updatedAt:   m.UpdatedAt,
		deletedAt:   m.DeletedAt,
	}
}

func (s *Subcategory) ToModel() models.Subcategory {
	return models.Subcategory{
		ID:         s.id,
		Name:       s.name,
		CategoryID: s.category_id,
		CreatedAt:  s.createdAt,
		UpdatedAt:  s.updatedAt,
		DeletedAt:  s.deletedAt,
	}
}

func (s *Subcategory) validate() error {
	if s.name == "" {
		return fmt.Errorf("name is required")
	}
	if s.category_id == "" {
		return fmt.Errorf("category_id is required")
	}

	return nil
}

func FromID(id string) *Subcategory {
	return &Subcategory{id: id}
}

func (s *Subcategory) ID() string {
	return s.id
}

func (s *Subcategory) Name() string {
	return s.name
}

func (s *Subcategory) CategoryID() string {
	return s.category_id
}

func (s *Subcategory) CreatedAt() time.Time {
	return s.createdAt
}

func (s *Subcategory) DeletedAt() *time.Time {
	return s.deletedAt
}
