package src

import (
	"math"

	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/mat"
)

type CDResult struct {
	Gap   float64
	Eps   float64
	NIter int
}

// Author: Pascal Masschelier, https://github.com/pa-m
// v https://github.com/pa-m/sklearn/blob/master/linear_model/cdfast.go
//
// Coordinate descent algorithm for Elastic-Net
// v https://github.com/scikit-learn/scikit-learn/blob/a24c8b464d094d2c468a16ea9f8bf8d42d949f84/sklearn/linear_model/cd_fast.pyx
func CoordinateDescent(w *mat.Dense, l1reg, l2reg float64, X *mat.Dense, Y *mat.Dense, maxIter int, tol float64, rng *rand.Rand, random, positive bool) *CDResult {
	/*
	   coordinate descent algorithm
	       for Elastic-Net regression
	       We minimize
	       (1/2) * norm(y - X w, 2)^2 + alpha norm(w, 1) + (beta/2) norm(w, 2)^2
	*/
	gap := tol + 1.
	dwtol := tol

	NSamples, NFeatures := X.Dims()
	_ = NSamples
	_, NTasks := Y.Dims()

	var dwii, wmax, dwmax, dualNormXtA, cons, XtAaxis1norm, wiiabsmax, l21norm float64
	var ii, fIter, nIter int
	tmp := mat.NewVecDense(NTasks, nil)
	wii := mat.NewVecDense(NTasks, nil)

	R, Y2, XtA, RY, xw := &mat.Dense{}, &mat.Dense{}, &mat.Dense{}, &mat.Dense{}, &mat.Dense{}
	RNorm, wNorm, ANorm, nn, tmpfactor := 0., 0., 0., 0., 0.
	XtA = mat.NewDense(NFeatures, NTasks, nil)

	normColsX := make([]float64, NFeatures)
	Xmat := X.RawMatrix()
	for jX := 0; jX < Xmat.Rows*Xmat.Stride; jX = jX + Xmat.Stride {
		for i, v := range Xmat.Data[jX : jX+Xmat.Cols] {
			normColsX[i] += v * v
		}
	}
	R.Mul(X, w)
	R.Sub(Y, R)
	Y2.MulElem(Y, Y)
	tol *= mat.Sum(Y2)

	for nIter = 0; nIter < maxIter; nIter++ {
		wmax, dwmax = 0., 0.
		for fIter = 0; fIter < NFeatures; fIter++ {
			if random {
				if rng != nil {
					ii = rng.Intn(NFeatures)
				} else {
					ii = rand.Intn(NFeatures)
				}
			} else {
				ii = fIter
			}
			if normColsX[ii] == 0. {
				continue
			}
			wii.CopyVec(w.RowView(ii)) // store previous value
			if mat.Norm(wii, 2) != 0. {
				xw.Mul(X.ColView(ii), wii.T())
				R.Add(R, xw)
			}
			tmp.MulVec(R.T(), X.ColView(ii))
			nn = mat.Norm(tmp, 2)
			if l1reg < nn {
				tmpfactor = math.Max(1.-l1reg/nn, 0) / (normColsX[ii] + l2reg)

			} else {
				tmpfactor = 0.
			}
			w.RowView(ii).(*mat.VecDense).ScaleVec(tmpfactor, tmp)
			if mat.Norm(w.RowView(ii), 2) != 0. {
				xw.Mul(X.ColView(ii), w.RowView(ii).T())
				R.Sub(R, xw)
			}
			dwii = 0.
			for o := 0; o < NTasks; o++ {
				v := math.Abs(w.At(ii, o) - wii.AtVec(o))
				if v > dwii {
					dwii = v
				}
			}

			if dwii > dwmax {
				dwmax = dwii
			}
			wiiabsmax = 0.
			wmat := w.RawMatrix()
			for jw := 0; jw < wmat.Rows*wmat.Stride; jw = jw + wmat.Stride {
				for _, v := range wmat.Data[jw : jw+wmat.Cols] {
					v = math.Abs(v)
					if v > wiiabsmax {
						wiiabsmax = v
					}
				}
			}
			if wiiabsmax > wmax {
				wmax = wiiabsmax
			}

		}

		if wmax == 0. || dwmax/wmax < dwtol || nIter == maxIter-1 {
			XtA.Mul(X.T(), R)
			XtAmat := XtA.RawMatrix()
			Wmat := w.RawMatrix()
			for jXtA, jW := 0, 0; jXtA < XtAmat.Rows*XtAmat.Stride; jXtA, jW = jXtA+XtAmat.Stride, jW+Wmat.Stride {
				for i := range XtAmat.Data[jXtA : jXtA+XtAmat.Cols] {
					XtAmat.Data[jXtA+i] -= l2reg * Wmat.Data[jW+i]
				}
			}
			dualNormXtA = 0.
			for ii := 0; ii < NFeatures; ii++ {
				XtAaxis1norm = mat.Norm(XtA.RowView(ii), 2)
				if XtAaxis1norm > dualNormXtA {
					dualNormXtA = XtAaxis1norm
				}
			}
			RNorm = mat.Norm(R, 2)
			wNorm = mat.Norm(w, 2)
			if dualNormXtA > l1reg {
				cons = l1reg / dualNormXtA
				ANorm = RNorm * cons
				gap = .5 * (RNorm*RNorm + ANorm*ANorm)
			} else {
				cons = 1.
				gap = RNorm * RNorm
			}
			RY.MulElem(R, Y)
			l21norm = 0.
			for ii = 0; ii < NFeatures; ii++ {
				l21norm += mat.Norm(w.RowView(ii), 2)
			}
			gap += l1reg*l21norm - cons*mat.Sum(RY) + .5*l2reg*(1.+cons*cons)*(wNorm*wNorm)
			if gap < tol {
				// return if we have reached the desired tolerance
				break
			}
		}

	}
	return &CDResult{Gap: gap, Eps: tol, NIter: nIter + 1}
}
