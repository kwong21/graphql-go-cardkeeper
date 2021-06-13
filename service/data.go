package service

import (
	"log"

	"github.com/kwong21/graphql-go-cardkeeper/models"
)

type DataService interface {
	GetTeamsByLeague(league string) []models.Team
	AddTeam(name string, abbr string, league string) (models.Team, error)
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

func (d DatabaseService) AddTeam(name string, abbr string, league string) (models.Team, error) {
	newTeam := models.Team{
		Name:   name,
		Abbr:   abbr,
		League: league,
	}

	team, err := d.client.AddTeam(newTeam)

	return team, err
}
