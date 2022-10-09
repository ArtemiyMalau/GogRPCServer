package main

import (
	"context"
	"grpc_server/config"
	"grpc_server/internal/adapters/db/mongodb"
	"grpc_server/internal/adapters/grpc"
	"grpc_server/internal/application/api"
	"log"
)

func main() {
	cfg := config.GetConfig()
	dbAdapter, err := db.NewAdapter(context.Background(), cfg.MongoDB.Host, cfg.MongoDB.Port,
		cfg.MongoDB.Username, cfg.MongoDB.Password, cfg.MongoDB.Database, cfg.MongoDB.AuthDB,
		cfg.MongoDB.Collections.Product, cfg.MongoDB.Collections.Document,
	)
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
		return
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application)
	grpcAdapter.Run()
}
