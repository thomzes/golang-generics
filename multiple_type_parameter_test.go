package golanggenerics

import (
	"fmt"
	"testing"
)

func MultipleTypeParameter[T1 any, T2 any](param1 T1, param2 T2) {
	fmt.Println(param1)
	fmt.Println(param2)
}

func TestMultipleTypeParameter(t *testing.T) {
	MultipleTypeParameter[string, int]("Thomas", 11)
	MultipleTypeParameter[int, string](100, "Bambang")
}
