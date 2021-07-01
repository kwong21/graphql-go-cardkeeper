package resolver

import (
	"strconv"

	"github.com/kwong21/graphql-go-cardkeeper/models"
)

type PlayerQueryArgs struct {
	FirstName string
	LastName  string
	TeamName  string
	Team      TeamResolver
}

type PlayerResolver struct {
	p *models.Player
}

// ID resolves the unique identifier for the PLayer
func (pr *PlayerResolver) ID() *string {
	id := &pr.p.ID

	s := strconv.FormatUint(uint64(*id), 10)

	return &s
}

// FirstName resolves the given name of the Player
func (pr *PlayerResolver) FirstName() *string {
	return &pr.p.FirstName
}

// LastName resolves the family name of the Player
func (pr *PlayerResolver) LastName() *string {
	return &pr.p.LastName
}

// Team returns the name of the team the player plays for
func (pr *PlayerResolver) Team() *TeamResolver {
	tr := TeamResolver{
		t: &pr.p.Team,
	}

	return &tr
}
