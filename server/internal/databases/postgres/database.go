package postgres

import (
	"fmt"
	"log"
	"log/slog"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
)

type Postgres struct {
	db *sqlx.DB
}

func New(username, password, host, port, database string) *Postgres {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, database)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal(err)
		return &Postgres{}
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return &Postgres{db: db}
}

func (d *Postgres) Close() error {
	slog.Info("Disconnected from database")
	return d.db.Close()
}

func (d *Postgres) DB() *sqlx.DB {
	return d.db
}
