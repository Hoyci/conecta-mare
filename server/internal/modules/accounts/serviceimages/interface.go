package serviceimages

import "github.com/jmoiron/sqlx"

type ServiceImagesRepository interface {
	CreateTx(tx *sqlx.Tx, serviceImg *ServiceImage) error
}
