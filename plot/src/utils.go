//go:build js && wasm
// +build js,wasm

package src

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/color"
	"io"
	"syscall/js"

	"gonum.org/v1/plot/plotter"
)

func WriterToBase64String(writer io.WriterTo) string {
	var buf bytes.Buffer
	writer.WriteTo(&buf)
	return base64.StdEncoding.EncodeToString(buf.Bytes())
}

func HexToRGBA(hex string) color.RGBA {
	var c color.RGBA
	switch len(hex) {
	case 4:
		fmt.Sscanf(hex, "#%1x%1x%1x%1x", &c.R, &c.G, &c.B, &c.A)
		c.R *= 17
		c.G *= 17
		c.B *= 17
		c.A = 255
	case 7:
		fmt.Sscanf(hex, "#%02x%02x%02x", &c.R, &c.G, &c.B)
		c.A = 255
	case 9:
		fmt.Sscanf(hex, "#%02x%02x%02x%02x", &c.R, &c.G, &c.B, &c.A)
	default:
		fmt.Println("Invalid hex color:", hex)
	}
	return c
}

func XYFromJSObject(object js.Value) plotter.XYs {
	var (
		x = object.Get("x")
		y = object.Get("y")
	)
	if x.IsUndefined() || y.IsUndefined() {
		return nil
	}
	return XYFromJSValues(x, y)
}

func XYFromJSValues(x, y js.Value) plotter.XYs {
	xy := make(plotter.XYs, x.Length())
	for i := range xy {
		xy[i].X = x.Index(i).Float()
		xy[i].Y = y.Index(i).Float()
	}
	return xy
}
