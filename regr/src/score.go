package src

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

type MatConst struct {
	Rows, Columns int
	Value         float64
}

func (m MatConst) Dims() (int, int) { return m.Rows, m.Columns }

func (m MatConst) At(i, j int) float64 { return m.Value }

func (m MatConst) T() mat.Matrix { return MatConst{Rows: m.Columns, Columns: m.Rows, Value: m.Value} }

func R2Score(yTrue, yPred mat.Matrix, sampleWeight *mat.Dense, multioutput string) *mat.Dense {
	nSamples, nOutputs := yTrue.Dims()
	if sampleWeight == nil {
		sampleWeight = mat.DenseCopyOf(MatConst{Rows: nSamples, Columns: 1, Value: 1.})
	}
	numerator := mat.NewDense(1, nOutputs, nil)
	diff := mat.NewDense(nSamples, nOutputs, nil)
	diff.Sub(yPred, yTrue)
	diff2 := mat.NewDense(nSamples, nOutputs, nil)
	diff2.MulElem(diff, diff)
	numerator.Mul(sampleWeight.T(), diff2)

	sampleWeightSum := mat.Sum(sampleWeight)

	yTrueAvg := mat.NewDense(1, nOutputs, nil)
	yTrueAvg.Mul(sampleWeight.T(), yTrue)
	yTrueAvg.Scale(1./sampleWeightSum, yTrueAvg)

	diff2.Apply(func(i int, j int, _ float64) float64 {
		v := yTrue.At(i, j) - yTrueAvg.At(0, j)
		return v * v
	}, diff2)
	denominator := mat.NewDense(1, nOutputs, nil)
	denominator.Mul(sampleWeight.T(), diff2)

	r2score := mat.NewDense(1, nOutputs, nil)
	r2score.Apply(func(i int, j int, v float64) float64 {
		d := math.Max(denominator.At(i, j), 1e-20)
		return 1. - numerator.At(i, j)/d
	}, r2score)
	switch multioutput {
	case "raw_values":
		return r2score
	case "variance_weighted":
		r2 := mat.NewDense(1, 1, nil)
		r2.Mul(denominator, r2score.T())
		sumden := mat.Sum(denominator)
		r2.Scale(1./sumden, r2)
		return r2
	default: // "uniform_average":
		return mat.NewDense(1, 1, []float64{mat.Sum(r2score) / float64(nOutputs)})
	}

}

func MeanSquaredError(yTrue, yPred mat.Matrix, sampleWeight *mat.Dense, multioutput string) *mat.Dense {
	nSamples, nOutputs := yTrue.Dims()
	tmp := mat.NewDense(1, nOutputs, nil)
	tmp.Apply(func(_ int, j int, v float64) float64 {
		N, D := 0., 0.
		for i := 0; i < nSamples; i++ {
			ydiff := yPred.At(i, j) - yTrue.At(i, j)
			w := 1.
			if sampleWeight != nil {
				w = sampleWeight.At(0, j)
			}
			N += w * (ydiff * ydiff)
			D += w
		}
		return N / D
	}, tmp)
	switch multioutput {
	case "raw_values":
		return tmp
	default: // "uniform_average":
		return mat.NewDense(1, 1, []float64{mat.Sum(tmp) / float64(nOutputs)})
	}
}
