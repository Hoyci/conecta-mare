package locations

import "github.com/jmoiron/sqlx"

type LocationsRepository interface {
	CreateTx(tx *sqlx.Tx, location *Location) error
}
