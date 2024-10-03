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
}

func (regr *MCLogisticRegression) Fit(x, y mat.Matrix) {
	_, regr.cols = y.Dims()
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
	rows, _ := x.Dims()
	probs := mat.NewDense(rows, regr.cols, nil)
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

func (regr *MCLogisticRegression) Loss(yTrue, yPred mat.Matrix) float64 {
	loss := 0.
	for j := 0; j < regr.cols; j++ {
		yj := mat.Col(nil, j, yTrue)
		Y := mat.NewDense(len(yj), 1, yj)
		ypj := mat.Col(nil, j, yPred)
		YP := mat.NewDense(len(ypj), 1, ypj)
		loss += regr.Models[j].Loss(Y, YP)
	}
	return loss / float64(regr.cols)
}
