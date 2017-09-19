package main

import (
	"fmt"

	"github.com/Axect/Numeric/array"
	"github.com/Axect/Numeric/stats"
)

// Vector is type alias
type Vector = array.Vector

func main() {
	X := stats.NormalDist(2, 1, 1000)
	Y := stats.NormalDist(5, 3, 1000)
	mx, my := stats.Mean(X), stats.Mean(Y)
	sx, sy := stats.Std(X), stats.Std(Y)
	C := stats.CovMatrix(X, Y)
	R := stats.Cor(X, Y)
	fmt.Println(mx, my)
	fmt.Println(sx, sy)
	fmt.Println(C)
	fmt.Println(R)
}
