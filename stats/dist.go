package stats

import (
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
