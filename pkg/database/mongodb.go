package database

import (
	"context"
	"errors"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client
var dbName string

func GetCollection(name string) *mongo.Collection {
	return mongoClient.Database(dbName).Collection(name)
}

func StartMongoDB() error {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		return errors.New("MONGODB_URI environment variable not set")
	}

	database := os.Getenv("DATABASE")
	if database == "" {
		return errors.New("DATABASE environment variable not set")
	} else {
		dbName = database
	}

	var err error
	mongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	err = mongoClient.Ping(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	return nil
}

func CloseMongoDB() error {
	err := mongoClient.Disconnect(context.Background())
	if err != nil {
		return fmt.Errorf("failed to close MongoDB connection: %w", err)
	}
	return nil
}
