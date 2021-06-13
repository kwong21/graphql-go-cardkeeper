package resolver

import (
	"strconv"

	"github.com/kwong21/graphql-go-cardkeeper/models"
)

type TeamQueryArgs struct {
	Name   string
	League string
	Abbr   string
}
type TeamResolver struct {
	t *models.Team
}

// ID resolves the unique identifier for the Team
func (tr *TeamResolver) ID() *string {
	id := &tr.t.ID

	s := strconv.FormatUint(uint64(*id), 10)

	return &s
}

// Name resolves the name of the Team
func (tr *TeamResolver) Name() *string {
	return &tr.t.Name
}

// Abbr resolves the abbrieviated name of the Team
func (tr *TeamResolver) Abbr() *string {
	return &tr.t.Abbr
}

// League returns the professional league the Team belongs to
func (tr *TeamResolver) League() *string {
	return &tr.t.League
}
