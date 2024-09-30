//go:build js && wasm
// +build js,wasm

package src

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/color"
	"io"
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
