//go:build js && wasm
// +build js,wasm

package src

import "syscall/js"

// ref: mat.Dense
// fit: solve least squares
// predict: predict y from x
// save: save model
// load: load model
// note: separate wasm/js glue
func ABCD(this js.Value, args []js.Value) interface{} {
	return nil
}
