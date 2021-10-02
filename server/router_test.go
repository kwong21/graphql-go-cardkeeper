// + build unit

package server_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
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

func TestRouterWithGraphQL(t *testing.T) {
	t.Run("should return graphiql page", func(t *testing.T) {
		r := new(resolver.QueryResolver)

		router := server.NewRouter(r)

		w, err := performRequest(router, "GET", "/graphql", "")

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func performRequest(r *gin.Engine, method, path, data string) (*httptest.ResponseRecorder, error) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest(method, path, strings.NewReader(data))

	if err != nil {
		return nil, err
	}

	r.ServeHTTP(w, req)

	return w, err
}
