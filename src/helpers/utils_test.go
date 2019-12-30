package helpers

import (
	"math/big"
	"reflect"
	"testing"
)

func TestIntToBigInt(t *testing.T) {
	t.Parallel()

	bigInt := IntToBigInt(1)
	expected, _ := new(big.Int).SetString("1", 10)

	if reflect.TypeOf(bigInt) != reflect.TypeOf(expected) {
		t.Errorf("Wrong conversion")
	}
}

func TestTranspose(t *testing.T) {
	t.Parallel()

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

	transposed := Transpose(matrix)

	if !reflect.DeepEqual(expected, transposed) {
		t.Errorf("transponsed [][]string does not match with the expected")
	}
}

func TestMatrixToChannel(t *testing.T) {
	t.Parallel()

	matrix := [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}

	ch := matrixToChannel(matrix)
	chResult := []int{}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	if reflect.TypeOf(ch).String() != "<-chan int" {
		t.Errorf("matrixToChannel does not return <-chan int")
	}

	for n := range ch {
		chResult = append(chResult, n)
	}

	if !reflect.DeepEqual(chResult, expected) {
		t.Errorf("<-chan int does not represent matrix [][]string values")
	}

}
