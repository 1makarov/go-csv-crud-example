package postgres

import (
	"github.com/1makarov/go-csv-crud-example/internal/db"
	"github.com/jmoiron/sqlx"
)

func Open(cfg db.ConfigDB) (*sqlx.DB, error) {
	return sqlx.Connect("pgx", cfg.String())
}
