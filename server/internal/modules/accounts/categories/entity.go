package categories

import (
	"conecta-mare-server/internal/database/models"
	"conecta-mare-server/pkg/exceptions"
	"conecta-mare-server/pkg/uid"
	"fmt"
	"time"
)

type Category struct {
	id        string
	name      string
	icon      string
	createdAt time.Time
	updatedAt *time.Time
	deletedAt *time.Time
}

func New(name, icon string) (*Category, error) {
	category := Category{
		id:        uid.New("cat"),
		name:      name,
		icon:      icon,
		createdAt: time.Now(),
		updatedAt: nil,
	}

	if err := category.validate(name, icon); err != nil {
		return nil, exceptions.MakeApiError(err)
	}

	return &category, nil
}

func NewFromModel(m models.Category) *Category {
	return &Category{
		id:        m.ID,
		name:      m.Name,
		icon:      m.Icon,
		createdAt: m.CreatedAt,
		updatedAt: m.UpdatedAt,
		deletedAt: m.DeletedAt,
	}
}

func (c *Category) ToModel() models.Category {
	return models.Category{
		ID:        c.id,
		Name:      c.name,
		Icon:      c.icon,
		CreatedAt: c.createdAt,
		UpdatedAt: c.updatedAt,
		DeletedAt: c.deletedAt,
	}
}

func (c *Category) validate(name, icon string) error {
	if name == "" || icon == "" {
		return fmt.Errorf("name and icon are required")
	}
	return nil
}

func FromID(id string) *Category {
	return &Category{id: id}
}

func (c *Category) ID() string {
	return c.id
}

func (c *Category) Name() string {
	return c.name
}

func (c *Category) Icon() string {
	return c.icon
}

func (c *Category) CreatedAt() time.Time {
	return c.createdAt
}

func (c *Category) UpdatedAt() *time.Time {
	return c.updatedAt
}

func (c *Category) DeletedAt() *time.Time {
	return c.deletedAt
}
