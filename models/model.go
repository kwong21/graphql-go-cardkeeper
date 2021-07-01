package models

import (
	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	Name   string `gorm:"not null;uniqueIndex:idx_team;default:null"`
	Abbr   string `gorm:"not null;uniqueIndex:idx_team;default:null"`
	League string
}

type Player struct {
	gorm.Model
	FirstName string `gorm:"not null;uniqueIndex:idx_player;default:null"`
	LastName  string `gorm:"not null;uniqueIndex:idx_player;default:null"`
	TeamID    uint
	Team      Team
}

type Config struct {
	Database database
}

type database struct {
	Server   string
	User     string
	Password string
	Database string
	Port     string
}
