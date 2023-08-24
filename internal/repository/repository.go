package repository

import (
	"github.com/Moon1it/SerbLangBot/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type User interface {
	Create(newUser *models.User) error
	GetByChatID(chatID int64) (*models.User, error)
	GetStatsByChatID(chatID int64) (*models.UserStats, error)
	UpdateActiveExercise(chatID int64, activeExercise *models.Exercise) error
	UpdateTopicProgress(chatID int64, topicID int64, topicProgress *models.TopicStats) error
}

type Message interface {
	GetServiceMessage(name string) (*models.ServiceMessage, error)
}

type Topic interface {
	GetTopicsCount() (int64, error)
	GetAllTopics() ([]models.Topic, error)
}

type Exercise interface {
	GetByTopic(topicName string) (*models.Exercise, error)
	GetAnswerByQuestion(question string) (*models.Exercise, error)
}

type Repository struct {
	User
	Message
	Topic
	Exercise
}

func InitRepository(db *mongo.Client) *Repository {
	return &Repository{
		User:     InitUserRepo(db),
		Message:  InitMessageRepo(db),
		Topic:    InitTopicRepo(db),
		Exercise: InitExerciseRepo(db),
	}
}
