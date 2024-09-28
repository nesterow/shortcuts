package main

import (
	"fmt"

	"golang.org/x/exp/rand"

	"gonum.org/v1/gonum/stat"
)

//go:export example
func Example() {
	var (
		xs      = make([]float64, 100)
		ys      = make([]float64, 100)
		weights []float64
	)

	line := func(x float64) float64 {
		return 1 + 3*x
	}

	for i := range xs {
		xs[i] = float64(i)
		ys[i] = line(xs[i]) + 0.1*rand.NormFloat64()
	}

	// Do not force the regression line to pass through the origin.
	origin := false

	alpha, beta := stat.LinearRegression(xs, ys, weights, origin)
	r2 := stat.RSquared(xs, ys, weights, alpha, beta)

	fmt.Printf("Estimated offset is: %.6f\n", alpha)
	fmt.Printf("Estimated slope is:  %.6f\n", beta)
	fmt.Printf("R^2: %.6f\n", r2)

}

func main() {
	fmt.Println("Stats initialized")
}
