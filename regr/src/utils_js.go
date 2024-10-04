//go:build js && wasm
// +build js,wasm

package src

import (
	"syscall/js"

	"gonum.org/v1/gonum/mat"
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

func MatrixToJSArray(m mat.Matrix) []interface{} {
	r, c := m.Dims()
	data := make([]float64, r*c)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			data[i*c+j] = m.At(i, j)
		}
	}
	return ToJSArray(data)
}

func MatrixToJSArray2D(m mat.Matrix) []interface{} {
	r, c := m.Dims()
	data := make([][]interface{}, r)
	for i := 0; i < r; i++ {
		data[i] = make([]interface{}, c)
		for j := 0; j < c; j++ {
			data[i][j] = m.At(i, j)
		}
	}
	return ToJSArray(data)
}
