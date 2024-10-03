package src

import (
	"fmt"
	"testing"
)

func TestLogisticRegression(t *testing.T) {
	X := [][]float64{
		{.1, .1, .1},
		{.2, .2, .2},
		{.1, .1, .1},
		{.2, .2, .2},
	}
	//Y := [][]float64{{1}, {0}, {1}, {0}}
	Y := [][]float64{
		{0},
		{1},
		{0},
		{1},
	}
	XDense := Array2DToDense(X)
	YDense := Array2DToDense(Y)
	epochs := 1000
	regr := &LogisticRegression{
		Epochs:       epochs,
		LearningRate: .01,
	}
	regr.Fit(XDense, YDense, nil)
	fmt.Println(regr.Weights, regr.Bias)
	yPred := regr.Predict(XDense)
	fmt.Println(YDense, yPred, regr.Loss(YDense, yPred))
}

func TestMCLogisticRegression(t *testing.T) {
	X := [][]float64{
		{.1, .1, .1},
		{.2, .2, .2},
		{.1, .1, .1},
		{.2, .2, .2},
	}
	//Y := [][]float64{{1}, {0}, {1}, {0}}
	Y := [][]float64{
		{0, 1},
		{1, 0},
		{0, 1},
		{1, 0},
	}
	XDense := Array2DToDense(X)
	YDense := Array2DToDense(Y)
	epochs := 1000
	regr := &MCLogisticRegression{
		Epochs:       epochs,
		LearningRate: .01,
	}
	regr.Fit(XDense, YDense)
	// fmt.Println(regr.Weights, regr.Bias)
	yPred := regr.Predict(XDense)
	fmt.Println(YDense, yPred) //, regr.Loss(YDense, yPred))
}
