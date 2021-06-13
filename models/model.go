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
