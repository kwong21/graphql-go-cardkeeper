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

	return args.Get(0).(models.Team), nil
}
