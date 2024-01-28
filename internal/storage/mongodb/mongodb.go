package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewConnection(cfg Config) (*mongo.Database, error) {
	ctx := context.Background()

	// mongodb server options
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)

	clientOptions := options.Client().
		ApplyURI(cfg.URI).
		SetServerAPIOptions(serverAPIOptions)

	// creating mongodb client
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("mongoDB connection failed: %w", err)
	}

	// ping mongodb
	if er := client.Ping(ctx, nil); er != nil {
		return nil, fmt.Errorf("mongoDB ping failed: %w", er)
	}

	return client.Database(cfg.Database), nil
}
