// + build unit

package resolver_test

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/kwong21/graphql-go-cardkeeper/mocks"
	"github.com/kwong21/graphql-go-cardkeeper/models"
	"github.com/kwong21/graphql-go-cardkeeper/resolver"
	"github.com/kwong21/graphql-go-cardkeeper/schema"
	"github.com/stretchr/testify/mock"
)

var mockHockeyTeam = models.Team{
	Name:   "Burnaby Skaters",
	Abbr:   "BBY",
	League: "nhl",
}

var mockBaseBallTeam = models.Team{
	Name:   "Vancouver Canadians",
	Abbr:   "CND",
	League: "mlb",
}

var mockPlayer = models.Player{
	FirstName: "Viktor",
	LastName:  "Zykov",
	Team:      mockHockeyTeam,
}

var mockPlayers = []models.Player{mockPlayer}

func getTestFixtures() (*graphql.Schema, *mocks.MockDataService, *mocks.MockLoggerClient) {
	m := new(mocks.MockDataService)
	l := new(mocks.MockLoggerClient)

	resolver := &resolver.QueryResolver{
		DataService:   m,
		LoggerService: l,
	}

	schema, _ := schema.String()
	rootSchema := graphql.MustParseSchema(schema, resolver)

	l.On("Info", mock.AnythingOfType("string"))
	l.On("Error", mock.AnythingOfType("string"))

	return rootSchema, m, l
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
