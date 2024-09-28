//go:build js && wasm
// +build js,wasm

package main

import (
	"fmt"

	"l12.xyz/x/shortcuts/stat/src"

	"syscall/js"
)

func InitStatExports(this js.Value, args []js.Value) interface{} {
	exports := args[0]
	exports.Set("Bhattacharyya", js.FuncOf(src.Bhattacharyya))
	exports.Set("BivariateMoment", js.FuncOf(src.BivariateMoment))
	exports.Set("ChiSquare", js.FuncOf(src.ChiSquare))
	exports.Set("CircularMean", js.FuncOf(src.CircularMean))
	exports.Set("Correlation", js.FuncOf(src.Correlation))
	exports.Set("Covariance", js.FuncOf(src.Covariance))
	exports.Set("CrossEntropy", js.FuncOf(src.CrossEntropy))
	exports.Set("Entropy", js.FuncOf(src.Entropy))
	exports.Set("ExKurtosis", js.FuncOf(src.ExKurtosis))
	exports.Set("GeometricMean", js.FuncOf(src.GeometricMean))
	exports.Set("HarmonicMean", js.FuncOf(src.HarmonicMean))
	exports.Set("Hellinger", js.FuncOf(src.Hellinger))
	exports.Set("Histogram", js.FuncOf(src.Histogram))
	exports.Set("JensenShannon", js.FuncOf(src.JensenShannon))
	exports.Set("Kendall", js.FuncOf(src.Kendall))
	exports.Set("KolmogorovSmirnov", js.FuncOf(src.KolmogorovSmirnov))
	exports.Set("KullbackLeibler", js.FuncOf(src.KullbackLeibler))
	exports.Set("LinearRegression", js.FuncOf(src.LinearRegression))
	exports.Set("Mean", js.FuncOf(src.Mean))
	exports.Set("MeanStdDev", js.FuncOf(src.MeanStdDev))
	exports.Set("MeanVariance", js.FuncOf(src.MeanVariance))
	exports.Set("Mode", js.FuncOf(src.Mode))
	exports.Set("Moment", js.FuncOf(src.Moment))
	exports.Set("MomentAbout", js.FuncOf(src.MomentAbout))
	exports.Set("PopMeanStdDev", js.FuncOf(src.PopMeanStdDev))
	exports.Set("PopMeanVariance", js.FuncOf(src.PopMeanVariance))
	exports.Set("PopStdDev", js.FuncOf(src.PopStdDev))
	exports.Set("PopVariance", js.FuncOf(src.PopVariance))
	exports.Set("Quantile", js.FuncOf(src.Quantile))
	exports.Set("RNoughtSquared", js.FuncOf(src.RNoughtSquared))
	exports.Set("ROC", js.FuncOf(src.ROC))
	exports.Set("RSquared", js.FuncOf(src.RSquared))
	exports.Set("RSquaredFrom", js.FuncOf(src.RSquaredFrom))
	exports.Set("Skew", js.FuncOf(src.Skew))
	exports.Set("SortWeighted", js.FuncOf(src.SortWeighted))
	exports.Set("SortWeightedLabeled", js.FuncOf(src.SortWeightedLabeled))
	exports.Set("StdDev", js.FuncOf(src.StdDev))
	exports.Set("StdErr", js.FuncOf(src.StdErr))
	exports.Set("StdScore", js.FuncOf(src.StdScore))
	exports.Set("TOC", js.FuncOf(src.TOC))
	exports.Set("Variance", js.FuncOf(src.Variance))
	return nil
}

func main() {
	wait := make(chan struct{}, 0)
	js.Global().Set("__InitStatExports", js.FuncOf(InitStatExports))
	fmt.Println("Stats initialized")
	<-wait
}
