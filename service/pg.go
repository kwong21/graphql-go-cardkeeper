package service

import (
	"fmt"

	"github.com/kwong21/graphql-go-cardkeeper/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBClient interface {
	GetTeamsByLeague(string) []models.Team
	AddTeam(models.Team) (models.Team, error)
}

type PostgresClient struct {
	client *gorm.DB
}

func NewPGClient(config models.Config) (*PostgresClient, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		config.Database.Server, config.Database.User, config.Database.Password, config.Database.Database, config.Database.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	c := &PostgresClient{
		client: db,
	}

	c.runMigrations()
	return c, nil
}

func (pg PostgresClient) GetTeamsByLeague(league string) []models.Team {
	var teams []models.Team

	pg.client.Where("league = ?", league).Find(&teams)

	// log.Info("Found %d of teams in league %s", result.RowsAffected, league)

	return teams
}

func (pg PostgresClient) AddTeam(team models.Team) (models.Team, error) {
	result := pg.client.Create(&team)

	return team, result.Error
}

func (pg PostgresClient) runMigrations() {
	pg.client.AutoMigrate(&models.Team{})
}
