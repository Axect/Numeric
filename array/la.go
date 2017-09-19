package array

import (
	"log"
)

// Vector is Type Alias in Go 1.9
type Vector = []float64

// Matrix is Type Alias
type Matrix = []Vector

// Create generates arithmetic Sequence
func Create(init, step, end float64) Vector {
	Length := int((end - init + 1.) / step)
	A := make(Vector, Length, Length)
	s := init
	for i := range A {
		switch i {
		case 0:
			A[i] = s
		default:
			s += step
			A[i] = s
		}
	}
	return A
}

// Transpose transpose matrix
func Transpose(A Matrix) Matrix {
	Temp := Zeros(len(A[0]), len(A))

	// Transpose
	for i := range A {
		for j := range A[i] {
			Temp[j][i] = A[i][j]
		}
	}
	return Temp
}

// Det calc determinant
func Det(A Matrix) float64 {
	CheckSquare(A)
	Core := A[0]
	// Special Case : A = [number]
	if len(A) == 1 {
		return A[0][0]
	}
	s := 0.
	for i := range Core {
		B := Smallize(A, 0, i)
		switch {
		case i%2 == 0:
			s += Core[i] * Det(B)
		default:
			s -= Core[i] * Det(B)
		}
	}
	return s
}

// Smallize makes smaller matrix
func Smallize(A Matrix, m, n int) Matrix {
	if m >= len(A) || n >= len(A[0]) {
		log.Fatal("No Proper Input")
	}
	B := Zeros(len(A)-1, len(A[0])-1)
	for i := range B {
		for j := range B[i] {
			switch {
			case i < m && j < n:
				B[i][j] = A[i][j]
			case i >= m && j < n:
				B[i][j] = A[i+1][j]
			case i < m && j >= n:
				B[i][j] = A[i][j+1]
			case i >= m && j >= n:
				B[i][j] = A[i+1][j+1]
			}
		}
	}
	return B
}

// CheckSquare check square matrix
func CheckSquare(A Matrix) {
	if len(A) != len(A[0]) {
		log.Fatal("Should input Square Matrix!")
	}
}
