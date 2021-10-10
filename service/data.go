package service

import (
	"errors"
	"log"

	"github.com/kwong21/graphql-go-cardkeeper/models"
)

type DataService interface {
	GetPlayerByName(firstName string, lastName string) ([]models.Player, error)
	GetAllTeams() ([]models.Team, error)
	GetTeamsByLeague(league string) []models.Team
	AddTeam(name string, abbr string, league string) (models.Team, error)
	AddPlayer(firstName string, lastName string, teamName string) (models.Player, error)
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

func (d DatabaseService) GetAllTeams() ([]models.Team, error) {
	return d.client.GetAllTeams()
}

func (d DatabaseService) GetTeamsByLeague(league string) []models.Team {
	return d.client.GetTeamsByLeague(league)
}

func (d DatabaseService) GetPlayerByName(firstName string, lastName string) ([]models.Player, error) {
	return d.client.GetPlayerByName(firstName, lastName)
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
	teams, err := d.client.GetTeamByName(teamName)

	if err != nil {
		return models.Player{}, err
	}

	if len(teams) < 1 {
		return models.Player{}, errors.New("Team does not exist. " + teamName)
	}

	team := teams[0]
	newPlayer := models.Player{
		FirstName: firstName,
		LastName:  lastName,
		TeamID:    team.ID,
		Team:      team,
	}

	player, err := d.client.AddPlayer(newPlayer)

	return player, err
}
