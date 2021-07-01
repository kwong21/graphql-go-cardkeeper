// + build unit

package resolver_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/gqltesting"
	"github.com/kwong21/graphql-go-cardkeeper/models"
	"github.com/kwong21/graphql-go-cardkeeper/resolver"
	"github.com/kwong21/graphql-go-cardkeeper/schema"
	"github.com/kwong21/graphql-go-cardkeeper/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func performRequest(r *gin.Engine, method, path, data string) (*httptest.ResponseRecorder, error) {

	if method == "POST" && data == "" {
		return nil, errors.New("requsted a POST but no payload provided.")
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest(method, path, strings.NewReader(data))

	if err != nil {
		return nil, err
	}

	r.ServeHTTP(w, req)

	return w, err
}

func TestQueryResolver_Graphiql(t *testing.T) {
	t.Run("should return graphiql page", func(t *testing.T) {
		r := new(resolver.QueryResolver)

		router := server.NewRouter(r)

		w, err := performRequest(router, "GET", "/graphql", "")

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestRootResolver_Team(t *testing.T) {
	m := new(MockDataService)
	resolver := &resolver.QueryResolver{
		DataService: m,
	}
	schema, _ := schema.String()
	rootSchema := graphql.MustParseSchema(schema, resolver)

	mockTeam := models.Team{
		Name:   "Burnaby Skaters",
		Abbr:   "BBY",
		League: "nhl",
	}

	m.On("GetTeamsByLeague", mock.Anything).Return(mockTeam)
	m.On("AddTeam", mock.Anything, mock.Anything, mock.Anything).Return(mockTeam, nil)

	ctx := context.WithValue(context.Background(), "dataServce", m)

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
	m.AssertExpectations(t)
}

func TestRootResolver_Player(t *testing.T) {
	m := new(MockDataService)
	resolver := &resolver.QueryResolver{
		DataService: m,
	}

	schema, _ := schema.String()
	rootSchema := graphql.MustParseSchema(schema, resolver)

	mockTeam := models.Team{
		Name:   "Burnaby Skaters",
		Abbr:   "BBY",
		League: "nhl",
	}

	mockPlayer := models.Player{
		FirstName: "Viktor",
		LastName:  "Zykov",
		Team:      mockTeam,
	}

	m.On("GetPlayerByName", mock.Anything, mock.Anything).Return(mockPlayer)
	m.On("AddPlayer", mock.Anything, mock.Anything, mock.Anything).Return(mockPlayer, nil)

	ctx := context.WithValue(context.Background(), "dataServce", m)

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
				"player": {
					"id": "0",
					"firstName": "Viktor",
					"lastName": "Zykov",
					"team": {
						"name": "Burnaby Skaters"
					}
				}
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
	m.AssertExpectations(t)
}

// Graphql Testing testing for errors is broken.

// func TestQueryResolver_Team_Failure(t *testing.T) {
// 	m := new(MockDataService)
// 	rs := &resolver.QueryResolver{
// 		DataService: m,
// 	}
// 	schema, _ := schema.String()

// 	rootSchema := graphql.MustParseSchema(schema, rs)

// 	m.On("AddTeam", mock.Anything, mock.Anything, mock.Anything).Return(models.Team{}, errors.New("Error"))
// 	ctx := context.WithValue(context.Background(), "dataServce", m)

// 	gqltesting.RunTest(t, &gqltesting.Test{
// 		Context: ctx,
// 		Schema:  rootSchema,
// 		Query: `
// 			mutation _ {
// 				addTeam(name: "1", abbr: "1", league: "1") {
// 				  id
// 				}
// 			  }
// 		`,
// 		ExpectedResult: "",
// 	},
// 	)
// }
