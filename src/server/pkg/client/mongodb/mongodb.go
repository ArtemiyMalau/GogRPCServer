package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/url"
)

func NewClient(ctx context.Context, host, port, username, password, database, authDB string) (*mongo.Database, error) {
	dbUrl := url.URL{
		Scheme: "mongodb",
		Host:   fmt.Sprintf("%s:%s", host, port),
	}

	if username != "" && password != "" {
		dbUrl.User = url.UserPassword(username, password)
	}

	getArgs := dbUrl.Query()
	if authDB != "" {
		getArgs.Add("authSource", authDB)
	}
	dbUrl.RawQuery = getArgs.Encode()

	log.Printf("MongoDB connection url: %s", dbUrl.String())

	// Client options
	clientOptions := options.Client().ApplyURI(dbUrl.String())

	// Connection
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongoDB due the error: %v", err)
	}

	// Pinging
	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping mongoDB due the error: %v", err)
	}

	return client.Database(database), nil
}
