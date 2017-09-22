package array

import "math"

// =============================================================================
// Vector Function
// =============================================================================

// Pow generate Powered Vector
func Pow(V Vector, n float64) Vector {
	W := make(Vector, len(V), len(V))
	for i := range V {
		W[i] = math.Pow(V[i], n)
	}
	return W
}

// Sin generate Sin Vector
func Sin(V Vector) Vector {
	W := make(Vector, len(V), len(V))
	for i := range V {
		W[i] = math.Sin(V[i])
	}
	return W
}

// Cos generate Cos Vector
func Cos(V Vector) Vector {
	W := make(Vector, len(V), len(V))
	for i := range V {
		W[i] = math.Cos(V[i])
	}
	return W
}

// Tan generate Tan Vector
func Tan(V Vector) Vector {
	W := make(Vector, len(V), len(V))
	for i := range V {
		W[i] = math.Tan(V[i])
	}
	return W
}

// Exp generate Exp Vector
func Exp(V Vector) Vector {
	W := make(Vector, len(V), len(V))
	for i := range V {
		W[i] = math.Exp(V[i])
	}
	return W
}
