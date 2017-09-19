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

	C := array.Zeros(10, 5)
	fmt.Println(C)
}
