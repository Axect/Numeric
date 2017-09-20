package array

import (
	"log"
	"math"
)

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
		B := Minor(A, 0, i)
		switch {
		case i%2 == 0:
			s += Core[i] * Det(B)
		default:
			s -= Core[i] * Det(B)
		}
	}
	return s
}

// Minor makes smaller matrix
func Minor(A Matrix, m, n int) Matrix {
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

func Inverse(A Matrix) Matrix {
	CheckInv(A)
	B := Zeros(len(A), len(A[0]))
	D := math.Abs(Det(A))
	// Row
	for i := range A {
		for j := range A[i] {
			B[i][j] = math.Pow(-1., float64(i+j+1)) / D * Det(Minor(A, j, i))
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

// CheckInv check invertible
func CheckInv(A Matrix) {
	if Det(A) == 0 {
		log.Fatal("Singular Matrix!")
	}
}
