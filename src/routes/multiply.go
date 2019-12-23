package routes

import (
	"fmt"
	"math/big"
	"net/http"
	"strconv"

	"../helpers"
)

// Multiply route to get a product of the integers in the matrix
func Multiply(responseWriter http.ResponseWriter, request *http.Request) {
	matrix, _ := helpers.GetMatrix(helpers.MULTIPLY, responseWriter, request)

	count, _ := new(big.Int).SetString("1", 10)

	for n := range matrix {
		nString := strconv.FormatInt(int64(n), 10)
		nBig, _ := new(big.Int).SetString(nString, 10)
		count = count.Mul(count, nBig)
	}

	fmt.Fprint(responseWriter, count)
}
