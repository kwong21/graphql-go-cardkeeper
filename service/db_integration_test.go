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

func TestIntegrationDBClient_PG(t *testing.T) {
	var conf models.Config
	mockLoggerClient := new(mocks.MockLoggerClient)

	_, err := toml.DecodeFile("../config_integration.toml", &conf)

	require.NoError(t, err)

	fixture := service.NewDBService(conf, mockLoggerClient)
	require.NotNil(t, fixture)

	t.Run("integration test - Add Team", func(t *testing.T) {
		team, err := fixture.AddTeam("test", "tst", "nhl")

		require.NoError(t, err)
		assert.NotNil(t, team)
		assert.NotEmpty(t, team.ID)
	})

	t.Run("integration test - Get all teams", func(t *testing.T) {
		_, err := fixture.AddTeam("test1", "test2", "mlb")
		teams, err := fixture.GetAllTeams()

		require.NoError(t, err)
		require.NotNil(t, teams)
		assert.Equal(t, 2, len(teams))
	})

	t.Run("integration test - Get Teams", func(t *testing.T) {
		team, err := fixture.AddTeam("test2", "tst2", "nhl")
		teams := fixture.GetTeamsByLeague("nhl")

		require.NoError(t, err)
		require.NotNil(t, team)
		assert.NotEmpty(t, teams)
	})

	t.Run("integration test - Add Player", func(t *testing.T) {
		player, err := fixture.AddPlayer("Brock", "Boeser", "test")

		require.NoError(t, err)
		require.NotNil(t, player)
		assert.NotEmpty(t, player.ID)
	})

	t.Run("integration test - error when adding player when team ID does not exist", func(t *testing.T) {
		player, err := fixture.AddPlayer("Brock", "Boeser", "Some other team")

		assert.Error(t, err)
		assert.Empty(t, player)
	})
}
