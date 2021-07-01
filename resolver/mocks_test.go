// + build unit

package resolver_test

import (
	"github.com/kwong21/graphql-go-cardkeeper/models"
	"github.com/stretchr/testify/mock"
)

type MockDataService struct {
	mock.Mock
}

func (s *MockDataService) GetTeamsByLeague(league string) []models.Team {
	var teams []models.Team
	args := s.Called(league)
	t := args.Get(0).(models.Team)

	teams = append(teams, t)

	return teams
}

func (s *MockDataService) AddTeam(name string, league string, abbr string) (models.Team, error) {
	args := s.Called(name, league, abbr)

	return args.Get(0).(models.Team), args.Error(1)
}

func (s *MockDataService) GetPlayerByName(firstName string, lastName string) models.Player {
	args := s.Called(firstName, lastName)

	return args.Get(0).(models.Player)
}

func (s *MockDataService) AddPlayer(firstName string, lastName string, teamName string) (models.Player, error) {
	args := s.Called(firstName, lastName, teamName)

	return args.Get(0).(models.Player), args.Error(1)
}
