//go:build js && wasm
// +build js,wasm

package main

import (
	"l12.xyz/x/shortcuts/regr/src"

	"syscall/js"
)

func InitRegrExports(this js.Value, args []js.Value) interface{} {
	exports := args[0]
	exports.Set("Linear", js.FuncOf(src.NewLinearRegressionJS))
	exports.Set("ElasticNet", js.FuncOf(src.NewElasticNetJS))
	exports.Set("Lasso", js.FuncOf(src.NewLassoJS))
	return nil
}

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("__InitRegrExports", js.FuncOf(InitRegrExports))
	<-c
}
