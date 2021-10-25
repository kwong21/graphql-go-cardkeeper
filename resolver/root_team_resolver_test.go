// + build unit

package resolver_test

import (
	"context"
	"testing"

	"github.com/graph-gophers/graphql-go/gqltesting"
	"github.com/kwong21/graphql-go-cardkeeper/models"
	"github.com/stretchr/testify/mock"
)

const (
	teamData key = iota
	allTeamData
)

func TestRootResolver_Team_NoResults(t *testing.T) {
	rootSchema, mockDataService, _ := getTestFixtures()

	mockDataService.On("GetTeamsByLeague", mock.Anything).Return(&[]*models.TeamResolver{}, nil)
	ctx := context.WithValue(context.Background(), teamData, mockDataService)

	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Context: ctx,
			Schema:  rootSchema,
			Query: `
			{
				team(league: "1") {
					id
					name
					abbr
					league
				}
			}
			`,
			ExpectedResult: `
			{
				"team": []
			}
			`,
		},
	})
	mockDataService.AssertExpectations(t)
}

func TestRootResolver_Teams(t *testing.T) {
	rootSchema, mockDataService, _ := getTestFixtures()

	mockTeams := []*models.TeamResolver{
		&models.TeamResolver{T: &mockHockeyTeam},
		&models.TeamResolver{T: &mockBaseBallTeam},
	}

	mockDataService.On("GetAllTeams").Return(&mockTeams, nil)
	ctx := context.WithValue(context.Background(), allTeamData, mockDataService)

	gqltesting.RunTest(t, &gqltesting.Test{
		Context: ctx,
		Schema:  rootSchema,
		Query: `
		{
			teams {
			name 
			}
		}
		`,
		ExpectedResult: `
		{
			"teams": [{
				"name": "Burnaby Skaters"
			},
			{
				"name": "Vancouver Canadians"	
			}]
		}
		`,
	})
}

func TestRootResolver_Team(t *testing.T) {
	rootSchema, mockDataService, _ := getTestFixtures()

	nhlTeams := []models.Team{
		mockHockeyTeam,
	}

	mockDataService.On("GetTeamsByLeague", mock.Anything).Return(nhlTeams)
	mockDataService.On("AddTeam", mock.Anything, mock.Anything, mock.Anything).Return(mockHockeyTeam, nil)

	ctx := context.WithValue(context.Background(), teamData, mockDataService)

	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Context: ctx,
			Schema:  rootSchema,
			Query: `
			{
				team(league: "1") {
					id
					name
					abbr
					league
				}
			}
			`,
			ExpectedResult: `
			{
				"team": [{
					"id": "0",
					"name": "Burnaby Skaters",
					"abbr": "BBY",
					"league": "nhl"
				}]
			}
			`,
		},
		{
			Context: ctx,
			Schema:  rootSchema,
			Query: `
			mutation _ {
				addTeam(name: "1", abbr: "1", league: "1") {
				  id
				}
			  }
			`,
			ExpectedResult: `
			{
				"addTeam": {
					"id": "0"
				}
			}
			`,
		},
	})
	mockDataService.AssertExpectations(t)
}
