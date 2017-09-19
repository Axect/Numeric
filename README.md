# Go Numeric

Go Package for Numerical Calculations & Statistics

## 0. Notation

* Vector is just []float64
* Matrix is just [][]float64

For example
```go
package main

import (
    "github.com/Axect/Numeric/array"
    "github.com/Axect/Numeric/stats"
)

func main() {
    A := array.Create(1, 1, 100) // [1, 2, 3, 4, ... , 100]
    m := stats.Mean(A) // Mean
    v := stats.Var(A) // Variance
    
    B := array.Create(-100, 1, -1)
    array.Write([]Vector{A, B}, "test.csv") // Write to csvfile
}
```

## 1. Array

Array has some vector operations.

* Create (Vector)
* Transpose (Matrix)
* Zeros (Matrix)
* Eyes (Matrix)
* Sum
* Pow
* Sub
* Add
* Mul
* Inner
* AddConst (Vector, float64)
* SubConst
* DivisionConst
* Float2Int ([]float64 -> []int64)
* Write ([]Vector -> csv)

## 2. Stats

Stats has statistics functions

* Mean
* Var
* Std
* Cov
* CovMatrix
* Cor (Correlation Coefficient)
* NormalDist