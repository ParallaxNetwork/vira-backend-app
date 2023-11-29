package utils

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongodbConnect() (*mongo.Database, context.Context, error) {
	var ctx = context.Background()
	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
			return nil, ctx, err
	}

	return client.Database("vira"), ctx, nil
}