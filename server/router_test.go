// + build unit

package server_test

import (
	"testing"

	"github.com/kwong21/graphql-go-cardkeeper/resolver"
	"github.com/kwong21/graphql-go-cardkeeper/server"
	"github.com/stretchr/testify/assert"
)

func TestRouterInit(t *testing.T) {
	t.Run("router should instantiate without error", func(t *testing.T) {
		r := server.NewRouter(new(resolver.QueryResolver))

		assert.NotNil(t, r)
	})
}
