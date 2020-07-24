package postgres

import (
	"pqredis/config"
	"pqredis/monitoring/logger"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // init PG driver as required by sqlx
)

func New(cfg config.Config, logger logger.Logger) (*sqlx.DB, error) {
	db, err := newDBHandle(cfg.Postgres().String(), cfg.Postgres().MaxPoolSize)
	if err != nil {
		return db, err
	}
	logger.Infof("database connection initialized with pool size %d", cfg.Postgres().MaxPoolSize)
	return db, err
}

func newDBHandle(conStr string, maxPoolSize int) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", conStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxPoolSize)
	return db, nil
}
