package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc_server/config"
	"grpc_server/internal/ports"
	"grpc_server/proto"
	"log"
	"net"
)

type Adapter struct {
	proto.UnimplementedProductsExplorerServer
	api ports.APIPort
}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{api: api}
}

func (a Adapter) Run() {
	cfg := config.GetConfig()

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.IP, cfg.Listen.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterProductsExplorerServer(grpcServer, a)
	log.Printf("Server started at %s:%s\n", cfg.Listen.IP, cfg.Listen.Port)
	if cfg.Environment == "development" {
		reflection.Register(grpcServer)
	}
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve grpc server due to error %v", err)
	}
}
