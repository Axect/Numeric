package stats

import (
	"github.com/Axect/Numeric/array"
)

type Vector = []float64

func Mean(V Vector) float64 {
	s := array.Sum(V)
	N := len(V)
	return s / float64(N)
}
