package handlers

import (
	"fmt"
	"math/big"
	"net/http"
	"sync"

	"../helpers"
)

// Multiply route to get a product of the integers in the matrix
func Multiply(responseWriter http.ResponseWriter, request *http.Request) {
	matrix, _ := helpers.GetMatrix(helpers.MULTIPLY, responseWriter, request)
	count, _ := new(big.Int).SetString("1", 10)
	wg := sync.WaitGroup{}

	for n := range matrix {
		wg.Add(1)

		go func(n int) {
			defer wg.Done()

			nBig := helpers.IntToBigInt(n)
			count = count.Mul(count, nBig)
		}(n)
	}

	wg.Wait()

	fmt.Fprint(responseWriter, count)
}
