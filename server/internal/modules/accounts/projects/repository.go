package projects

import (
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) ProjectsRepository {
	return &repository{db}
}

func (r *repository) CreateTx(tx *sqlx.Tx, project *Project) error {
	model := project.ToModel()
	query := `
		INSERT INTO projects (
				id, user_profile_id, name, description, created_at
			) VALUES (
				:id, :user_profile_id, :name, :description, :created_at
		)`

	_, err := tx.NamedExec(query, model)
	return err
}
