package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	Client *mongo.Client
}

func ConnectToMongoDB() (*MongoClient, error) {

	uri := os.Getenv("MONGODB_URI")

	opts := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Check if the connection is successful
	if err := client.Ping(context.TODO(), nil); err != nil {
		_ = client.Disconnect(context.TODO()) // Attempt to disconnect on failure
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	log.Println("Successfully connected to MongoDB!")

	return &MongoClient{Client: client}, nil
}

func (m *MongoClient) CloseMongoDB() error {
	err := m.Client.Disconnect(context.Background())
	if err != nil {
		return fmt.Errorf("failed to close MongoDB connection: %w", err)
	}
	return nil
}
