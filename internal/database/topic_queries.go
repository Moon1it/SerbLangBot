package database

import (
	"context"

	"github.com/Moon1it/SerbLangBot/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetTopicsCount(db *mongo.Database) (int64, error) {
	// Get the Topics collection
	topics := db.Collection("Topics")

	// Create a filter to search for all topics
	filter := bson.M{} // An empty filter will return all documents in the collection

	// Get the count of documents that satisfy the filter
	count, err := topics.CountDocuments(context.TODO(), filter)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func GetAllTopics(db *mongo.Database) ([]models.Topic, error) {
	// Get the Topics collection
	topics := db.Collection("Topics")

	// Create a filter to search for all topics
	var filter bson.M // a nil filter will return all documents in the collection

	cursor, err := topics.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	// Slice to store the results
	var result []models.Topic

	// Iterate through the query results and add them to the slice
	for cursor.Next(context.TODO()) {
		var topic models.Topic
		if err := cursor.Decode(&topic); err != nil {
			return nil, err
		}
		result = append(result, topic)
	}

	return result, nil
}
