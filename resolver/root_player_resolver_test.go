// + build unit

package resolver_test

import (
	"context"
	"testing"

	"github.com/graph-gophers/graphql-go/gqltesting"
	"github.com/kwong21/graphql-go-cardkeeper/models"
	"github.com/stretchr/testify/mock"
)

type key int

const (
	dataserviceKey key = iota
)

func TestRootResolver_Player_NoResults(t *testing.T) {
	rootSchema, mockDataService, _ := getTestFixtures()

	mockDataService.On("GetPlayerByName", mock.Anything, mock.Anything).Return([]models.Player{}, nil)

	ctx := context.WithValue(context.Background(), dataserviceKey, mockDataService)

	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Context: ctx,
			Schema:  rootSchema,
			Query: `
			{
				player(firstName: "Viktor", lastName: "Zykov") {
					id
					firstName
					lastName
					team {
						name
					}
				}
			}
			`,
			ExpectedResult: `
			{
				"player": []
			}
			`,
		},
	})
	mockDataService.AssertExpectations(t)
}

func TestRootResolver_Player(t *testing.T) {
	rootSchema, mockDataService, _ := getTestFixtures()

	mockDataService.On("GetPlayerByName", mock.Anything, mock.Anything).Return(mockPlayers, nil)
	mockDataService.On("AddPlayer", mock.Anything, mock.Anything, mock.Anything).Return(mockPlayer, nil)

	ctx := context.WithValue(context.Background(), dataserviceKey, mockDataService)

	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Context: ctx,
			Schema:  rootSchema,
			Query: `
			{
				player(firstName: "Viktor", lastName: "Zykov") {
					id
					firstName
					lastName
					team {
						name
					}
				}
			}
			`,
			ExpectedResult: `
			{
				"player": [{
					"id": "0",
					"firstName": "Viktor",
					"lastName": "Zykov",
					"team": {
						"name": "Burnaby Skaters"
					}
				}]
			}
			`,
		},
		{
			Context: ctx,
			Schema:  rootSchema,
			Query: `
			mutation _ {
				addPlayer(firstName: "Ollie", lastName: "Inu", teamName: "Dogs") {
				  id
				}
			  }
			`,
			ExpectedResult: `
			{
				"addPlayer": {
					"id": "0"
				}
			}
			`,
		},
	})
	mockDataService.AssertExpectations(t)
}
