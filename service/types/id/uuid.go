package id

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
)

// MarshalUUID returns a GraphQL Marshaler for the uuid.UUID type.
func MarshalUUID(uuid uuid.UUID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		// Write the UUID as a quoted string to the GraphQL response.
		_, _ = io.WriteString(w, strconv.Quote(uuid.String()))
	})
}

// UnmarshalUUID returns a uuid.UUID from the provided GraphQL variable value.
func UnmarshalUUID(variableValue interface{}) (u uuid.UUID, err error) {
	// Check that the variable value is a string.
	s, ok := variableValue.(string)
	if !ok {
		// Return an error if the variable value is not a string.
		return u, fmt.Errorf("invalid type %T, expected string", variableValue)
	}

	// Parse the string into a UUID.
	return uuid.Parse(s)
}
