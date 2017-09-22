package main

import (
	"github.com/Axect/Numeric/Lagrange"
	"github.com/Axect/Numeric/array"
)

type Vector = array.Vector

func main() {
	X := Vector{1., 2., 3., 4.}
	Y := Vector{1., 4., 9., 16.}
	L := Lagrange.LPolynomial(X, Y)
	T := array.Create(0., 0.01, 400)
	Q := make(Vector, len(T), len(T))
	for i, t := range T {
		Q[i] = L(t)
	}
	array.Write([]Vector{X, Y, T, Q}, "Data/Lagrange.csv")
}
