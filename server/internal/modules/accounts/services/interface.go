package services

import "github.com/jmoiron/sqlx"

type ServicesRepository interface {
	CreateTx(tx *sqlx.Tx, service *Service) error
}
