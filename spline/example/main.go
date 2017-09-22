package main

import (
	"fmt"
	"os/exec"
	"sync"

	"github.com/Axect/Numeric/array"
	"github.com/Axect/Numeric/spline"
)

// Vector is alias
type Vector = []float64

var wg sync.WaitGroup

func main() {
	// Create Array
	X := array.Create(0, 1, 40)
	Y := array.Sin(X)
	// Spline
	Sp := spline.NewCubic(X, Y)

	T := array.Create(0, .1, 40)
	S := Sp.Evaluate(T)
	Result := []Vector{T, S}
	array.Write(Result, "Data/test.csv")
	wg.Add(1)
	go Routine()
	wg.Wait()
	fmt.Println()
	fmt.Println("All Process Finished!")

}

// Routine for julia
func Routine() {
	defer wg.Done()

	var (
		cmdOut []byte
		err    error
	)

	if cmdOut, err = exec.Command("julia", "plot.jl").Output(); err != nil {
		panic(err)
	}
	comp := string(cmdOut)
	fmt.Println()
	fmt.Println(comp)
	fmt.Println("Plot Finished!")
	return
}
