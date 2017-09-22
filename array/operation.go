package array

import (
	"log"
	"math"
)

// Sum sums all
func Sum(V Vector) float64 {
	s := 0.
	for _, elem := range V {
		s += elem
	}
	return s
}

// Pow generate Powered Vector
func Pow(V Vector, n float64) Vector {
	W := make(Vector, len(V), len(V))
	for i := range V {
		W[i] = math.Pow(V[i], n)
	}
	return W
}

// SubConst is Vector scalar operation
func SubConst(V Vector, x float64) Vector {
	W := make(Vector, len(V), len(V))
	for i := range W {
		W[i] = V[i] - x
	}
	return W
}

// AddConst is Vector scalar operation
func AddConst(V Vector, x float64) Vector {
	W := make(Vector, len(V), len(V))
	for i := range W {
		W[i] = V[i] + x
	}
	return W
}

// Add is Vector Vector operation
func Add(V, W Vector) Vector {
	LengthCheck(V, W)
	X := make(Vector, len(V), len(V))
	for i := range X {
		X[i] = V[i] + W[i]
	}
	return X
}

// Sub is Vector Vector operation
func Sub(V, W Vector) Vector {
	LengthCheck(V, W)
	X := make(Vector, len(V), len(V))
	for i := range X {
		X[i] = V[i] - W[i]
	}
	return X
}

// DivisionConst is Vector Scalar operation
func DivisionConst(V Vector, s float64) Vector {
	W := make(Vector, len(V), len(V))
	if s == 0. {
		log.Fatal("Zero Division Error!")
	}
	for i := range W {
		W[i] = V[i] / s
	}
	return W
}

// Mul is Vector Vector operation
func Mul(V, W Vector) Vector {
	LengthCheck(V, W)
	X := make(Vector, len(V), len(V))
	for i := range X {
		X[i] = V[i] * W[i]
	}
	return X
}

// Inner is inner product
func Inner(V, W Vector) float64 {
	LengthCheck(V, W)
	s := 0.
	for i := range V {
		s += V[i] * W[i]
	}
	return s
}

// LengthCheck generates error
func LengthCheck(V, W Vector) {
	if len(V) != len(W) {
		log.Fatal("Length is not matched!")
	}
}

// Product products array
func Product(V Vector) float64 {
	p := 1.
	for i := range V {
		p *= V[i]
	}
	return p
}
