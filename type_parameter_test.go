package golanggenerics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestSample(t *testing.T) {
	var result string = Length[string]("Thomas")
	assert.Equal(t, "Thomas", result)

	var resultNumber int = Length[int](11)
	assert.Equal(t, 11, resultNumber)
}
