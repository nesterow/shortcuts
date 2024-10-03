package src

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

type LogisticRegression struct {
	Epochs       int
	Weights      *mat.Dense
	Bias         float64
	LearningRate float64
}

func sigmoidFunction(x float64) float64 {
	if x < 0 {
		return 1. - 1./(1.+math.Exp(x))
	}
	return 1. / (1. + math.Exp(-x))
}

// binary cross-entropy Loss
func (regr *LogisticRegression) Loss(yTrue, yPred mat.Matrix) float64 {
	ep := 1e-9
	y1 := &mat.Dense{}
	y1.Apply(func(i, j int, v float64) float64 {
		return v * math.Log1p(yPred.At(i, j)+ep)
	}, yTrue)
	y2 := &mat.Dense{}
	y2.Apply(func(i, j int, v float64) float64 {
		return (1. - v) * math.Log1p(1.-yPred.At(i, j)+ep)
	}, yTrue)
	sum := &mat.Dense{}
	sum.Add(y1, y2)
	w, h := yTrue.Dims()
	return -(mat.Sum(sum) / float64(w*h))
}

func (regr *LogisticRegression) forward(X mat.Matrix) mat.Matrix {
	coef := &mat.Dense{}
	coef.Mul(X, regr.Weights)
	coef.Apply(func(i, j int, v float64) float64 {
		return sigmoidFunction(v + regr.Bias)
	}, coef)
	return coef
}

func (regr *LogisticRegression) grad(x, yTrue, yPred mat.Matrix) (*mat.Dense, float64) {
	nSamples, _ := x.Dims()
	deriv := &mat.Dense{}
	deriv.Sub(yPred, yTrue)
	dw := &mat.Dense{}
	dw.Mul(x.T(), deriv)
	dw.Apply(func(i, j int, v float64) float64 {
		return 1. / float64(nSamples) * v
	}, dw)
	db := (1. / float64(nSamples)) * mat.Sum(deriv)
	return dw, db
}

func (regr *LogisticRegression) backprop(x, y mat.Matrix) float64 {
	_, c := x.Dims()
	_, cy := y.Dims()
	if regr.Weights == nil {
		regr.Weights = mat.NewDense(c, cy, nil)
	}
	if regr.LearningRate == 0 {
		regr.LearningRate = 0.01
	}
	yPred := regr.forward(x)
	loss := regr.Loss(y, yPred)
	dw, db := regr.grad(x, y, yPred)
	regr.Weights.Sub(regr.Weights, dw)
	regr.Bias -= regr.LearningRate * db
	return loss
}

func (regr *LogisticRegression) Fit(X, Y mat.Matrix, losses *[]float64) {
	for i := 0; i < regr.Epochs; i++ {
		regr.backprop(X, Y)
		if losses != nil {
			*losses = append(*losses, regr.Loss(Y, regr.forward(X)))
		}
	}

}

func (regr *LogisticRegression) Predict(X mat.Matrix) mat.Matrix {
	coef := &mat.Dense{}
	coef.Mul(X, regr.Weights)
	coef.Apply(func(i, j int, v float64) float64 {
		p := sigmoidFunction(v + regr.Bias)
		if p > .5 {
			return 1.
		}
		return 0.
	}, coef)
	return coef
}
