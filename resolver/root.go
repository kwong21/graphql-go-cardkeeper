package resolver

import (
	"context"
	"fmt"

	"github.com/graph-gophers/graphql-go/errors"
	"github.com/kwong21/graphql-go-cardkeeper/service"
)

type QueryResolver struct {
	DataService   service.DataService
	LoggerService service.Logger
}

func NewRoot(s service.DataService, l service.Logger) (*QueryResolver, error) {
	return &QueryResolver{
		DataService:   s,
		LoggerService: l,
	}, nil
}

// Team resolves the query to GET teams belonging in `league`
func (r QueryResolver) Team(ctx context.Context, args TeamQueryArgs) (*[]*TeamResolver, error) {
	var resolved = make([]*TeamResolver, 0)

	r.LoggerService.Info("Retrieving teams in league " + args.League)
	teams := r.DataService.GetTeamsByLeague(args.League)

	for _, t := range teams {
		resolved = append(resolved, &TeamResolver{t: &t})
	}

	r.LoggerService.Info(fmt.Sprintf("Found %d teams", len(resolved)))

	return &resolved, nil
}

// AddTeam resolves the query to POST a team to the database
func (r QueryResolver) AddTeam(ctx context.Context, args TeamQueryArgs) (*TeamResolver, error) {
	t, err := r.DataService.AddTeam(args.Name, args.Abbr, args.League)

	if err != nil {
		r.LoggerService.Error("Could not add team to database")
		r.LoggerService.Error(err.Error())

		qe := &errors.QueryError{
			Message: fmt.Sprintf("error: Unable to add new team: %s", err),
		}
		return nil, qe
	}

	resolved := &TeamResolver{t: &t}

	return resolved, nil
}

// Player resolves the query to GET players with `first_name` and `last_name`
func (r QueryResolver) Player(ctx context.Context, args PlayerQueryArgs) (*[]*PlayerResolver, error) {
	var resolved = make([]*PlayerResolver, 0)

	r.LoggerService.Info(fmt.Sprintf("Retrieving player with first_name %s and last_name %s", args.FirstName, args.LastName))
	players, err := r.DataService.GetPlayerByName(args.FirstName, args.LastName)

	for _, p := range players {
		resolved = append(resolved, &PlayerResolver{p: &p})
	}

	if err != nil {
		r.LoggerService.Error(fmt.Sprintf("Could not retrieve player with first_name %s and last_name %s", args.FirstName, args.LastName))
		r.LoggerService.Error(err.Error())

		qe := &errors.QueryError{
			Message: err.Error(),
		}

		return nil, qe
	}

	return &resolved, nil
}

// AddPlayer reoslves the query to POST a player to the database
func (r QueryResolver) AddPlayer(ctx context.Context, args PlayerQueryArgs) (*PlayerResolver, error) {
	p, err := r.DataService.AddPlayer(args.FirstName, args.LastName, args.TeamName)

	if err != nil {
		r.LoggerService.Error("Could not add player to database")
		r.LoggerService.Error(err.Error())

		qe := &errors.QueryError{
			Message: fmt.Sprintf("error: Unable to add new player: %s", err),
		}
		return nil, qe
	}

	resolved := &PlayerResolver{p: &p}

	return resolved, nil
}
