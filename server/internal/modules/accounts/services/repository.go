package services

import (
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) ServicesRepository {
	return &repository{db}
}

func (r *repository) CreateTx(tx *sqlx.Tx, service *Service) error {
	model := service.ToModel()
	query := `
		INSERT INTO services (
				id, user_profile_id, name, description, created_at
			) VALUES (
				:id, :user_profile_id, :name, :description, :created_at
		)`

	_, err := tx.NamedExec(query, model)
	return err
}
