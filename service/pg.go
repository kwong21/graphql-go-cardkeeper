package service

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/kwong21/graphql-go-cardkeeper/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

type DBClient interface {
	GetTeamsByLeague(string) []models.Team
	GetTeamByName(string) ([]models.Team, error)
	GetPlayerByName(string, string) ([]models.Player, error)
	AddTeam(models.Team) (models.Team, error)
	AddPlayer(models.Player) (models.Player, error)
}

type PostgresClient struct {
	client *gorm.DB
}

func NewPGClient(config models.Config) (*PostgresClient, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		config.Database.Server, config.Database.User, config.Database.Password, config.Database.Database, config.Database.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      logger.Info,
				Colorful:      false,
			},
		),
	})

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

	pg.client.Where("league = ?", league).First(&teams)

	return teams
}

func (pg PostgresClient) AddTeam(team models.Team) (models.Team, error) {
	result := pg.client.Create(&team)

	return team, result.Error
}

func (pg PostgresClient) GetPlayerByName(firstName string, lastName string) ([]models.Player, error) {
	var players []models.Player

	r := pg.client.Where("first_name = ? AND last_name = ?", firstName, lastName).Find(&players)

	return players, r.Error
}

func (pg PostgresClient) AddPlayer(player models.Player) (models.Player, error) {
	result := pg.client.Omit(clause.Associations).Create(&player)

	return player, result.Error
}

func (pg PostgresClient) GetTeamByName(teamName string) ([]models.Team, error) {
	var teams []models.Team

	r := pg.client.Where("name = ?", teamName).Find(&teams)

	return teams, r.Error
}

func (pg PostgresClient) runMigrations() {
	pg.client.AutoMigrate(&models.Team{}, &models.Player{})
}
