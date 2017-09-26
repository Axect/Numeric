package stats

import (
	"math"
	"math/rand"
	"time"
)

// NormalDist is Normal Distribution
func NormalDist(m, sigma float64, N int) Vector {
	V := make(Vector, N, N)
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := range V {
		V[i] = rand.NormFloat64()*sigma + m
	}
	return V
}

// BinomialDist is Binomial Distribution
func BinomialDist(n, p float64, N int) Vector {
	m := n * p
	std := math.Sqrt(n * p * (1 - p))
	return NormalDist(m, std, N)
}
