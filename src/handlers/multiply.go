package handlers

import (
	"fmt"
	"math/big"
	"net/http"

	"../helpers"
)

// Multiply route to get a product of the integers in the matrix
func Multiply(responseWriter http.ResponseWriter, request *http.Request) {
	chMatrix, _ := helpers.GetMatrix(helpers.MULTIPLY, responseWriter, request)
	count, _ := new(big.Int).SetString("1", 10)

	for n := range chMatrix {
		nBig := helpers.IntToBigInt(n)
		count = count.Mul(count, nBig)
	}

	fmt.Fprint(responseWriter, count)
}
