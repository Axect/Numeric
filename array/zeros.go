package array

// Zeros make M*N matrix
func Zeros(M, N int) Matrix {
	A := make(Matrix, M, M)
	for i := range A {
		A[i] = make(Vector, N, N)
	}
	return A
}