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

func (s *MockDataService) GetAllTeams() (*[]*models.TeamResolver, error) {
	args := s.Called()

	return args.Get(0).(*[]*models.TeamResolver), args.Error(1)
}

func (s *MockDataService) GetTeamsByLeague(league string) (*[]*models.TeamResolver, error) {
	args := s.Called(league)

	return args.Get(0).(*[]*models.TeamResolver), args.Error(1)
}

func (s *MockDataService) AddTeam(name string, league string, abbr string) (*models.TeamResolver, error) {
	args := s.Called(name, league, abbr)

	return args.Get(0).(*models.TeamResolver), args.Error(1)
}

func (s *MockDataService) GetAllPlayers() (*[]*models.PlayerResolver, error) {
	args := s.Called()

	return args.Get(0).(*[]*models.PlayerResolver), args.Error(1)
}

func (s *MockDataService) GetPlayerByID(id string) (*[]*models.PlayerResolver, error) {
	args := s.Called(id)

	return args.Get(0).(*[]*models.PlayerResolver), args.Error(1)
}

func (s *MockDataService) GetPlayersOnTeam(team string) (*[]*models.PlayerResolver, error) {
	args := s.Called(team)

	return args.Get(0).(*[]*models.PlayerResolver), args.Error(1)
}

func (s *MockDataService) AddPlayer(playerInput models.PlayerInputArgs) (*models.PlayerResolver, error) {
	args := s.Called(playerInput)

	return args.Get(0).(*models.PlayerResolver), args.Error(1)
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
