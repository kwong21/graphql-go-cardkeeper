// +build integration

package service_test

import (
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/kwong21/graphql-go-cardkeeper/models"
	"github.com/kwong21/graphql-go-cardkeeper/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntegrationDBClient_PG(t *testing.T) {
	var conf models.Config

	_, err := toml.DecodeFile("../config_local.toml", &conf)

	require.NoError(t, err)

	fixture := service.New(conf)
	require.NotNil(t, fixture)

	t.Run("integration test - Add Team", func(t *testing.T) {
		team, err := fixture.AddTeam("test", "tst", "nhl")

		require.NoError(t, err)
		assert.NotNil(t, team)
		assert.NotEmpty(t, team.ID)
	})

	t.Run("integration test - Get Teams", func(t *testing.T) {
		team, err := fixture.AddTeam("test2", "tst2", "nhl")
		teams := fixture.GetTeamsByLeague("nhl")

		require.NoError(t, err)
		require.NotNil(t, team)
		assert.NotEmpty(t, teams)
	})
}
