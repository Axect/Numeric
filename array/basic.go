package array

import "fmt"

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

// Zeros make M*N matrix
func Zeros(M, N int) Matrix {
	A := make(Matrix, M, M)
	for i := range A {
		A[i] = make(Vector, N, N)
	}
	return A
}

// Eyes generates diag matrix
func Eyes(N int) Matrix {
	I := Zeros(N, N)
	for i := range I {
		I[i][i] = 1.
	}
	return I
}

// MatrixForm print matrix form
func MatrixForm(A Matrix) {
	for i := range A {
		for j := range A[i] {
			switch {
			case j == len(A[i])-1:
				fmt.Println(A[i][j], " |")
			case j == 0:
				fmt.Printf("| %v\t", A[i][j])
			default:
				fmt.Printf("%v\t", A[i][j])
			}
		}
	}
}
