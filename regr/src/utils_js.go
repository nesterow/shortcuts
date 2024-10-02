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

func JSFloatArray2D(arg js.Value) [][]float64 {
	arr := make([][]float64, arg.Length())
	for i := 0; i < len(arr); i++ {
		arr[i] = make([]float64, arg.Index(i).Length())
	}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < arg.Index(i).Length(); j++ {
			arr[i][j] = arg.Index(i).Index(j).Float()
		}
	}
	return arr
}

func ToJSArray[T any](arr []T) []interface{} {
	jsArr := make([]interface{}, len(arr))
	for i, v := range arr {
		jsArr[i] = v
	}
	return jsArr
}
