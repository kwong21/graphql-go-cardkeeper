package service

import (
	"errors"
	"fmt"

	"github.com/kwong21/graphql-go-cardkeeper/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DBClient interface {
	GetTeamsByLeague(string) []models.Team
	GetTeamByName(string) (models.Team, error)
	GetPlayerByName(string, string) models.Player
	AddTeam(models.Team) (models.Team, error)
	AddPlayer(models.Player) (models.Player, error)
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

func (pg PostgresClient) GetPlayerByName(firstName string, lastName string) models.Player {
	var player models.Player

	pg.client.Where("firstName = ? AND lastName = ?", firstName, lastName).Find(&player)

	return player
}

func (pg PostgresClient) AddPlayer(player models.Player) (models.Player, error) {
	result := pg.client.Omit(clause.Associations).Create(&player)

	return player, result.Error
}

func (pg PostgresClient) GetTeamByName(teamName string) (models.Team, error) {
	var team models.Team

	pg.client.Where("name = ?", teamName).Find(&team)

	if team.ID <= 0 {
		return team, errors.New(fmt.Sprintf("team with name %s does not exist", teamName))
	}

	return team, nil
}

func (pg PostgresClient) runMigrations() {
	pg.client.AutoMigrate(&models.Team{}, &models.Player{})
}
