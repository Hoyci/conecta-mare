package clickhouse

import (
	"fmt"
	"log"
	"log/slog"

	_ "github.com/ClickHouse/clickhouse-go/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
)

type ClickHouse struct {
	db *sqlx.DB
}

func New(username, password, host, port, database string) *ClickHouse {
	dsn := fmt.Sprintf("tcp://%s:%s@%s:%s?database=%s", username, password, host, port, database)

	db, err := sqlx.Connect("clickhouse", dsn)
	if err != nil {
		log.Fatalf("Error connecting to ClickHouse: %v", err)
		return &ClickHouse{}
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging ClickHouse: %v", err)
	}

	return &ClickHouse{db: db}
}

func (c *ClickHouse) Close() error {
	slog.Info("Disconnected from ClickHouse")
	return c.db.Close()
}

func (c *ClickHouse) DB() *sqlx.DB {
	return c.db
}
