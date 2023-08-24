package repository

import (
	"context"
	"os"

	"github.com/Moon1it/SerbLangBot/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TopicRepo struct {
	Collection *mongo.Collection
}

func InitTopicRepo(db *mongo.Client) *TopicRepo {
	dbName := os.Getenv("DATABASE")
	topicCollection := db.Database(dbName).Collection("Topics")
	return &TopicRepo{Collection: topicCollection}
}

// GetTopicsCount gets the count of all topics in the database.
func (t *TopicRepo) GetTopicsCount() (int64, error) {

	// Get the count of documents that satisfy the filter
	count, err := t.Collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		return 0, err
	}

	return count, nil
}

// GetAllTopics gets all topics from the database.
func (t *TopicRepo) GetAllTopics() ([]models.Topic, error) {

	cursor, err := t.Collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	// Slice to store the results
	var topicSlice []models.Topic

	// Iterate through the query results and add them to the slice
	for cursor.Next(context.TODO()) {
		var topic models.Topic
		if err := cursor.Decode(&topic); err != nil {
			return nil, err
		}
		topicSlice = append(topicSlice, topic)
	}

	return topicSlice, nil
}
