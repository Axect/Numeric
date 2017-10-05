package main

import (
	"fmt"

	ar "github.com/Axect/Numeric/array"
	ss "github.com/Axect/Numeric/stats"
)

// Vector is type alias
type Vector = ar.Vector

func main() {
	X := ss.NormalDist(2, 1, 100000)
	Y := ss.NormalDist(5, 3, 100000)
	C := ss.CovMatrix(X, Y)
	R := ss.Cor(X, Y)
	ar.MatrixForm(C)
	fmt.Println(R)

	D := ss.Binomial(10, 1./5)
	fmt.Println(D(1))
}
