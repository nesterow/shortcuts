//go:build js && wasm
// +build js,wasm

package src

import (
	"syscall/js"
)

func NewLogisticRegressionJS(this js.Value, args []js.Value) interface{} {
	var (
		epochs       = args[0].Get("epochs").Int()
		learningRate = args[0].Get("learningRate").Float()
	)
	reg := &MCLogisticRegression{
		Epochs:       epochs,
		LearningRate: learningRate,
	}
	loss := 1.0
	obj := js.Global().Get("Object").New()
	obj.Set("fit", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		X := JSFloatArray2D(args[0])
		Y := JSFloatArray2D(args[1])
		XDense := Array2DToDense(X)
		YDense := Array2DToDense(Y)
		reg.Fit(XDense, YDense)
		return nil
	}))
	obj.Set("predict", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		X := JSFloatArray2D(args[0])
		XDense := Array2DToDense(X)
		Y := reg.Predict(XDense)
		loss = reg.Loss(XDense, Y)
		return MatrixToJSArray2D(Y)
	}))
	obj.Set("loss", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return loss
	}))
	return obj
}
