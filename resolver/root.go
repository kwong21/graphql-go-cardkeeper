package resolver

import (
	"context"
	"fmt"

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
		return nil, customError{
			Code:    "400",
			Message: fmt.Sprintf("Unable to add new team: %s", err),
		}
	}

	resolved := &TeamResolver{t: &t}

	return resolved, nil
}
