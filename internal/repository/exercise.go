package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/Moon1it/SerbLangBot/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ExerciseRepo struct {
	Collection *mongo.Collection
}

func InitExerciseRepo(db *mongo.Client) *ExerciseRepo {
	dbName := os.Getenv("DATABASE")
	exerciseCollection := db.Database(dbName).Collection("Exercises")
	return &ExerciseRepo{Collection: exerciseCollection}
}

// GetExerciseByTopic retrieves a random exercise for a given topic name from the database.
// Returns the found exercise or an error if not found.
func (e *ExerciseRepo) GetByTopic(topicName string) (*models.Exercise, error) {
	pipeline := bson.A{
		bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "Topics"},
				{Key: "localField", Value: "topicID"},
				{Key: "foreignField", Value: "topicID"},
				{Key: "as", Value: "Topic"},
			}},
		},
		bson.D{
			{Key: "$match", Value: bson.M{
				"Topic.name": topicName,
			}},
		},
		bson.D{
			{Key: "$sample", Value: bson.M{"size": 1}},
		},
		bson.D{
			{Key: "$project", Value: bson.M{
				"_id":   0,
				"Topic": 0,
			}},
		},
	}

	cursor, err := e.Collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, fmt.Errorf("failed to execute aggregation: %w", err)
	}
	defer cursor.Close(context.Background())

	if !cursor.Next(context.Background()) {
		return nil, fmt.Errorf("no exercises found for topic: %s", topicName)
	}

	var exercise models.Exercise
	if err := cursor.Decode(&exercise); err != nil {
		return nil, fmt.Errorf("failed to decode exercise: %w", err)
	}

	return &exercise, nil
}

// GetAnswerByQuestion retrieves an exercise by its question from the database.
// Returns the found exercise or an error if not found.
func (e *ExerciseRepo) GetAnswerByQuestion(question string) (*models.Exercise, error) {
	pipeline := bson.A{
		bson.D{
			{Key: "$match", Value: bson.M{"question": question}},
		},
		bson.D{
			{Key: "$sample", Value: bson.M{"size": 1}},
		},
	}

	cursor, err := e.Collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, fmt.Errorf("failed to execute aggregation: %w", err)
	}
	defer cursor.Close(context.TODO())

	if !cursor.Next(context.TODO()) {
		return nil, fmt.Errorf("exercise not found for question: %s", question)
	}

	var exercise models.Exercise
	if err := cursor.Decode(&exercise); err != nil {
		return nil, fmt.Errorf("failed to decode exercise: %w", err)
	}

	return &exercise, nil
}
