package resolver

import (
	"context"
	"fmt"

	"github.com/graph-gophers/graphql-go/errors"
	"github.com/kwong21/graphql-go-cardkeeper/service"
)

type QueryResolver struct {
	DataService service.DataService
}

func NewRoot(s service.DataService) (*QueryResolver, error) {
	return &QueryResolver{
		DataService: s,
	}, nil
}

func (r QueryResolver) Team(ctx context.Context, args TeamQueryArgs) (*[]*TeamResolver, error) {
	var resolved = make([]*TeamResolver, 0)

	teams := r.DataService.GetTeamsByLeague(args.League)

	for _, t := range teams {
		resolved = append(resolved, &TeamResolver{t: &t})
	}

	return &resolved, nil
}

func (r QueryResolver) AddTeam(ctx context.Context, args TeamQueryArgs) (*TeamResolver, error) {
	t, err := r.DataService.AddTeam(args.Name, args.Abbr, args.League)

	if err != nil {
		qe := &errors.QueryError{
			Message: fmt.Sprintf("error: Unable to add new team: %s", err),
		}
		return nil, qe
	}

	resolved := &TeamResolver{t: &t}

	return resolved, nil
}

func (r QueryResolver) Player(ctx context.Context, args PlayerQueryArgs) (*PlayerResolver, error) {
	p := r.DataService.GetPlayerByName(args.FirstName, args.LastName)

	return &PlayerResolver{p: &p}, nil
}

func (r QueryResolver) AddPlayer(ctx context.Context, args PlayerQueryArgs) (*PlayerResolver, error) {
	p, err := r.DataService.AddPlayer(args.FirstName, args.LastName, args.TeamName)

	if err != nil {
		qe := &errors.QueryError{
			Message: fmt.Sprintf("error: Unable to add new player: %s", err),
		}
		return nil, qe
	}

	resolved := &PlayerResolver{p: &p}

	return resolved, nil
}
