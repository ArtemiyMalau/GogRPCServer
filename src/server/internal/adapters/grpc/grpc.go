package grpc

import (
	"context"
	"grpc_server/proto"
)

func (a Adapter) Fetch(request *proto.FetchRequest, stream proto.ProductsExplorer_FetchServer) error {
	return nil
}

func (a Adapter) List(ctx context.Context, request *proto.ListRequest) (*proto.ListResponse, error) {
	return nil, nil
}
