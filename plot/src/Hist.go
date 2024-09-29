//go:build js && wasm
// +build js,wasm

package src

import (
	"image/color"

	"syscall/js"

	"github.com/gonum/stat/distuv"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func HistPlot(this js.Value, args []js.Value) interface{} {
	var (
		values = args[0]
		bins   = args[1]
		opts   js.Value
		text   = "Histogram"
	)

	if len(args) == 3 {
		opts = args[2]
		if !opts.Get("title").IsUndefined() {
			text = opts.Get("title").String()
		}
	}

	v := make(plotter.Values, values.Length())
	for i := range v {
		v[i] = values.Index(i).Float()
	}
	p := plot.New()
	p.Title.Text = text
	h, err := plotter.NewHist(v, bins.Int())
	if err != nil {
		panic(err)
	}
	h.Normalize(1)
	p.Add(h)
	norm := plotter.NewFunction(distuv.UnitNormal.Prob)
	norm.Color = color.RGBA{R: 255, A: 255}
	norm.Width = vg.Points(2)
	p.Add(norm)
	writer, err := p.WriterTo(6*vg.Inch, 4*vg.Inch, "png")
	if err != nil {
		panic(err)
	}
	b64string := WriterToBase64String(writer)
	return b64string
}
