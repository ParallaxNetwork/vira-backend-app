package utils

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongodbConnect() (*mongo.Database, context.Context, error) {
	err := godotenv.Load(".env")
	mongourl := os.Getenv("MONGODB_URL")

	if err != nil {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}

	var ctx = context.Background()
	clientOptions := options.Client()
	clientOptions.ApplyURI(mongourl)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
			return nil, ctx, err
	}

	return client.Database("vira"), ctx, nil
}