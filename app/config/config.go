package config

import (
	"time"

	"go.uber.org/zap"
)

type ConnType string

var (
	ConnTypeTCP ConnType = "tcp"
	ConnTypeUDP ConnType = "udp"
)

type Config struct {
	Log     *zap.Logger
	Server  Server
	Storage Storage
}

type Storage struct {
	PgPort         string        `env:"POSTGRES_PORT" envDefault:"5434"`
	PgHost         string        `env:"POSTGRES_HOST" envDefault:"localhost"`
	PgName         string        `env:"POSTGRES_DBNAME" envDefault:"postgres"`
	PgUser         string        `env:"POSTGRES_USER" envDefault:"postgres"`
	PgPassword     string        `env:"POSTGRES_PASSWORD" envDefault:"qwerty"`
	PgPingEnabled  bool          `env:"POSTGRES_PING_ENABLED" envDefault:"true"`
	PgPingInterval time.Duration `env:"POSTGRES_PING_INTERVAL" envDefault:"40m"`
	PgMaxOpenConn  int           `env:"POSTGRES_MAX_OPEN_CONN" envDefault:"10"`
	PgIdleConn     int           `env:"POSTGRES_MAX_IDLE_CONN" envDefault:"10"`
	PgSSLMode      string        `env:"POSTGRES_SSL_MODE" envDefault:"disable"`
	MigrationPath  string        `env:"MIGRATE_PATH" envDefault:"file://scheme"`
}

type Server struct {
	Addr     string   `env:"SERVER_ADDR" envDefault:":80"`
	ConnType ConnType `env:"SERVER_CONN_TYPE" envDefault:"tcp"`
}
