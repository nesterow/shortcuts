package src

import (
	"fmt"
	"testing"
)

func TestLogisticRegression(t *testing.T) {
	X := [][]float64{{.1, .1, .1}, {.2, .2, .2}, {.1, .1, .1}, {.2, .2, .2}}
	Y := [][]float64{{0}, {1}, {0}, {1}}
	XDense := Array2DToDense(X)
	YDense := Array2DToDense(Y)
	epochs := 1000
	regr := &LogisticRegression{
		LearningRate: .1,
	}
	regr.Fit(XDense, YDense, epochs, nil)
	fmt.Println(regr.Weights, regr.Bias)
	yPred := regr.Predict(XDense)
	fmt.Println(YDense, yPred, regr.Loss(YDense, yPred))
}
