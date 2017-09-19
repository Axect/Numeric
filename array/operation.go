package array

import (
	"log"
	"math"
)

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
	if len(V) != len(W) {
		log.Fatal("Length is not matched!")
	}
	X := make(Vector, len(V), len(V))
	for i := range X {
		X[i] = V[i] + W[i]
	}
	return X
}

// Sub is Vector Vector operation
func Sub(V, W Vector) Vector {
	if len(V) != len(W) {
		log.Fatal("Length is not matched!")
	}
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
	if len(V) != len(W) {
		log.Fatal("Length is not matched!")
	}
	X := make(Vector, len(V), len(V))
	for i := range X {
		X[i] = V[i] * W[i]
	}
	return X
}
