package helpers

import (
	"fmt"
	"reflect"
	"testing"
)

func Testtranspose(t *testing.T) {
	matrix := [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}

	expected := [][]string{
		{"1", "4", "7"},
		{"2", "5", "8"},
		{"3", "6", "9"},
	}

	_, transposed := transpose(matrix)

	fmt.Println(transposed)

	if !reflect.DeepEqual(expected, transposed) {
		t.ErrorF("transponsed [][]string does not match with the expected")
	}
}
