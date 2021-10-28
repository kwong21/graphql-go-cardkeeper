package models

import (
	"strconv"

	"gorm.io/gorm"
)

// PlayerInputArgs contains the fields for adding a Player to the API
type PlayerInputArgs struct {
	TeamID    string
	FirstName string
	LastName  string
}

// Player encapsulates information for a Player
type Player struct {
	gorm.Model
	FirstName string `gorm:"not null;uniqueIndex:idx_player;default:null"`
	LastName  string `gorm:"not null;uniqueIndex:idx_player;default:null"`
	TeamID    uint
	TeamName  string
	Team      Team `gorm:"embedded"`
}

// PlayerResolver contains the resolved GraphQL representation of the Player
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

// TeamName resolves the name of the team the Player plays for
func (Pr *PlayerResolver) TeamName() *string {
	return nullableValueOf(Pr.P.TeamName)
}

// Team returns the name of the team the player plays for
func (Pr *PlayerResolver) Team() *TeamResolver {
	tr := TeamResolver{
		T: &Pr.P.Team,
	}

	return &tr
}
