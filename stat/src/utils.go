//go:build js && wasm
// +build js,wasm

package src

import (
	"syscall/js"
)

func JSFloatArray(arg js.Value) []float64 {
	arr := make([]float64, arg.Length())
	for i := 0; i < len(arr); i++ {
		arr[i] = arg.Index(i).Float()
	}
	return arr
}

func JSBoolArray(arg js.Value) []bool {
	arr := make([]bool, arg.Length())
	for i := 0; i < len(arr); i++ {
		arr[i] = arg.Index(i).Bool()
	}
	return arr
}
