package projectimages

import (
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) ProjectImagesRepository {
	return &repository{db}
}

func (r *repository) CreateTx(tx *sqlx.Tx, projectImg *ProjectImage) error {
	model := projectImg.ToModel()
	query := `
		INSERT INTO project_images (
				id, project_id, url, ordering, created_at
			) VALUES (
				:id, :project_id, :url, :ordering, :created_at
		)
	`

	_, err := tx.NamedExec(query, model)
	return err
}
