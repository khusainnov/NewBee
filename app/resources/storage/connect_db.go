package storage

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/khusainnov/newbee/app/config"
	"gitlab.com/khusainnov/driver/postgres"
	"go.uber.org/zap"
)

type ClientImpl struct {
	db *sqlx.DB
}

// TODO: implement connect to Postgres DB
func New(log *zap.Logger, cfg config.Storage) (*ClientImpl, error) {
	db, err := postgres.NewPostgresDB(
		postgres.ConfigPG{
			Host:         cfg.PgHost,
			Port:         cfg.PgPort,
			User:         cfg.PgUser,
			Password:     cfg.PgPassword,
			DBName:       cfg.PgName,
			SSLMode:      cfg.PgSSLMode,
			MaxOpenConns: cfg.PgMaxOpenConn,
			MaxIdleConns: cfg.PgIdleConn,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database, %w", err)
	}

	go func() {
		t := time.NewTicker(cfg.PgPingInterval)

		for range t.C {
			if err := db.Ping(); err != nil {
				log.Warn("failed to ping db", zap.Error(err))
			}
		}
	}()

	return &ClientImpl{db: db}, nil
}

func (db *ClientImpl) GetDB() *sqlx.DB {
	return db.db
}

// TODO: add command to run migration separate of main code
