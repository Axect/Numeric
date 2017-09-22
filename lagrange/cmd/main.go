package main

import (
	"fmt"
	"os/exec"
	"sync"

	"github.com/Axect/Numeric/array"
	"github.com/Axect/Numeric/lagrange"
)

var wg sync.WaitGroup

// Vector is just []float64
type Vector = array.Vector

func main() {
	X := Vector{1., 2., 3., 4.}
	Y := Vector{1., 4., 9., 16.}
	L := lagrange.LPolynomial(X, Y)
	T := array.Create(0., 0.01, 400)
	Q := make(Vector, len(T), len(T))
	for i, t := range T {
		Q[i] = L(t)
	}
	array.Write([]Vector{T, Q}, "Data/Lagrange.csv")

	wg.Add(1)
	go Routine("cmd/plot.jl")
	wg.Wait()

	fmt.Println("All Process Finished!")
}

// Routine for run julia
func Routine(Directory string) {
	defer wg.Done()

	Julia := "julia"

	var (
		cmdOut []byte
		err    error
	)

	if cmdOut, err = exec.Command(Julia, Directory).Output(); err != nil {
		panic(err)
	}
	comp := string(cmdOut)
	fmt.Println(comp)
	fmt.Println("Complete to Plot!")
	fmt.Println()
	return
}
