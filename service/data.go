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
	AddPlayer(firstName string, lastName string, teamName string) (*models.PlayerResolver, error)
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

func (d DatabaseService) AddPlayer(firstName string, lastName string, teamName string) (*models.PlayerResolver, error) {
	teams, err := d.client.GetTeamByName(teamName)

	if err != nil {
		return nil, err
	}

	if len(teams) < 1 {
		return nil, errors.New("Team does not exist. " + teamName)
	}

	team := teams[0]
	newPlayer := models.Player{
		FirstName: firstName,
		LastName:  lastName,
		TeamID:    team.ID,
		Team:      team,
	}

	player, err := d.client.AddPlayer(newPlayer)
	resolved := &models.PlayerResolver{P: &player}

	return resolved, err
}
