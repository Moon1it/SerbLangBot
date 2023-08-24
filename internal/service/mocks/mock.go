// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	models "github.com/Moon1it/SerbLangBot/internal/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	gomock "github.com/golang/mock/gomock"
)

// MockUser is a mock of User interface.
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
}

// MockUserMockRecorder is the mock recorder for MockUser.
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance.
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUser) CreateUser(message *tgbotapi.Message) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", message)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserMockRecorder) CreateUser(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUser)(nil).CreateUser), message)
}

// GetUser mocks base method.
func (m *MockUser) GetUser(message *tgbotapi.Message) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", message)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserMockRecorder) GetUser(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUser)(nil).GetUser), message)
}

// GetUserProgress mocks base method.
func (m *MockUser) GetUserProgress(chatID int64) (*tgbotapi.MessageConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserProgress", chatID)
	ret0, _ := ret[0].(*tgbotapi.MessageConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserProgress indicates an expected call of GetUserProgress.
func (mr *MockUserMockRecorder) GetUserProgress(chatID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserProgress", reflect.TypeOf((*MockUser)(nil).GetUserProgress), chatID)
}

// UpdateTopicProgress mocks base method.
func (m *MockUser) UpdateTopicProgress(chatID int64, answer int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTopicProgress", chatID, answer)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTopicProgress indicates an expected call of UpdateTopicProgress.
func (mr *MockUserMockRecorder) UpdateTopicProgress(chatID, answer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTopicProgress", reflect.TypeOf((*MockUser)(nil).UpdateTopicProgress), chatID, answer)
}

// MockNavigation is a mock of Navigation interface.
type MockNavigation struct {
	ctrl     *gomock.Controller
	recorder *MockNavigationMockRecorder
}

// MockNavigationMockRecorder is the mock recorder for MockNavigation.
type MockNavigationMockRecorder struct {
	mock *MockNavigation
}

// NewMockNavigation creates a new mock instance.
func NewMockNavigation(ctrl *gomock.Controller) *MockNavigation {
	mock := &MockNavigation{ctrl: ctrl}
	mock.recorder = &MockNavigationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNavigation) EXPECT() *MockNavigationMockRecorder {
	return m.recorder
}

// GetMainMenuMessage mocks base method.
func (m *MockNavigation) GetMainMenuMessage(chatID int64) (*tgbotapi.MessageConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMainMenuMessage", chatID)
	ret0, _ := ret[0].(*tgbotapi.MessageConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMainMenuMessage indicates an expected call of GetMainMenuMessage.
func (mr *MockNavigationMockRecorder) GetMainMenuMessage(chatID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMainMenuMessage", reflect.TypeOf((*MockNavigation)(nil).GetMainMenuMessage), chatID)
}

// GetStartMessage mocks base method.
func (m *MockNavigation) GetStartMessage(message *tgbotapi.Message) (*tgbotapi.MessageConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStartMessage", message)
	ret0, _ := ret[0].(*tgbotapi.MessageConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStartMessage indicates an expected call of GetStartMessage.
func (mr *MockNavigationMockRecorder) GetStartMessage(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStartMessage", reflect.TypeOf((*MockNavigation)(nil).GetStartMessage), message)
}

// GetTopicMenuMessage mocks base method.
func (m *MockNavigation) GetTopicMenuMessage(chatID int64) (*tgbotapi.MessageConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTopicMenuMessage", chatID)
	ret0, _ := ret[0].(*tgbotapi.MessageConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTopicMenuMessage indicates an expected call of GetTopicMenuMessage.
func (mr *MockNavigationMockRecorder) GetTopicMenuMessage(chatID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTopicMenuMessage", reflect.TypeOf((*MockNavigation)(nil).GetTopicMenuMessage), chatID)
}

// MockExercise is a mock of Exercise interface.
type MockExercise struct {
	ctrl     *gomock.Controller
	recorder *MockExerciseMockRecorder
}

// MockExerciseMockRecorder is the mock recorder for MockExercise.
type MockExerciseMockRecorder struct {
	mock *MockExercise
}

// NewMockExercise creates a new mock instance.
func NewMockExercise(ctrl *gomock.Controller) *MockExercise {
	mock := &MockExercise{ctrl: ctrl}
	mock.recorder = &MockExerciseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExercise) EXPECT() *MockExerciseMockRecorder {
	return m.recorder
}

// GetExercisePoll mocks base method.
func (m *MockExercise) GetExercisePoll(chatID int64, exerciseType string) (*tgbotapi.SendPollConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExercisePoll", chatID, exerciseType)
	ret0, _ := ret[0].(*tgbotapi.SendPollConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExercisePoll indicates an expected call of GetExercisePoll.
func (mr *MockExerciseMockRecorder) GetExercisePoll(chatID, exerciseType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExercisePoll", reflect.TypeOf((*MockExercise)(nil).GetExercisePoll), chatID, exerciseType)
}