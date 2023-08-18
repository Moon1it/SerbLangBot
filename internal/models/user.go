package models

import "time"

type User struct {
	ID             string    `bson:"_id"`            // Unique identifier for the user
	Name           string    `bson:"name"`           // Name of the user
	ChatID         int64     `bson:"chatID"`         // Identifier for the user's chat
	ActiveExercise Exercise  `bson:"activeExercise"` // Currently active exercise for the user
	Stats          UserStats `bson:"stats"`          // Statistics and progress of the user
	RegisteredAt   time.Time `bson:"registeredAt"`   // Date of user registration
}

type UserStats struct {
	ProgressByTopics map[int64]TopicStats `bson:"progressByTopics"` // Progress of the user in various topics
}

type TopicStats struct {
	AllSolved        int `bson:"allSolved"`        // Total number of exercises solved in the topic
	SuccessfulSolved int `bson:"successfulSolved"` // Number of exercises successfully solved in the topic
}
