wasm:
	GOOS=js GOARCH=wasm tinygo build -o stats/mod.wasm ./stats/main.go