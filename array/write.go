package array

import (
	"github.com/Axect/csv"
)

// Write has a format : V = [X, Y, Z] where X,Y,Z are Vectors
func Write(V []Vector, name string) {
	Txt := csv.Float2Str(V)
	Tr := csv.Transpose(Txt)
	csv.Write(Tr, name)
}
