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
	dsn := fmt.Sprintf("clickhouse://%s:%s@%s:%s/%s", username, password, host, port, database)

	db, err := sqlx.Connect("clickhouse", dsn)
	if err != nil {
		log.Fatalf("Erro ao conectar no ClickHouse: %v", err)
		return &ClickHouse{}
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Erro ao pingar ClickHouse: %v", err)
	}

	slog.Info("Conectado ao ClickHouse")

	return &ClickHouse{db: db}
}

func (c *ClickHouse) Close() error {
	slog.Info("Desconectado do ClickHouse")
	return c.db.Close()
}

func (c *ClickHouse) DB() *sqlx.DB {
	return c.db
}
