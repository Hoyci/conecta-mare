package certifications

import "github.com/jmoiron/sqlx"

type CertificationsRepository interface {
	CreateTx(tx *sqlx.Tx, certification *Certification) error
	// FindByID(id string) (*Certification, error)
	// FindByUserProfileID(userProfileID string) ([]*Certification, error)
	// Update(certification *Certification) error
	// Delete(id string) error
}

//
// type Service interface {
// 	Create(userProfileID, institution, courseName string, startDate time.Time, endDate *time.Time) (*Certification, error)
// 	GetByID(id string) (*Certification, error)
// 	GetByUserProfileID(userProfileID string) ([]*Certification, error)
// 	Update(id, institution, courseName string, startDate time.Time, endDate *time.Time) (*Certification, error)
// 	Delete(id string) error
// }
