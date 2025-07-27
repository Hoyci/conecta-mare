package locations

import (
	"fmt"

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
	fmt.Println(model.CommunityID)
	query := `
		INSERT INTO locations (
				id, user_profile_id, street, number, complement, community_id, created_at
			) VALUES (
				:id, :user_profile_id, :street, :number, :complement, :community_id, :created_at
		)`

	_, err := tx.NamedExec(query, model)
	return err
}
