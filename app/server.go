package app

import (
	"context"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/khusainnov/newbee/app/config"
	"github.com/khusainnov/newbee/app/processors"
	"go.uber.org/zap"
)

func Run(ctx context.Context, cfg *config.Config) {
	lis, err := net.Listen(config.TCPServer, cfg.ServerCfg.Addr)
	if err != nil {
		cfg.Log.Fatal("failed to listen addr", zap.String("Addr", cfg.ServerCfg.Addr))
	}

	proceesor := processors.New()

	if err = rpc.Register(proceesor); err != nil {
		cfg.Log.Fatal("failed to register processor", zap.Error(err))
	}

	cfg.Log.Info("running rpc server", zap.String("Port", cfg.ServerCfg.Addr))
	for {
		conn, err := lis.Accept()
		if err != nil {
			cfg.Log.Fatal("failed to accept connections")
		}

		go func(conn net.Conn) {
			defer conn.Close()
			jsonrpc.ServeConn(conn)
		}(conn)
	}
}
