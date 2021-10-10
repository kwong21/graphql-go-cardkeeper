// + build unit

package mocks

import (
	"github.com/kwong21/graphql-go-cardkeeper/models"
	"github.com/stretchr/testify/mock"
)

type MockDataService struct {
	mock.Mock
}

type MockLoggerClient struct {
	mock.Mock
}

func (s *MockDataService) GetAllTeams() ([]models.Team, error) {
	args := s.Called()

	return args.Get(0).([]models.Team), args.Error(1)
}

func (s *MockDataService) GetTeamsByLeague(league string) []models.Team {
	args := s.Called(league)

	return args.Get(0).([]models.Team)
}

func (s *MockDataService) AddTeam(name string, league string, abbr string) (models.Team, error) {
	args := s.Called(name, league, abbr)

	return args.Get(0).(models.Team), args.Error(1)
}

func (s *MockDataService) GetPlayerByName(firstName string, lastName string) ([]models.Player, error) {
	args := s.Called(firstName, lastName)

	return args.Get(0).([]models.Player), args.Error(1)
}

func (s *MockDataService) AddPlayer(firstName string, lastName string, teamName string) (models.Player, error) {
	args := s.Called(firstName, lastName, teamName)

	return args.Get(0).(models.Player), args.Error(1)
}

func (l *MockLoggerClient) Warn(msg string) {
	l.Called(msg)
}

func (l *MockLoggerClient) Debug(msg string) {
	l.Called(msg)
}

func (l *MockLoggerClient) Info(msg string) {
	l.Called(msg)
}

func (l *MockLoggerClient) Error(msg string) {
	l.Called(msg)
}
