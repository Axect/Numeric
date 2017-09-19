package main

import (
	"fmt"

	"github.com/Axect/Numeric/array"
)

// Vector is type Alias
type Vector = []float64

func main() {
	A := array.Create(1, 1, 100)
	B := array.Create(-100, 1, -1)
	array.Write([]Vector{A, B}, "test.csv")
	E := array.Zeros(2, 2)
	E[0][0] = 1
	E[0][1] = 2
	E[1][0] = 3
	E[1][1] = 4
	fmt.Println(array.Det(E))
}
