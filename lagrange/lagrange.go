package lagrange

import (
	"log"

	"github.com/Axect/Numeric/array"
)

// Normal means normal float64 function
type Normal func(float64) float64

// Vector is type alias
type Vector = []float64

// Mul generates new normal function by product with float number
func Mul(f Normal, x float64) Normal {
	return func(z float64) float64 {
		return f(z) * x
	}
}

// Lagrange is Lagrange interpolation
func Lagrange(xseq Vector, i int) Normal {
	if i >= len(xseq) || i < 0 {
		log.Fatal("Input proper integer") // Trivial Error Filter (index not matched)
	}
	lapl := func(x float64) float64 {
		tempseq := make(Vector, len(xseq), len(xseq))
		for j := 0; j < len(tempseq); j++ {
			if xseq[j] != xseq[i] {
				tempseq[j] = (x - xseq[j]) / (xseq[i] - xseq[j])
			} else {
				tempseq[j] = 1.
			}
		}
		return array.Product(tempseq)
	}
	return lapl
}

// LPolynomial means Lagrange polynomial
func LPolynomial(xseq, yseq Vector) Normal {
	var templapl Normal
	result := func(x float64) float64 {
		s := 0.
		for i, yval := range yseq {
			templapl = Lagrange(xseq, i)
			templapl = Mul(templapl, yval)
			s += templapl(x)
		}
		return s
	}
	return result
}
