package stats

// Factorial is factorial
func Factorial(n float64) float64 {
	p := 1.
	for i := 1; i <= int(n); i++ {
		p *= float64(i)
	}
	return p
}

// Perm is permutation
func Perm(n, r float64) float64 {
	nfac := Factorial(n)
	rfac := Factorial(n - r)
	return nfac / rfac
}

// Comb is combination
func Comb(n, r float64) float64 {
	perm := Perm(n, r)
	rfac := Factorial(r)
	return perm / rfac
}

// Homonomial is Combination with repetition
func Homonomial(n, r float64) float64 {
	return Comb(n+r-1, r)
}
