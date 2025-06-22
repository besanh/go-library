package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUn(t *testing.T) {
	u := NewUtil()

	assert.NotNil(t, u)
	assert.IsType(t, &Util{}, u)
}
