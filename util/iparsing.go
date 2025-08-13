package util

import (
	"fmt"

	"github.com/google/uuid"
)

// UUIDPtrToStringPtr converts a uuid.UUID to a *string.
// Returns nil if the input is nil.
func (i *Util) UUIDPtrToStringPtr(u uuid.UUID) *string {
	s := u.String()
	return &s
}

// StringPtrToUUIDPtr parses a string (containing a UUID) back to a *uuid.UUID.
// Returns (nil, nil) if the input is nil, or an error if the string isn’t a valid UUID.
func (i *Util) StringPtrToUUIDPtr(s string) (*uuid.UUID, error) {
	u, err := uuid.Parse(s)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID %q: %w", s, err)
	}
	return &u, nil
}
