package app

import (
	"net/http"

	"github.com/caarlos0/env/v6"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
	"github.com/khusainnov/newbee/app/api"
	"github.com/khusainnov/newbee/app/config"
	"github.com/khusainnov/newbee/app/resources/storage"
	"go.uber.org/zap"
)

func Run() {
	cfg := newConfig()
	log := cfg.Log

	client, err := storage.New(log, cfg.Storage)
	if err != nil {
		log.Fatal("failed to init storage", zap.Error(err))
	}

	apis := api.NewAPI(log, client.GetDB())

	router := mux.NewRouter()

	rpcServer := rpc.NewServer()
	rpcServer.RegisterCodec(json2.NewCodec(), "application/json")

	if err := rpcServer.RegisterService(apis, ""); err != nil {
		log.Fatal("failed to register service", zap.Error(err))
	}

	router.Handle("/jsonrpc/v2", rpcServer)

	log.Info("Starting server on", zap.String("PORT", cfg.Server.Addr))
	if err := http.ListenAndServe(cfg.Server.Addr, router); err != nil {
		log.Fatal("Error starting server", zap.Error(err))
	}
}

func newConfig() *config.Config {
	cfg := &config.Config{}

	log, _ := zap.NewDevelopment()
	cfg.Log = log
	if err := env.Parse(cfg); err != nil {
		cfg.Log.Fatal("cannot parse config", zap.Error(err))
	}

	return cfg
}
