package serviceimages

import (
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) ServiceImagesRepository {
	return &repository{db}
}

func (r *repository) CreateTx(tx *sqlx.Tx, serviceImg *ServiceImage) error {
	model := serviceImg.ToModel()
	query := `
		INSERT INTO service_images (
				id, service_id, url, ordering, created_at
			) VALUES (
				:id, :service_id, :url, :ordering, :created_at
		)
	`

	_, err := tx.NamedExec(query, model)
	return err
}
