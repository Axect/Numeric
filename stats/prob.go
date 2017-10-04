package stats

import (
	"math"
)

// Dist is type alias for dist
type Dist = func(float64) float64

// Factorial is factorial
func Factorial(n float64) float64 {
	p := 1.
	for i := 1; i <= int(n+0.4); i++ {
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

// Binomial is probability for binomial
func Binomial(N, p float64) Dist {
	s := 0.
	q := float64(1) - p
	var r float64
	return func(n float64) float64 {
		for i := 0; i <= int(n+0.1); i++ {
			r = float64(i)
			s += Comb(N, r) * math.Pow(p, r) * math.Pow(q, N-r)
		}
		return s
	}
}

// Poisson calc poisson distribution's probability
func Poisson(mu, n float64) float64 {
	return math.Pow(mu, n) / Factorial(n) * math.Exp(-n)
}
