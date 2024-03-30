package config

import "go.uber.org/zap"

const TCPServer = "tcp"

type Config struct {
	Log       *zap.Logger
	ServerCfg Server
}

type Server struct {
	Addr string
}
