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
		values    = args[0]
		bins      = args[1]
		opts      js.Value
		text      = "Histogram"
		lineColor = color.RGBA{R: 255, G: 255, B: 255, A: 255}
		fillColor = color.RGBA{R: 50, G: 50, B: 50, A: 255}
		normColor = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	)

	if len(args) == 3 {
		opts = args[2]
		if !opts.Get("title").IsUndefined() {
			text = opts.Get("title").String()
		}
		if !opts.Get("lineColor").IsUndefined() {
			lineColor = HexToRGBA(opts.Get("lineColor").String())
		}
		if !opts.Get("fillColor").IsUndefined() {
			fillColor = HexToRGBA(opts.Get("fillColor").String())
		}
		if !opts.Get("normColor").IsUndefined() {
			normColor = HexToRGBA(opts.Get("normColor").String())
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
	h.LineStyle.Color = lineColor
	h.FillColor = fillColor
	h.Normalize(1)
	p.Add(h)
	norm := plotter.NewFunction(distuv.UnitNormal.Prob)
	norm.Color = normColor
	norm.Width = vg.Points(2)
	p.Add(norm)
	writer, err := p.WriterTo(6*vg.Inch, 4*vg.Inch, "png")
	if err != nil {
		panic(err)
	}
	b64string := WriterToBase64String(writer)
	return b64string
}
