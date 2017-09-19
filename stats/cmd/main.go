package main

import (
	"fmt"

	"github.com/Axect/Numeric/array"
	"github.com/Axect/Numeric/stats"
)

// Vector is type alias
type Vector = array.Vector

func main() {
	X := Vector{0., 2., 1., -3.}
	Y := Vector{2., 5., 4., -1.}
	fmt.Println(stats.Var(X), stats.Var(Y), stats.Cov(X, Y))
	fmt.Println(stats.CovMatrix(X, Y))
	fmt.Println(stats.Cor(X, Y))
}
