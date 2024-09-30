//go:build js && wasm
// +build js,wasm

package src

import (
	"image/color"
	"syscall/js"

	"github.com/gonum/stat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

func Plot(this js.Value, args []js.Value) interface{} {
	plt := plot.New()
	opts := args[0]
	var (
		title  = opts.Get("title").String()
		XLabel = opts.Get("XLabel").String()
		YLabel = opts.Get("YLabel").String()
		width  = 7.5
		height = 4.25
	)
	if !opts.Get("width").IsUndefined() {
		width = opts.Get("width").Float()
	}
	if !opts.Get("height").IsUndefined() {
		height = opts.Get("height").Float()
	}
	if opts.Get("title").IsUndefined() {
		title = "Plot"
	}
	if opts.Get("XLabel").IsUndefined() {
		XLabel = "X"
	}
	if opts.Get("YLabel").IsUndefined() {
		YLabel = "Y"
	}

	plt.Title.Text = title
	plt.X.Label.Text = XLabel
	plt.Y.Label.Text = YLabel
	plt.Add(plotter.NewGrid())

	lines := make([]plot.Plotter, 0)
	for i, arg := range args {
		if i == 0 {
			continue
		}
		PlotterLinesFromJSObject(arg, plt, &lines)
	}
	plt.Add(lines...)
	plt.Legend.Top = true

	writer, err := plt.WriterTo(vg.Length(width)*vg.Inch, vg.Length(height)*vg.Inch, "png")
	if err != nil {
		panic(err)
	}
	b64string := WriterToBase64String(writer)
	return b64string
}

func PlotterLinesFromJSObject(object js.Value, plt *plot.Plot, plotters *[]plot.Plotter) {
	var (
		typ         = object.Get("type").String()
		data        = object.Get("data")
		legend      = ""
		glyphColor  = object.Get("glyphColor")
		glyphRadius = object.Get("glypRadius")
		glyphShape  = object.Get("glyphShape")
		lineWidth   = object.Get("lineWidth")
		lineColor   = object.Get("lineColor")
		lineDashes  = object.Get("lineDashes")
	)

	shapeGlyph := func(g *draw.GlyphStyle) {
		if !glyphShape.IsUndefined() {
			shape := glyphShape.String()
			switch shape {
			case "cross":
				g.Shape = draw.CrossGlyph{}
			case "plus":
				g.Shape = draw.PlusGlyph{}
			case "ring":
				g.Shape = draw.RingGlyph{}
			case "circle":
				g.Shape = draw.CircleGlyph{}
			case "square":
				g.Shape = draw.SquareGlyph{}
			case "triangle":
				g.Shape = draw.TriangleGlyph{}
			case "pyramid":
				g.Shape = draw.PyramidGlyph{}
			}
		}
	}

	radiusGlyph := func(r *vg.Length) {
		if !glyphRadius.IsUndefined() {
			*r = vg.Points(glyphRadius.Float())
		}
	}

	dashLine := func(l *[]vg.Length) {
		if !lineDashes.IsUndefined() {
			w := lineDashes.Index(0).Float()
			h := lineDashes.Index(1).Float()
			*l = []vg.Length{vg.Points(w), vg.Points(h)}
		}
	}

	applyColor := func(p *color.Color, hex js.Value) {
		if !hex.IsUndefined() {
			*p = HexToRGBA(hex.String())
		}
	}

	if !object.Get("legend").IsUndefined() {
		legend = object.Get("legend").String()
	}
	if data.IsUndefined() {
		panic("data is undefined")
	}

	XYs := XYFromJSObject(data)
	if XYs == nil {
		XYs = XYFromJSValues(data.Index(0), data.Index(1))
	}

	switch typ {
	case "scatter":
		s, err := plotter.NewScatter(XYs)
		if err != nil {
			panic(err)
		}
		applyColor(&s.GlyphStyle.Color, glyphColor)
		shapeGlyph(&s.GlyphStyle)
		radiusGlyph(&s.GlyphStyle.Radius)
		*plotters = append(*plotters, s)
		if legend != "" {
			plt.Legend.Add(legend, s)
		}
	case "line":
		l, err := plotter.NewLine(XYs)
		if err != nil {
			panic(err)
		}
		applyColor(&l.LineStyle.Color, lineColor)
		dashLine(&l.LineStyle.Dashes)
		if !lineWidth.IsUndefined() {
			l.LineStyle.Width = vg.Points(lineWidth.Float())
		}
		*plotters = append(*plotters, l)
		if legend != "" {
			plt.Legend.Add(legend, l)
		}
	case "linePoints":
		l, lp, err := plotter.NewLinePoints(XYs)
		if err != nil {
			panic(err)
		}
		if !lineWidth.IsUndefined() {
			l.LineStyle.Width = vg.Points(lineWidth.Float())
		}
		applyColor(&l.LineStyle.Color, lineColor)
		dashLine(&l.LineStyle.Dashes)
		shapeGlyph(&lp.GlyphStyle)
		applyColor(&lp.GlyphStyle.Color, glyphColor)
		radiusGlyph(&lp.GlyphStyle.Radius)
		*plotters = append(*plotters, l, lp)
		if legend != "" {
			plt.Legend.Add(legend, l, lp)
		}
	case "trend":
		X := make([]float64, len(XYs))
		Y := make([]float64, len(XYs))
		min := XYs[0].Y
		max := XYs[0].Y
		for i, xy := range XYs {
			X[i] = xy.X
			Y[i] = xy.Y
			if xy.Y < min {
				min = xy.Y
			}
			if xy.Y > max {
				max = xy.Y
			}
		}
		a, b := stat.LinearRegression(X, Y, nil, false)
		l, err := plotter.NewLine(plotter.XYs{
			{X: X[0], Y: a + b*X[0]},
			{X: X[len(X)-1], Y: a + b*X[len(X)-1]},
		})
		if err != nil {
			panic(err)
		}
		applyColor(&l.LineStyle.Color, lineColor)
		dashLine(&l.LineStyle.Dashes)
		if !lineWidth.IsUndefined() {
			l.LineStyle.Width = vg.Points(lineWidth.Float())
		}
		*plotters = append(*plotters, l)
		if legend != "" {
			plt.Legend.Add(legend, l)
		}
	}
}
