//go:build js && wasm
// +build js,wasm

package src

import (
	"syscall/js"

	"gonum.org/v1/gonum/stat"
)

func Bhattacharyya(this js.Value, args []js.Value) interface{} {
	var (
		p = JSFloatArray(args[0])
		q = JSFloatArray(args[1])
	)
	return stat.Bhattacharyya(p, q)
}

func BivariateMoment(this js.Value, args []js.Value) interface{} {
	var (
		r       = args[0].Float()
		s       = args[1].Float()
		xs      = JSFloatArray(args[2])
		ys      = JSFloatArray(args[3])
		weights = JSFloatArray(args[4])
	)
	if len(weights) == 0 {
		weights = nil
	}
	return stat.BivariateMoment(r, s, xs, ys, weights)
}

func CDF(this js.Value, args []js.Value) interface{} {
	var (
		q       = args[0].Float()
		xs      = JSFloatArray(args[2])
		weights = JSFloatArray(args[3])
	)
	return stat.CDF(q, 1, xs, weights)
}

func ChiSquare(this js.Value, args []js.Value) interface{} {
	var (
		p = JSFloatArray(args[0])
		q = JSFloatArray(args[1])
	)
	return stat.ChiSquare(p, q)
}

func CircularMean(this js.Value, args []js.Value) interface{} {
	var (
		xs      = JSFloatArray(args[0])
		weights = JSFloatArray(args[1])
	)
	if len(weights) == 0 {
		weights = nil
	}
	return stat.CircularMean(xs, weights)
}

func Correlation(this js.Value, args []js.Value) interface{} {
	var (
		xs      = JSFloatArray(args[0])
		ys      = JSFloatArray(args[1])
		weights = JSFloatArray(args[2])
	)
	if len(weights) == 0 {
		weights = nil
	}
	return stat.Correlation(xs, ys, weights)
}

func CorrelationMatrix(this js.Value, args []js.Value) interface{} {
	panic("not implemented")
}

func Covariance(this js.Value, args []js.Value) interface{} {
	var (
		xs      = JSFloatArray(args[0])
		ys      = JSFloatArray(args[1])
		weights = JSFloatArray(args[2])
	)
	if len(weights) == 0 {
		weights = nil
	}
	return stat.Covariance(xs, ys, weights)
}

func CovarianceMatrix(this js.Value, args []js.Value) interface{} {
	panic("not implemented")
}

func CrossEntropy(this js.Value, args []js.Value) interface{} {
	var (
		p = JSFloatArray(args[0])
		q = JSFloatArray(args[1])
	)
	return stat.CrossEntropy(p, q)
}

func Entropy(this js.Value, args []js.Value) interface{} {
	var (
		p = JSFloatArray(args[0])
	)
	return stat.Entropy(p)
}

func ExKurtosis(this js.Value, args []js.Value) interface{} {
	var (
		xs      = JSFloatArray(args[0])
		weights = JSFloatArray(args[1])
	)
	if len(weights) == 0 {
		weights = nil
	}
	return stat.ExKurtosis(xs, weights)
}

func GeometricMean(this js.Value, args []js.Value) interface{} {
	var (
		xs      = JSFloatArray(args[0])
		weights = JSFloatArray(args[1])
	)
	if len(weights) == 0 {
		weights = nil
	}
	return stat.GeometricMean(xs, weights)
}

func HarmonicMean(this js.Value, args []js.Value) interface{} {
	var (
		xs      = JSFloatArray(args[0])
		weights = JSFloatArray(args[1])
	)
	if len(weights) == 0 {
		weights = nil
	}
	return stat.HarmonicMean(xs, weights)
}

func Hellinger(this js.Value, args []js.Value) interface{} {
	var (
		p = JSFloatArray(args[0])
		q = JSFloatArray(args[1])
	)
	return stat.Hellinger(p, q)
}

func Histogram(this js.Value, args []js.Value) interface{} {
	var (
		counts  = JSFloatArray(args[0])
		divs    = JSFloatArray(args[1])
		xs      = JSFloatArray(args[2])
		weights = JSFloatArray(args[3])
	)
	if len(weights) == 0 {
		weights = nil
	}
	return stat.Histogram(counts, divs, xs, weights)
}

func JensenShannon(this js.Value, args []js.Value) interface{} {
	var (
		p = JSFloatArray(args[0])
		q = JSFloatArray(args[1])
	)
	return stat.JensenShannon(p, q)
}

func Kendall(this js.Value, args []js.Value) interface{} {
	var (
		xs      = JSFloatArray(args[0])
		ys      = JSFloatArray(args[1])
		weights = JSFloatArray(args[2])
	)
	if len(weights) == 0 {
		weights = nil
	}
	return stat.Kendall(xs, ys, weights)
}

func KolmogorovSmirnov(this js.Value, args []js.Value) interface{} {
	var (
		xs       = JSFloatArray(args[0])
		xWeights = JSFloatArray(args[1])
		ys       = JSFloatArray(args[2])
		yWeights = JSFloatArray(args[3])
	)
	return stat.KolmogorovSmirnov(xs, xWeights, ys, yWeights)
}

func KullbackLeibler(this js.Value, args []js.Value) interface{} {
	var (
		p = JSFloatArray(args[0])
		q = JSFloatArray(args[1])
	)
	return stat.KullbackLeibler(p, q)
}

func LinearRegression(this js.Value, args []js.Value) interface{} {
	var (
		xs      = JSFloatArray(args[0])
		ys      = JSFloatArray(args[1])
		weights = JSFloatArray(args[2])
		origin  = args[3].Bool()
	)
	if len(weights) == 0 {
		weights = nil
	}
	alpha, beta := stat.LinearRegression(xs, ys, weights, origin)
	return map[string]interface{}{
		"alpha": alpha,
		"beta":  beta,
	}
}

func Mahalanobis(this js.Value, args []js.Value) interface{} {
	panic("not implemented")
}

func Mean(this js.Value, args []js.Value) interface{} {
	var (
		xs      = JSFloatArray(args[0])
		weights = JSFloatArray(args[1])
	)
	if len(weights) == 0 {
		weights = nil
	}
	return stat.Mean(xs, weights)
}

func MeanStdDev(this js.Value, args []js.Value) interface{} {
	var (
		xs      = JSFloatArray(args[0])
		weights = JSFloatArray(args[1])
	)
	if len(weights) == 0 {
		weights = nil
	}
	mean, stdDev := stat.MeanStdDev(xs, weights)
	return map[string]interface{}{
		"mean":   mean,
		"stdDev": stdDev,
	}
}

func MeanVariance(this js.Value, args []js.Value) interface{} {
	var (
		xs      = JSFloatArray(args[0])
		weights = JSFloatArray(args[1])
	)
	if len(weights) == 0 {
		weights = nil
	}
	mean, variance := stat.MeanVariance(xs, weights)
	return map[string]interface{}{
		"mean":     mean,
		"variance": variance,
	}
}

func Mode(this js.Value, args []js.Value) interface{} {
	var (
		xs      = JSFloatArray(args[0])
		weights = JSFloatArray(args[1])
	)
	if len(weights) == 0 {
		weights = nil
	}
	value, count := stat.Mode(xs, weights)
	return map[string]interface{}{
		"value": value,
		"count": count,
	}
}

func Moment(this js.Value, args []js.Value) interface{} {
	var (
		r       = args[0].Float()
		xs      = JSFloatArray(args[1])
		weights = JSFloatArray(args[2])
	)
	if len(weights) == 0 {
		weights = nil
	}
	return stat.Moment(r, xs, weights)
}

func MomentAbout(this js.Value, args []js.Value) interface{} {
	var (
		r       = args[0].Float()
		xs      = JSFloatArray(args[1])
		mean    = args[2].Float()
		weights = JSFloatArray(args[3])
	)
	if len(weights) == 0 {
		weights = nil
	}
	return stat.MomentAbout(r, xs, mean, weights)
}

func PopMeanStdDev(this js.Value, args []js.Value) interface{} {
	var (
		xs      = JSFloatArray(args[0])
		weights = JSFloatArray(args[1])
	)
	if len(weights) == 0 {
		weights = nil
	}
	mean, stdDev := stat.PopMeanStdDev(xs, weights)
	return map[string]interface{}{
		"mean":   mean,
		"stdDev": stdDev,
	}
}

func PopMeanVariance(this js.Value, args []js.Value) interface{} {
	var (
		xs      = JSFloatArray(args[0])
		weights = JSFloatArray(args[1])
	)
	if len(weights) == 0 {
		weights = nil
	}
	mean, variance := stat.PopMeanVariance(xs, weights)
	return map[string]interface{}{
		"mean":     mean,
		"variance": variance,
	}
}

func PopStdDev(this js.Value, args []js.Value) interface{} {
	var (
		xs      = JSFloatArray(args[0])
		weights = JSFloatArray(args[1])
	)
	if len(weights) == 0 {
		weights = nil
	}
	return stat.PopStdDev(xs, weights)
}

func PopVariance(this js.Value, args []js.Value) interface{} {
	var (
		xs      = JSFloatArray(args[0])
		weights = JSFloatArray(args[1])
	)
	if len(weights) == 0 {
		weights = nil
	}
	return stat.PopVariance(xs, weights)
}

func Quantile(this js.Value, args []js.Value) interface{} {
	var (
		q       = args[0].Float()
		typ     = args[1].String()
		xs      = JSFloatArray(args[2])
		weights = JSFloatArray(args[3])
	)
	if typ == "empirical" {
		return stat.Quantile(q, 1, xs, weights)
	} else if typ == "linear" {
		return stat.Quantile(q, 4, xs, weights)
	} else {
		return stat.Quantile(q, 2, xs, weights)
	}
}

func RNoughtSquared(this js.Value, args []js.Value) interface{} {
	var (
		xs      = JSFloatArray(args[0])
		ys      = JSFloatArray(args[1])
		weights = JSFloatArray(args[2])
		beta    = args[3].Float()
	)
	if len(weights) == 0 {
		weights = nil
	}
	return stat.RNoughtSquared(xs, ys, weights, beta)
}

func ROC(this js.Value, args []js.Value) interface{} {
	var (
		cutoffs = JSFloatArray(args[0])
		y       = JSFloatArray(args[1])
		classes = JSBoolArray(args[2])
		weights = JSFloatArray(args[3])
	)
	tpr, fpr, tresh := stat.ROC(cutoffs, y, classes, weights)
	return map[string]interface{}{
		"tpr":   tpr,
		"fpr":   fpr,
		"tresh": tresh,
	}
}

func RSquared(this js.Value, args []js.Value) interface{} {
	var (
		xs      = JSFloatArray(args[0])
		ys      = JSFloatArray(args[1])
		weights = JSFloatArray(args[2])
		alpha   = args[3].Float()
		beta    = args[4].Float()
	)
	if len(weights) == 0 {
		weights = nil
	}
	return stat.RSquared(xs, ys, weights, alpha, beta)
}

func RSquaredFrom(this js.Value, args []js.Value) interface{} {
	var (
		xs      = JSFloatArray(args[0])
		ys      = JSFloatArray(args[1])
		weights = JSFloatArray(args[2])
	)
	if len(weights) == 0 {
		weights = nil
	}
	return stat.RSquaredFrom(xs, ys, weights)
}

func Skew(this js.Value, args []js.Value) interface{} {
	var (
		xs      = JSFloatArray(args[0])
		weights = JSFloatArray(args[1])
	)
	if len(weights) == 0 {
		weights = nil
	}
	return stat.Skew(xs, weights)
}

func SortWeighted(this js.Value, args []js.Value) interface{} {
	var (
		xs      = JSFloatArray(args[0])
		weights = JSFloatArray(args[1])
	)
	if len(weights) == 0 {
		weights = nil
	}
	stat.SortWeighted(xs, weights)
	return xs
}

func SortWeightedLabeled(this js.Value, args []js.Value) interface{} {
	var (
		xs      = JSFloatArray(args[0])
		labels  = JSBoolArray(args[1])
		weights = JSFloatArray(args[2])
	)
	if len(weights) == 0 {
		weights = nil
	}
	stat.SortWeightedLabeled(xs, labels, weights)
	return map[string]interface{}{
		"xs":     xs,
		"labels": labels,
	}
}

func StdDev(this js.Value, args []js.Value) interface{} {
	var (
		xs      = JSFloatArray(args[0])
		weights = JSFloatArray(args[1])
	)
	if len(weights) == 0 {
		weights = nil
	}
	return stat.StdDev(xs, weights)
}

func StdErr(this js.Value, args []js.Value) interface{} {
	var (
		std  = args[0].Float()
		size = args[1].Float()
	)
	return stat.StdErr(std, size)
}

func StdScore(this js.Value, args []js.Value) interface{} {
	var (
		x    = args[0].Float()
		mean = args[1].Float()
		std  = args[2].Float()
	)
	return stat.StdScore(x, mean, std)
}

func TOC(this js.Value, args []js.Value) interface{} {
	var (
		classes = JSBoolArray(args[0])
		weights = JSFloatArray(args[1])
	)
	min, ntp, max := stat.TOC(classes, weights)
	return map[string]interface{}{
		"min": min,
		"ntp": ntp,
		"max": max,
	}
}

func Variance(this js.Value, args []js.Value) interface{} {
	var (
		xs      = JSFloatArray(args[0])
		weights = JSFloatArray(args[1])
	)
	if len(weights) == 0 {
		weights = nil
	}
	return stat.Variance(xs, weights)
}
