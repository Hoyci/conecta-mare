package projectimages

import (
	"conecta-mare-server/internal/database/models"
	"conecta-mare-server/pkg/exceptions"
	"fmt"
	"time"
)

type ProjectImage struct {
	id        string
	projectID string
	url       string
	ordering  int
	createdAt time.Time
}

func New(id, projectID, url string, ordering int) (*ProjectImage, error) {
	img := &ProjectImage{
		id:        id,
		projectID: projectID,
		url:       url,
		ordering:  ordering,
		createdAt: time.Now(),
	}

	if err := img.validate(); err != nil {
		return nil, exceptions.MakeApiError(err)
	}

	return img, nil
}

func NewFromModel(m models.ProjectImage) *ProjectImage {
	return &ProjectImage{
		id:        m.ID,
		projectID: m.ProjectID,
		url:       m.URL,
		ordering:  m.Ordering,
		createdAt: m.CreatedAt,
	}
}

func (img *ProjectImage) ToModel() models.ProjectImage {
	return models.ProjectImage{
		ID:        img.id,
		ProjectID: img.projectID,
		URL:       img.url,
		Ordering:  img.ordering,
		CreatedAt: img.createdAt,
	}
}

func (img *ProjectImage) validate() error {
	if img.projectID == "" {
		return fmt.Errorf("project_id is required")
	}
	if img.url == "" {
		return fmt.Errorf("url is required")
	}
	return nil
}

func (si *ProjectImage) ID() string           { return si.id }
func (si *ProjectImage) ProjectID() string    { return si.projectID }
func (si *ProjectImage) URL() string          { return si.url }
func (si *ProjectImage) Ordering() int        { return si.ordering }
func (si *ProjectImage) CreatedAt() time.Time { return si.createdAt }
