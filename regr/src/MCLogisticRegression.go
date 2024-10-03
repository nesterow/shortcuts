package src

import (
	"sync"

	"gonum.org/v1/gonum/mat"
)

type MCLogisticRegression struct {
	Epochs       int
	LearningRate float64
	Models       []*LogisticRegression
	cols         int
	rows         int
}

func (regr *MCLogisticRegression) Fit(x, y mat.Matrix) {
	regr.rows, regr.cols = y.Dims()
	regr.Models = make([]*LogisticRegression, regr.cols)
	for j := 0; j < regr.cols; j++ {
		regr.Models[j] = &LogisticRegression{
			Epochs:       regr.Epochs,
			LearningRate: regr.LearningRate,
		}
		yj := mat.Col(nil, j, y)
		Y := mat.NewDense(len(yj), 1, yj)
		regr.Models[j].Fit(x, Y, nil)
	}
}

func (regr *MCLogisticRegression) Predict(x mat.Matrix) mat.Matrix {
	probs := mat.NewDense(regr.rows, regr.cols, nil)
	wg := sync.WaitGroup{}
	wg.Add(regr.cols)
	for j := 0; j < regr.cols; j++ {
		go func(j int) {
			pred := regr.Models[j].Predict(x).(*mat.Dense)
			probs.SetCol(j, mat.Col(nil, 0, pred))
			wg.Done()
		}(j)
	}
	wg.Wait()
	return probs
}
