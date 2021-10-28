// + build unit

package resolver_test

import (
	"context"
	"testing"

	"github.com/graph-gophers/graphql-go/gqltesting"
	"github.com/kwong21/graphql-go-cardkeeper/models"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type key int

const (
	playerData key = iota
)

func TestRootResolver_Player_NoResults(t *testing.T) {
	rootSchema, mockDataService, _ := getTestFixtures()

	mockDataService.On("GetPlayerByID", mock.Anything).Return(&[]*models.PlayerResolver{}, nil)

	ctx := context.WithValue(context.Background(), playerData, mockDataService)

	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Context: ctx,
			Schema:  rootSchema,
			Query: `
			{
				player(id: 1) {
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

func TestRootResolver_Players(t *testing.T) {
	rootSchema, mockDataService, _ := getTestFixtures()

	mockTeamTwo := models.Team{
		Name:   "JYP Twice",
		Abbr:   "JYP",
		League: "kbl",
	}

	mockPlayerTwo := &models.Player{
		Model: gorm.Model{
			ID: 99,
		},
		FirstName: "Sana",
		LastName:  "Minatozaki",
		Team:      mockTeamTwo,
	}

	testPlayers := &[]*models.PlayerResolver{
		&models.PlayerResolver{P: &mockPlayer},
		&models.PlayerResolver{P: mockPlayerTwo},
	}

	mockDataService.On("GetAllPlayers").Return(testPlayers, nil)
	mockDataService.On("GetPlayersOnTeam", "JYP Twice").Return(&[]*models.PlayerResolver{
		&models.PlayerResolver{P: mockPlayerTwo}}, nil)
	ctx := context.WithValue(context.Background(), playerData, mockDataService)

	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Context: ctx,
			Schema:  rootSchema,
			Query: `
			{
				players {
				firstName
				team {
					name
				}
			}
			}
			`,
			ExpectedResult: `
			{
				"players":
				[
					{
						"firstName": "Viktor",
						"team": {
						"name": "Burnaby Skaters"
						}
					},
					{
						"firstName": "Sana",
						"team": {
							"name": "JYP Twice"
						}
					}
				]
			}
		`,
		},
	})
}

func TestRootResolver_Player(t *testing.T) {
	rootSchema, mockDataService, _ := getTestFixtures()

	mockResolvedPlayer := &[]*models.PlayerResolver{
		&models.PlayerResolver{P: &mockPlayer},
	}

	mockDataService.On("GetPlayerByID", mock.Anything).Return(mockResolvedPlayer, nil)
	mockDataService.On("AddPlayer", mock.Anything).Return(&models.PlayerResolver{P: &mockPlayer}, nil)

	ctx := context.WithValue(context.Background(), playerData, mockDataService)

	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Context: ctx,
			Schema:  rootSchema,
			Query: `
			{
				player(id: 99) {
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
					"id": "99",
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
			mutation addPlayer {
				addPlayer(
					player: {
						firstName: "Ollie", lastName: "Inu", teamID: "1"
					}) {
				  id
				}
			  }
			`,
			ExpectedResult: `
			{
				"addPlayer": {
					"id": "99"
				}
			}
			`,
		},
	})
	mockDataService.AssertExpectations(t)
}
