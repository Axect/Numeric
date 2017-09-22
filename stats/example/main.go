package main

import (
	"fmt"

	"github.com/Axect/Numeric/array"
	"github.com/Axect/Numeric/stats"
)

// Vector is type alias
type Vector = array.Vector

func main() {
	X := stats.NormalDist(2, 1, 100000)
	Y := stats.NormalDist(5, 3, 100000)
	C := stats.CovMatrix(X, Y)
	R := stats.Cor(X, Y)
	array.MatrixForm(C)
	fmt.Println(R)

	D := stats.BinormialDist(1000, 1/6, 1000)
	m := stats.Mean(D)
	std := stats.Std(D)
	fmt.Println(m, std)
}
