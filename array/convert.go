package array

// Float2Int converts []float64 to []int
func Float2Int(A []float64) []int {
	B := make([]int, len(A), len(A))
	for i, elem := range A {
		B[i] = int(elem)
	}
	return B
}
