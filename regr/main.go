//go:build js && wasm
// +build js,wasm

package main

import (
	"l12.xyz/x/shortcuts/regr/src"

	"syscall/js"
)

func InitRegrExports(this js.Value, args []js.Value) interface{} {
	exports := js.Global().Get("Object").New()
	exports.Set("Linear", js.FuncOf(src.NewLinearRegressionJS))
	exports.Set("ElasticNet", js.FuncOf(src.NewElasticNetJS))
	exports.Set("Lasso", js.FuncOf(src.NewLassoJS))
	exports.Set("R2Score", js.FuncOf(src.R2ScoreJS))
	exports.Set("Logistic", js.FuncOf(src.NewLogisticRegressionJS))
	return exports
}

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("__InitRegrExports", js.FuncOf(InitRegrExports))
	<-c
}
