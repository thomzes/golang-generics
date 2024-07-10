package golanggenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func IsSame[T comparable](value1, value2 T) bool {
	if value1 == value2 {
		return true
	} else {
		return false
	}
}

func TestIsMe(t *testing.T) {
	assert.True(t, IsSame[int](1, 1))
	assert.True(t, IsSame[string]("1", "1"))
	assert.True(t, IsSame[bool](true, true))
}
