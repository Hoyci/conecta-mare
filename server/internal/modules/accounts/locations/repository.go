package locations

import (
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) LocationsRepository {
	return &repository{db}
}

func (r *repository) CreateTx(tx *sqlx.Tx, location *Location) error {
	model := location.ToModel()
	query := `
		INSERT INTO locations (
				id, user_profile_id, street, number, complement, neighborhood, created_at
			) VALUES (
				:id, :user_profile_id, :street, :number, :complement, :neighborhood, :created_at
		)`

	_, err := tx.NamedExec(query, model)
	return err
}
