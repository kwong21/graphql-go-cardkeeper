package service

import (
	"errors"
	"log"

	"github.com/kwong21/graphql-go-cardkeeper/models"
)

type DataService interface {
	GetAllPlayers() (*[]*models.PlayerResolver, error)
	GetPlayerByID(id string) (*[]*models.PlayerResolver, error)
	GetPlayersOnTeam(team string) (*[]*models.PlayerResolver, error)
	GetAllTeams() (*[]*models.TeamResolver, error)
	GetTeamsByLeague(league string) (*[]*models.TeamResolver, error)
	AddTeam(name string, abbr string, league string) (*models.TeamResolver, error)
	AddPlayer(player models.PlayerInputArgs) (*models.PlayerResolver, error)
}

type DatabaseService struct {
	client DBClient
}

func NewDBService(config models.Config, l Logger) DataService {

	c, err := NewPGClient(config)

	if err != nil {
		log.Fatalf("Not able to create database connection: %s", err)
	}

	return &DatabaseService{
		client: c,
	}
}

func (d DatabaseService) GetAllTeams() (*[]*models.TeamResolver, error) {
	var resolved = make([]*models.TeamResolver, 0)

	teams, err := d.client.GetAllTeams()

	if err != nil {
		return nil, err
	}

	for idx := range teams {
		resolved = append(resolved, &models.TeamResolver{T: &teams[idx]})

	}

	return &resolved, nil
}

func (d DatabaseService) GetTeamsByLeague(league string) (*[]*models.TeamResolver, error) {
	var resolved = make([]*models.TeamResolver, 0)

	teams := d.client.GetTeamsByLeague(league)

	for idx := range teams {
		resolved = append(resolved, &models.TeamResolver{T: &teams[idx]})
	}

	return &resolved, nil
}

func (d DatabaseService) GetAllPlayers() (*[]*models.PlayerResolver, error) {
	var resolved = make([]*models.PlayerResolver, 0)
	players, err := d.client.GetAllPlayers()

	if err != nil {
		return nil, err
	}

	for idx := range players {
		resolved = append(resolved, &models.PlayerResolver{P: &players[idx]})
	}

	return &resolved, nil
}

func (d DatabaseService) GetPlayerByID(id string) (*[]*models.PlayerResolver, error) {
	var resolved = make([]*models.PlayerResolver, 0)
	player, err := d.client.GetPlayerByID(id)

	if err != nil {
		return nil, err
	}

	if player != (models.Player{}) {
		resolved = append(resolved, &models.PlayerResolver{P: &player})
	}

	return &resolved, nil
}

func (d DatabaseService) GetPlayersOnTeam(team string) (*[]*models.PlayerResolver, error) {
	var resolved = make([]*models.PlayerResolver, 0)
	players, err := d.client.GetPlayersOnTeam(team)

	if err != nil {
		return nil, err
	}

	for idx := range players {
		resolved = append(resolved, &models.PlayerResolver{P: &players[idx]})
	}

	return &resolved, nil
}

func (d DatabaseService) AddTeam(name string, abbr string, league string) (*models.TeamResolver, error) {
	newTeam := models.Team{
		Name:   name,
		Abbr:   abbr,
		League: league,
	}

	team, err := d.client.AddTeam(newTeam)
	resolved := &models.TeamResolver{T: &team}

	return resolved, err
}

func (d DatabaseService) AddPlayer(playerInput models.PlayerInputArgs) (*models.PlayerResolver, error) {
	team, err := d.client.GetTeamByID(playerInput.TeamID)

	if err != nil {
		return nil, err
	}

	if (team == models.Team{}) {
		return nil, errors.New("Team with ID " + playerInput.TeamID + " does not exist")
	}

	newPlayer := models.Player{
		FirstName: playerInput.FirstName,
		LastName:  playerInput.LastName,
		TeamID:    team.ID,
		Team:      team,
	}

	player, err := d.client.AddPlayer(newPlayer)
	resolved := &models.PlayerResolver{P: &player}

	return resolved, err
}
