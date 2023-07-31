package models

type Exercise struct {
	ID       string   `bson:"_id"`
	TopicId  int      `bson:"topicId"`
	Question string   `bson:"question"`
	Variants []string `bson:"variants"`
	Answer   int      `bson:"answer"`
}
