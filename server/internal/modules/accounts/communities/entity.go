package communities

import (
	"conecta-mare-server/internal/databases/postgres/models"
	"conecta-mare-server/pkg/exceptions"
	"conecta-mare-server/pkg/uid"
	"fmt"
)

type Community struct {
	id      string
	name    string
	censoID int
}

func New(name string, censoID int) (*Community, error) {
	community := Community{
		id:      uid.New("community"),
		name:    name,
		censoID: censoID,
	}

	if err := community.validate(name, censoID); err != nil {
		return nil, exceptions.MakeApiError(err)
	}

	return &community, nil
}

func NewFromModel(m models.Community) *Community {
	return &Community{
		id:      m.ID,
		name:    m.Name,
		censoID: m.CensoID,
	}
}

func (c *Community) ToModel() models.Community {
	return models.Community{
		ID:      c.id,
		Name:    c.name,
		CensoID: c.censoID,
	}
}

func (c *Community) validate(name string, censoID int) error {
	if name == "" || censoID == 0 {
		return fmt.Errorf("name and censoID are required")
	}
	return nil
}

func (c *Community) ID() string {
	return c.id
}

func (c *Community) Name() string {
	return c.name
}

func (c *Community) CensoID() int {
	return c.censoID
}
