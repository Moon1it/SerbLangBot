package handler

import (
	"testing"
	"time"

	"github.com/Moon1it/SerbLangBot/internal/models"
	"github.com/Moon1it/SerbLangBot/internal/service"
	mock_service "github.com/Moon1it/SerbLangBot/internal/service/mocks"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const startMessage = `🌟Hello and welcome to the Serbian Language Learning Bot!🌟

I'm your dedicated language learning companion designed to help you master Serbian. Learning a new language comes with a host of exciting benefits:

📖 Start Learning: Begin your language lessons and practice Serbian at your own pace. Whether you're a complete beginner or looking to enhance your existing knowledge, we've got you covered!

🔄 Daily Practice: Engage in daily exercises to reinforce your language skills. Consistency is the key to language mastery, and our daily practice sessions will keep you on track.

🎯 Vocabulary: Access our comprehensive Serbian vocabulary database. Expand your word bank and express yourself fluently in Serbian.

🎮 Language Games: Have fun and learn through interactive games. Learning doesn't have to be boring - our language games make the process enjoyable and engaging!

📈 Progress & Stats: Track your progress and see how far you've come. Celebrate your achievements and stay motivated as you see yourself making steady progress.

To get started, simply type /menu and explore the various options available in our main menu. Let's embark on this language adventure together! 🚀😊`

func TestHandler_HandleMessage_Start(t *testing.T) {
	type mockBehavior func(user *mock_service.MockUser, navigation *mock_service.MockNavigation, message *tgbotapi.Message)

	testTable := []struct {
		name         string
		message      *tgbotapi.Message
		mockBehavior mockBehavior
		expected     *models.MessageResponse
	}{
		{
			name: "KnownUser",
			message: &tgbotapi.Message{
				Text: "/start",
				Chat: &tgbotapi.Chat{ID: 123456789},
			},
			mockBehavior: func(user *mock_service.MockUser, navigation *mock_service.MockNavigation, message *tgbotapi.Message) {
				user.EXPECT().GetUser(message).Return(&models.User{
					ID:             "user123",
					Name:           "John Doe",
					ChatID:         123456789,
					ActiveExercise: models.Exercise{
						// Заполните данные упражнения, если нужно
					},
					Stats: models.UserStats{
						ProgressByTopics: []models.TopicStats{
							{
								AllSolved:        10,
								SuccessfulSolved: 7,
							},
							{
								AllSolved:        5,
								SuccessfulSolved: 3,
							},
							// ... добавьте статистику по другим темам
						},
					},
					RegisteredAt: time.Now(),
				}, nil)
				navigation.EXPECT().GetStartMessage(message).Return(&tgbotapi.MessageConfig{
					Text: startMessage,
				}, nil)
			},
			expected: &models.MessageResponse{
				Message: &tgbotapi.MessageConfig{
					Text: startMessage,
				},
				Poll: nil,
			},
		},
		{
			name: "UnknownUser",
			message: &tgbotapi.Message{
				Text: "/start",
				Chat: &tgbotapi.Chat{ID: 987654321},
			},
			mockBehavior: func(user *mock_service.MockUser, navigation *mock_service.MockNavigation, message *tgbotapi.Message) {
				user.EXPECT().GetUser(message).Return(nil, service.ErrUserNotFound)
				user.EXPECT().CreateUser(message).Return(&models.User{
					ID:             "d33a11b5-442f-4d9e-b4f6-b173eece8bda",
					Name:           "John Doe",
					ChatID:         987654321,
					ActiveExercise: models.Exercise{},
					Stats: models.UserStats{
						ProgressByTopics: []models.TopicStats{
							{
								AllSolved:        10,
								SuccessfulSolved: 7,
							},
							{
								AllSolved:        5,
								SuccessfulSolved: 3,
							},
							{
								AllSolved:        10,
								SuccessfulSolved: 7,
							},
							{
								AllSolved:        5,
								SuccessfulSolved: 3,
							},
							{
								AllSolved:        11,
								SuccessfulSolved: 7,
							},
							{
								AllSolved:        5,
								SuccessfulSolved: 4,
							},
						},
					},
					RegisteredAt: time.Now(),
				}, nil)
				navigation.EXPECT().GetStartMessage(message).Return(&tgbotapi.MessageConfig{
					Text: startMessage,
				}, nil)
			},
			expected: &models.MessageResponse{
				Message: &tgbotapi.MessageConfig{
					Text: startMessage,
				},
				Poll: nil,
			},
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			user := mock_service.NewMockUser(ctrl)
			navigation := mock_service.NewMockNavigation(ctrl)

			test.mockBehavior(user, navigation, test.message)

			services := &service.Service{
				User:       user,
				Navigation: navigation,
			}
			handler := InitHandler(services)

			result, err := handler.HandleMessage(test.message)

			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		})
	}
}
