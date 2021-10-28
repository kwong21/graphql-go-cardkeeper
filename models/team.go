package models

import (
	"strconv"

	"gorm.io/gorm"
)

// Team contains information for a sports team
type Team struct {
	gorm.Model
	Name   string `gorm:"not null;uniqueIndex:idx_team;default:null"`
	Abbr   string `gorm:"not null;uniqueIndex:idx_team;default:null"`
	League string
}

// TeamResolver is the struct representation of the GraphQL object
type TeamResolver struct {
	T *Team
}

// ID resolves the unique identifier for the Team
func (Tr *TeamResolver) ID() *string {
	id := &Tr.T.ID

	s := strconv.FormatUint(uint64(*id), 10)

	return &s
}

// Name resolves the name of the Team
func (Tr *TeamResolver) Name() *string {
	return &Tr.T.Name
}

// Abbr resolves the abbrieviated name of the Team
func (Tr *TeamResolver) Abbr() *string {
	return &Tr.T.Abbr
}

// League returns the professional league the Team belongs to
func (Tr *TeamResolver) League() *string {
	return &Tr.T.League
}
