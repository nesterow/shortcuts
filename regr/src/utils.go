package src

import (
	"gonum.org/v1/gonum/mat"
)

func Array2DToDense(X [][]float64) *mat.Dense {
	dense := mat.NewDense(len(X), len(X[0]), nil)
	for i, row := range X {
		dense.SetRow(i, row)
	}
	return dense
}
