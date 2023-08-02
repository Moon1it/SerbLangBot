package models

type User struct {
	ID              string    `bson:"_id"`             // Unique identifier for the user
	Name            string    `bson:"name"`            // Name of the user
	ChatID          int64     `bson:"chatID"`          // Identifier for the user's chat
	CurrentExercise Exercise  `bson:"currentExercise"` // Currently active exercise for the user
	Stats           UserStats `bson:"stats"`           // Statistics and progress of the user
}

type UserStats struct {
	ProgressByTopics map[int64]TopicProgress `bson:"progressByTopics"` // Progress of the user in various topics
}

type TopicProgress struct {
	AllSolved        int `bson:"allSolved"`        // Total number of problems solved in the topic
	SuccessfulSolved int `bson:"successfulSolved"` // Number of problems successfully solved in the topic
}
