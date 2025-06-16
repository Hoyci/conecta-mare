package serviceimages

import (
	"conecta-mare-server/internal/database/models"
	"conecta-mare-server/pkg/exceptions"
	"conecta-mare-server/pkg/uid"
	"fmt"
	"time"
)

type ServiceImage struct {
	id        string
	serviceID string
	url       string
	ordering  int
	createdAt time.Time
}

func New(
	serviceID, url string,
	ordering int,
) (*ServiceImage, error) {
	img := &ServiceImage{
		id:        uid.New("serviceimg"),
		serviceID: serviceID,
		url:       url,
		ordering:  ordering,
		createdAt: time.Now(),
	}

	if err := img.validate(); err != nil {
		return nil, exceptions.MakeApiError(err)
	}

	return img, nil
}

func NewServiceImageFromModel(m models.ServiceImage) *ServiceImage {
	return &ServiceImage{
		id:        m.ID,
		serviceID: m.ServiceID,
		url:       m.URL,
		ordering:  m.Ordering,
		createdAt: m.CreatedAt,
	}
}

func (img *ServiceImage) ToModel() models.ServiceImage {
	return models.ServiceImage{
		ID:        img.id,
		ServiceID: img.serviceID,
		URL:       img.url,
		Ordering:  img.ordering,
		CreatedAt: img.createdAt,
	}
}

func (img *ServiceImage) validate() error {
	if img.serviceID == "" {
		return fmt.Errorf("service_id is required")
	}
	if img.url == "" {
		return fmt.Errorf("url is required")
	}
	return nil
}

func (si *ServiceImage) ID() string           { return si.id }
func (si *ServiceImage) ServiceID() string    { return si.serviceID }
func (si *ServiceImage) URL() string          { return si.url }
func (si *ServiceImage) Ordering() int        { return si.ordering }
func (si *ServiceImage) CreatedAt() time.Time { return si.createdAt }
