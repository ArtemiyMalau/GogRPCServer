package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc_client/pb"
	"log"
)

func NewGRPCClient(host, port string) (pb.ProductsExplorerClient, chan any) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	endConn := make(chan any)
	go func() {
		<-endConn
		
		err := conn.Close()
		if err != nil {
			return
		}
	}()
	return pb.NewProductsExplorerClient(conn), endConn
}
