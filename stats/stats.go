package stats

import (
	"math"

	"github.com/Axect/Numeric/array"
)

// Vector is type alias
type Vector = []float64

// Matrix is type alias
type Matrix = []Vector

// Mean is function for vector
func Mean(V Vector) float64 {
	s := array.Sum(V)
	N := len(V)
	return s / float64(N)
}

// Var is variance
func Var(V Vector) float64 {
	m := Mean(V)
	X := array.SubConst(V, m)
	Y := array.Pow(X, 2)
	s := array.Sum(Y)
	return s / float64(len(V)-1)
}

// Std is std
func Std(V Vector) float64 {
	v := Var(V)
	return math.Sqrt(v)
}

// CovMatrix is cov matrix
func CovMatrix(V, W Vector) Matrix {
	x1, y2 := Var(V), Var(W)
	x2 := Cov(V, W)
	y1 := x2
	Result := Matrix{Vector{x1, x2}, Vector{y1, y2}}
	return Result
}

// Cov is covariance
func Cov(V, W Vector) float64 {
	mx, my := Mean(V), Mean(W)
	X, Y := array.SubConst(V, mx), array.SubConst(W, my)
	Z := array.Mul(X, Y)
	S := array.Sum(Z)
	return S / float64(len(V)-1)
}

// Cor is Correlation Coefficient
func Cor(V, W Vector) float64 {
	s2 := Cov(V, W)
	sx, sy := Std(V), Std(W)
	return s2 / (sx * sy)
}
