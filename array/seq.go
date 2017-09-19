package array

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
	Temp := make(Matrix, len(A[0]), len(A[0]))
	// Make Space
	for j := range Temp {
		Temp[j] = make(Vector, len(A), len(A))
	}

	// Transpose
	for i := range A {
		for j := range A[i] {
			Temp[j][i] = A[i][j]
		}
	}
	return Temp
}
