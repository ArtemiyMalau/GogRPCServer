package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"grpc_server/internal/ports"
	"grpc_server/pkg/client/mongodb"
)

type Adapter struct {
	productC  *mongo.Collection
	documentC *mongo.Collection
}

var _ ports.DBPort = (*Adapter)(nil)

func NewAdapter(ctx context.Context, host, port, username, password, database, authDB string, productC, documentC string) (*Adapter, error) {
	client, err := mongodb.NewClient(ctx, host, port, username, password, database, authDB)
	if err != nil {
		return nil, err
	}

	return &Adapter{
		productC:  client.Collection(productC),
		documentC: client.Collection(documentC),
	}, nil
}
