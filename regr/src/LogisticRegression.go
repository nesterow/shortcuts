package src

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

type LogisticRegression struct {
	Epochs  int
	Weights *mat.Dense
	Bias    float64
	Losses  []float64
}

func sigmoidFunction(x float64) float64 {
	if x < 0 {
		return 1. - 1./(1.+math.Exp(x))
	}
	return 1. / (1. + math.Exp(-x))
}

func (regr *LogisticRegression) backprop(x, y mat.Matrix) float64 {
	_, c := x.Dims()
	ry, cy := y.Dims()
	regr.Bias = 0.1
	regr.Weights = mat.NewDense(cy, c, nil)
	coef := &mat.Dense{}

	coef.Mul(regr.Weights, x.T())
	coef.Apply(func(i, j int, v float64) float64 {
		return sigmoidFunction(v + regr.Bias)
	}, coef)

	diff := &mat.Dense{}
	diff.Sub(y.T(), coef)

	w := &mat.Dense{}
	w.Mul(diff, x)
	regr.Weights = w

	regr.Bias -= 0.1 * (mat.Sum(diff) / float64(c))

	// Loss
	yZeroLoss := &mat.Dense{}
	yZeroLoss.Apply(func(i, j int, v float64) float64 {
		return v * math.Log1p(coef.At(i, j)+1e-9)
	}, y.T())

	yOneLoss := &mat.Dense{}
	yOneLoss.Apply(func(i, j int, v float64) float64 {
		return (1. - v) * math.Log1p(1.-coef.At(i, j)+1e-9)
	}, y.T())

	sum := &mat.Dense{}
	sum.Add(yZeroLoss, yOneLoss)
	return mat.Sum(sum) / float64(ry+cy)
}

func (regr *LogisticRegression) Fit(X, Y mat.Matrix, epochs int) error {
	for i := 0; i < epochs; i++ {
		loss := regr.backprop(X, Y)
		regr.Losses = append(regr.Losses, loss)
	}
	return nil
}

func (regr *LogisticRegression) Predict(X mat.Matrix) mat.Matrix {
	coef := &mat.Dense{}
	coef.Mul(X, regr.Weights.T())
	coef.Apply(func(i, j int, v float64) float64 {
		p := sigmoidFunction(v + regr.Bias)
		if p > .5 {
			return 1.
		}
		return 0.
	}, coef)
	return coef
}
