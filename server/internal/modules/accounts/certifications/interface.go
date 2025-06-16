package certifications

import "github.com/jmoiron/sqlx"

type CertificationsRepository interface {
	CreateTx(tx *sqlx.Tx, certification *Certification) error
}
