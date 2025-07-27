package models

type Community struct {
	ID      string `db:"id"`
	Name    string `db:"name"`
	CensoID int    `db:"censo_id"`
}
