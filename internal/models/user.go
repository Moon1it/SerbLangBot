package models

type User struct {
	ID              string    `bson:"_id"`
	Name            string    `bson:"name"`
	ChatID          int64     `bson:"chatID"`
	CurrentExercise Exercise  `bson:"currentExercise"`
	Stats           UserStats `bson:"stats"`
}

type UserStats struct {
	TopicProgress map[int64]TopicProgress `bson:"topicProgress"`
}

type TopicProgress struct {
	AllSolved        int `bson:"allSolved"`
	SuccessfulSolved int `bson:"successfulSolved"`
}
