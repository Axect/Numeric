package array

// Eyes generates diag matrix
func Eyes(N int) Matrix {
	I := Zeros(N, N)
	for i := range I {
		I[i][i] = 1.
	}
	return I
}
