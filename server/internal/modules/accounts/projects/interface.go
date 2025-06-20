package projects

import "github.com/jmoiron/sqlx"

type ProjectsRepository interface {
	CreateTx(tx *sqlx.Tx, project *Project) error
}
