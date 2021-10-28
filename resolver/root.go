package resolver

import (
	"context"
	"fmt"

	"github.com/graph-gophers/graphql-go/errors"
	"github.com/kwong21/graphql-go-cardkeeper/models"
	"github.com/kwong21/graphql-go-cardkeeper/service"
)

type QueryResolver struct {
	DataService   service.DataService
	LoggerService service.Logger
}

type PlayerQueryArgs struct {
	ID        string
	FirstName string
	LastName  string
	TeamName  *string
}

type TeamQueryArgs struct {
	Name   string
	League string
	Abbr   string
}

func NewRoot(s service.DataService, l service.Logger) (*QueryResolver, error) {
	return &QueryResolver{
		DataService:   s,
		LoggerService: l,
	}, nil
}

// Teams resolves the query to GET all teams
func (r QueryResolver) Teams(ctx context.Context) (*[]*models.TeamResolver, error) {
	r.LoggerService.Info("Retrieving all teams")

	teams, err := r.DataService.GetAllTeams()

	if err != nil {
		r.LoggerService.Error("Could not retrieve all teams in the database")
		qe := r.createErrorResponse(err)

		return nil, qe
	}

	return teams, nil
}

// Team resolves the query to GET teams belonging in `league`
func (r QueryResolver) Team(ctx context.Context, args TeamQueryArgs) (*[]*models.TeamResolver, error) {
	r.LoggerService.Info("Retrieving teams in league " + args.League)
	teams, err := r.DataService.GetTeamsByLeague(args.League)

	if err != nil {
		r.LoggerService.Error("Could not get team from league" + args.League)

		qe := r.createErrorResponse(err)

		return nil, qe
	}

	return teams, nil
}

// AddTeam resolves the query to POST a team to the database
func (r QueryResolver) AddTeam(ctx context.Context, args TeamQueryArgs) (*models.TeamResolver, error) {
	t, err := r.DataService.AddTeam(args.Name, args.Abbr, args.League)

	if err != nil {
		r.LoggerService.Error("Could not add team to database")

		qe := r.createErrorResponse(err)

		return nil, qe
	}

	return t, nil
}

// Players resolves the query to GET all players or all players on a team if a param is provided
func (r QueryResolver) Players(ctx context.Context, args PlayerQueryArgs) (*[]*models.PlayerResolver, error) {
	var players *[]*models.PlayerResolver
	var err error

	if args.TeamName == nil {
		r.LoggerService.Info("GET all players")
		players, err = r.DataService.GetAllPlayers()
	} else {
		r.LoggerService.Info(fmt.Sprintf("GET players on team %s", *args.TeamName))
		players, err = r.DataService.GetPlayersOnTeam(*args.TeamName)
	}

	if err != nil {
		r.LoggerService.Error("Could not get Players from database")
		qe := r.createErrorResponse(err)
		return nil, qe
	}

	return players, nil
}

// Player resolves the query to GET a player with queried ID
func (r QueryResolver) Player(ctx context.Context, args PlayerQueryArgs) (*[]*models.PlayerResolver, error) {
	r.LoggerService.Info(fmt.Sprintf("Retrieving player with ID %s", args.ID))
	player, err := r.DataService.GetPlayerByID(args.ID)

	if err != nil {
		r.LoggerService.Error(fmt.Sprintf("Could not retrieve player with ID %s", args.ID))
		qe := r.createErrorResponse(err)

		return nil, qe
	}

	return player, nil
}

// AddPlayer reoslves the query to POST a player to the database
func (r QueryResolver) AddPlayer(ctx context.Context, args struct{ Player models.PlayerInputArgs }) (*models.PlayerResolver, error) {
	p, err := r.DataService.AddPlayer(args.Player)

	if err != nil {
		r.LoggerService.Error("Could not add player to database")
		qe := r.createErrorResponse(err)

		return nil, qe
	}

	return p, nil
}

func (r QueryResolver) createErrorResponse(err error) error {
	r.LoggerService.Error(err.Error())

	qe := &errors.QueryError{
		Message: err.Error(),
	}

	return qe
}
