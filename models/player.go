package models

import (
	"strconv"

	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	FirstName string `gorm:"not null;uniqueIndex:idx_player;default:null"`
	LastName  string `gorm:"not null;uniqueIndex:idx_player;default:null"`
	TeamID    uint
	Team      Team `gorm:"embedded"`
}

type PlayerResolver struct {
	P *Player
}

// ID resolves the unique identifier for the PLayer
func (Pr *PlayerResolver) ID() *string {
	id := &Pr.P.ID

	s := strconv.FormatUint(uint64(*id), 10)

	return &s
}

// FirstName resolves the given name of the Player
func (Pr *PlayerResolver) FirstName() *string {
	return &Pr.P.FirstName
}

// LastName resolves the family name of the Player
func (Pr *PlayerResolver) LastName() *string {
	return &Pr.P.LastName
}

// Team returns the name of the team the player plays for
func (Pr *PlayerResolver) Team() *TeamResolver {
	tr := TeamResolver{
		T: &Pr.P.Team,
	}

	return &tr
}
