package schema_test

import (
	"testing"

	"github.com/kwong21/graphql-go-cardkeeper/schema"
	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	s, err := schema.String()

	require.NoError(t, err)
	require.NotEmpty(t, s)
}
