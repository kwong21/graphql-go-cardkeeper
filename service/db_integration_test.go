// +build integration

package service_test

import (
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/kwong21/graphql-go-cardkeeper/mocks"
	"github.com/kwong21/graphql-go-cardkeeper/models"
	"github.com/kwong21/graphql-go-cardkeeper/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_IntegrationDBClient_PG(t *testing.T) {
	var conf models.Config
	mockLoggerClient := new(mocks.MockLoggerClient)

	_, err := toml.DecodeFile("../config_integration.toml", &conf)

	require.NoError(t, err)

	fixture := service.NewDBService(conf, mockLoggerClient)
	require.NotNil(t, fixture)

	t.Run("seed data", func(t *testing.T) {
		_, err := fixture.AddTeam("test1", "test1", "nhl") // id 1

		require.NoError(t, err)

		_, err = fixture.AddTeam("test2", "test2", "mlb") // id 2

		require.NoError(t, err)

		m1 := models.PlayerInputArgs{ // id 1
			FirstName: "Brock",
			LastName:  "Boeser",
			TeamID:    "1",
		}

		p1, err := fixture.AddPlayer(m1)

		require.NoError(t, err)
		require.NotNil(t, p1)

		m2 := models.PlayerInputArgs{ // id 2
			FirstName: "Sana",
			LastName:  "Minatozaki",
			TeamID:    "2",
		}

		p2, err := fixture.AddPlayer(m2)

		require.NoError(t, err)
		require.NotNil(t, p2)
	})

	t.Run("integration test - Get all teams", func(t *testing.T) {
		teams, err := fixture.GetAllTeams()

		require.NoError(t, err)
		require.NotNil(t, teams)
		assert.Equal(t, 2, len(*teams))
	})

	t.Run("integration test - Get Teams", func(t *testing.T) {
		teams, err := fixture.GetTeamsByLeague("nhl")

		require.NoError(t, err)
		require.NotNil(t, teams)
		assert.NotEmpty(t, teams)
		assert.Equal(t, 1, len(*teams))
	})

	t.Run("integration test - Get Players", func(t *testing.T) {
		players, err := fixture.GetAllPlayers()

		require.NoError(t, err)
		require.NotNil(t, players)
		assert.Equal(t, 2, len(*players))
	})

	t.Run("integration test - Get Player on team test2", func(t *testing.T) {
		players, err := fixture.GetPlayersOnTeam("test2")

		require.NoError(t, err)
		require.NotNil(t, players)
		assert.Equal(t, 1, len(*players))
	})

	t.Run("integration test - Get player by ID", func(t *testing.T) {
		player, err := fixture.GetPlayerByID("1")

		test := *player
		foundPlayer := test[0]

		require.NoError(t, err)
		require.NotNil(t, player)
		assert.Equal(t, uint(1), foundPlayer.P.ID)
	})

	t.Run("integration test - error when adding player when team ID does not exist", func(t *testing.T) {
		mockPlayerInput := models.PlayerInputArgs{
			FirstName: "Sana",
			LastName:  "Minatozaki",
			TeamID:    "6",
		}

		player, err := fixture.AddPlayer(mockPlayerInput)

		assert.Error(t, err)
		assert.Empty(t, player)
	})
}
