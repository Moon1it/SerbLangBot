package mongodb

import (
	"context"
	"fmt"
	"log"

	"github.com/Moon1it/SerbLangBot/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongoDB(cfg *config.Config) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(cfg.MongoString)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Check if the connection is successful
	err = client.Ping(context.Background(), nil)
	if err != nil {
		_ = client.Disconnect(context.Background()) // Attempt to disconnect on failure
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	log.Println("Successfully connected to MongoDB!")

	return client.Database(cfg.DBName), nil
}
