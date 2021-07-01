package service

import (
	"log"

	"github.com/kwong21/graphql-go-cardkeeper/models"
)

type DataService interface {
	GetPlayerByName(firstName string, lastName string) models.Player
	GetTeamsByLeague(league string) []models.Team
	AddTeam(name string, abbr string, league string) (models.Team, error)
	AddPlayer(firstName string, lastName string, teamName string) (models.Player, error)
}

type DatabaseService struct {
	client DBClient
}

func New(config models.Config) DataService {

	c, err := NewPGClient(config)

	if err != nil {
		log.Fatalf("Not able to create database connection: %s", err)
	}

	return &DatabaseService{
		client: c,
	}
}

func (d DatabaseService) GetTeamsByLeague(league string) []models.Team {
	teams := d.client.GetTeamsByLeague(league)

	return teams
}

func (d DatabaseService) GetPlayerByName(firstName string, lastName string) models.Player {
	player := d.client.GetPlayerByName(firstName, lastName)

	return player
}

func (d DatabaseService) AddTeam(name string, abbr string, league string) (models.Team, error) {
	newTeam := models.Team{
		Name:   name,
		Abbr:   abbr,
		League: league,
	}

	team, err := d.client.AddTeam(newTeam)

	return team, err
}

func (d DatabaseService) AddPlayer(firstName string, lastName string, teamName string) (models.Player, error) {
	t, err := d.client.GetTeamByName(teamName)

	if err != nil {
		return models.Player{}, err
	}

	newPlayer := models.Player{
		FirstName: firstName,
		LastName:  lastName,
		TeamID:    t.ID,
		Team:      t,
	}

	player, err := d.client.AddPlayer(newPlayer)

	return player, err
}
