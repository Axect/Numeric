package array

// Eyes generates diag matrix
func Eyes(N int) Matrix {
	I := make(Matrix, N, N)
	for i := range I {
		I[i] = make(Vector, N, N)
	}

	for i := range I {
		I[i][i] = 1.
	}
	return I
}
