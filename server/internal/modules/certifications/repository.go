package certifications

import (
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) CertificationsRepository {
	return &repository{db}
}

func (r *repository) CreateTx(tx *sqlx.Tx, certification *Certification) error {
	model := certification.ToModel()
	_, err := tx.NamedExec(`
			INSERT INTO certifications (
				id, user_profile_id, institution, course_name, start_date, end_date, created_at
			) VALUES (
				:id, :user_profile_id, :institution, :course_name, :start_date, :end_date, :created_at
			)
		`,
		&model,
	)
	return err
}
