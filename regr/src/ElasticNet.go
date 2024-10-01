//go:build js && wasm
// +build js,wasm

package src

import (
	"syscall/js"

	"gonum.org/v1/gonum/mat"
)

type ElasticNet struct {
	LinearRegression
	Tol, Alpha, L1Ratio float64
	MaxIter             int
	WarmStart, Positive bool
	CDResult            CDResult
}

func NewElasticNet(maxIter int, alpha float64) *ElasticNet {
	regr := &ElasticNet{}
	regr.Tol = 1e-4
	regr.Alpha = alpha
	regr.L1Ratio = .5
	regr.MaxIter = maxIter
	return regr
}

func NewLasso(maxIter int, alpha float64) *ElasticNet {
	regr := NewElasticNet(maxIter, alpha)
	regr.L1Ratio = 1.
	return regr
}

func (regr *ElasticNet) Fit(Xs, Ys [][]float64, random bool) error {
	X, Y := Array2DToDense(Xs), Array2DToDense(Ys)

	NSamples, NFeatures := X.Dims()
	_, NOutputs := Y.Dims()

	l1reg := regr.Alpha * regr.L1Ratio * float64(NSamples)
	l2reg := regr.Alpha * (1. - regr.L1Ratio) * float64(NSamples)
	if !regr.WarmStart {
		regr.Coef = mat.NewDense(NFeatures, NOutputs, nil)
	}
	regr.CDResult = *CoordinateDescent(regr.Coef, l1reg, l2reg, X, Y, regr.MaxIter, regr.Tol, nil, random, regr.Positive)
	return nil
}

func NewLassoJS(this js.Value, args []js.Value) interface{} {
	obj := js.Global().Get("Object").New()
	maxIter, alpha := getargs(args)
	regr := NewLasso(maxIter, alpha)
	setjsmethods(obj, regr)
	return obj
}

func NewElasticNetJS(this js.Value, args []js.Value) interface{} {
	obj := js.Global().Get("Object").New()
	maxIter, alpha := getargs(args)
	regr := NewElasticNet(maxIter, alpha)
	setjsmethods(obj, regr)
	return obj
}

func getargs(args []js.Value) (int, float64) {
	maxIter := 100
	if !args[0].IsUndefined() {
		maxIter = args[0].Int()
	}

	alpha := 1e-4
	if len(args) > 1 && !args[1].IsUndefined() {
		alpha = args[1].Float()
	}

	return maxIter, alpha
}

func setjsmethods(obj js.Value, regr *ElasticNet) {
	obj.Set("fit", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		X := JSFloatArray2D(args[0])
		Y := JSFloatArray2D(args[1])
		random := false
		if len(args) == 3 && !args[2].IsUndefined() {
			random = args[2].Bool()
		}
		return regr.Fit(X, Y, random)
	}))

	obj.Set("predict", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		X := JSFloatArray2D(args[0])
		Y, _ := regr.Predict(X)
		return ToJSArray(Y)
	}))
}
