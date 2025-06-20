package projectimages

import "github.com/jmoiron/sqlx"

type ProjectImagesRepository interface {
	CreateTx(tx *sqlx.Tx, projectImg *ProjectImage) error
}
