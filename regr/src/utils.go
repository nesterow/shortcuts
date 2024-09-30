//go:build js && wasm
// +build js,wasm

package src

import (
	"syscall/js"

	"gonum.org/v1/gonum/mat"
)

func Array2DToDense(X [][]float64) *mat.Dense {
	dense := mat.NewDense(len(X), len(X[0]), nil)
	for i, row := range X {
		dense.SetRow(i, row)
	}
	return dense
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
