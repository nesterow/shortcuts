package src

import (
	"fmt"
	"testing"
)

func TestLogisticRegression(t *testing.T) {
	X := [][]float64{{10.1, 10.1, 10.1}, {2.1, 2.1, 2.1}, {10.2, 10.2, 10.2}, {2.2, 2.2, 2.2}}
	Y := [][]float64{{0}, {1}, {0}, {1}}
	XDense := Array2DToDense(X)
	YDense := Array2DToDense(Y)
	epochs := 10
	regr := &LogisticRegression{}
	err := regr.Fit(XDense, YDense, epochs)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(regr.Weights, regr.Bias, regr.Losses)
	fmt.Println(YDense, regr.Predict(XDense))
}
