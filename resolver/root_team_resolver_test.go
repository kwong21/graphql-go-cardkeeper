// + build unit

package resolver_test

import (
	"context"
	"testing"

	"github.com/graph-gophers/graphql-go/gqltesting"
	"github.com/kwong21/graphql-go-cardkeeper/models"
	"github.com/stretchr/testify/mock"
)

func TestRootResolver_Team_NoResults(t *testing.T) {
	rootSchema, mockDataService, _ := getTestFixtures()

	mockDataService.On("GetTeamsByLeague", mock.Anything).Return([]models.Team{})
	ctx := context.WithValue(context.Background(), "dataservice", mockDataService)

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

func TestRootResolver_Team(t *testing.T) {
	rootSchema, mockDataService, _ := getTestFixtures()

	mockDataService.On("GetTeamsByLeague", mock.Anything).Return(mockTeams)
	mockDataService.On("AddTeam", mock.Anything, mock.Anything, mock.Anything).Return(mockTeam, nil)

	ctx := context.WithValue(context.Background(), "dataServce", mockDataService)

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
