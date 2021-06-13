// + build unit

package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kwong21/graphql-go-cardkeeper/controller"
	"github.com/stretchr/testify/assert"
)

func TestGrpahiQLHandler_shouldGiveStatusOK(t *testing.T) {
	t.Run("should return HTTP 200 when requesting homepage", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		fixture := new(controller.GraphqlController)
		fixture.GraphQLHandlerFunc(c)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}
