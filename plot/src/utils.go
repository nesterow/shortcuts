//go:build js && wasm
// +build js,wasm

package src

import (
	"bytes"
	"encoding/base64"
	"io"
)

func WriterToBase64String(writer io.WriterTo) string {
	var buf bytes.Buffer
	writer.WriteTo(&buf)
	return base64.StdEncoding.EncodeToString(buf.Bytes())
}
