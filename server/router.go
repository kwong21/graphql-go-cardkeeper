package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kwong21/graphql-go-cardkeeper/controller"
	"github.com/kwong21/graphql-go-cardkeeper/resolver"
)

func NewRouter(r *resolver.QueryResolver) *gin.Engine {
	router := gin.Default()

	graphql := controller.GraphqlController{
		Resolver: r,
	}

	router.GET("/graphql", graphql.GraphQLHandlerFunc)
	router.POST("/graphql", gin.WrapH(graphql.QueryFunc()))

	return router
}
