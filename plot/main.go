//go:build js && wasm
// +build js,wasm

package main

import (
	"l12.xyz/x/shortcuts/plot/src"

	"syscall/js"
)

func InitPlotExports(this js.Value, args []js.Value) interface{} {
	exports := args[0]
	exports.Set("Hist", js.FuncOf(src.HistPlot))
	return nil
}

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("__InitPlotExports", js.FuncOf(InitPlotExports))
	<-c
}
