//go:build js && wasm
// +build js,wasm

package src

import (
	"syscall/js"

	"gonum.org/v1/gonum/mat"
)

type LinearRegression struct {
	Coef *mat.Dense
}

func (reg *LinearRegression) Fit(X, Y [][]float64) error {
	XDense := Array2DToDense(X)
	YDense := Array2DToDense(Y)
	reg.Coef = new(mat.Dense)
	reg.Coef.Solve(XDense, YDense)
	return nil
}

func (reg *LinearRegression) Predict(X [][]float64) ([]float64, error) {
	YDense := new(mat.Dense)
	YDense.Mul(Array2DToDense(X), reg.Coef)
	return YDense.RawMatrix().Data, nil
}

func (reg *LinearRegression) Score(YT, Y [][]float64) float64 {
	TDense := Array2DToDense(YT)
	YDense := Array2DToDense(Y)
	return R2Score(TDense, YDense, nil, "raw_values").At(0, 0)
}

func (l *LinearRegression) Save() ([]byte, error) {
	return nil, nil
}

func (l *LinearRegression) Load(data []byte) error {
	return nil
}

func NewLinearRegressionJS(this js.Value, args []js.Value) interface{} {
	reg := new(LinearRegression)
	obj := js.Global().Get("Object").New()
	obj.Set("fit", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		X := JSFloatArray2D(args[0])
		Y := JSFloatArray2D(args[1])
		return reg.Fit(X, Y)
	}))
	obj.Set("predict", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		X := JSFloatArray2D(args[0])
		Y, _ := reg.Predict(X)
		return ToJSArray(Y)
	}))
	obj.Set("score", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		X := JSFloatArray2D(args[0])
		Y := JSFloatArray2D(args[1])
		return reg.Score(X, Y)
	}))
	return obj
}
