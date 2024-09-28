wasm:
	GOOS=js GOARCH=wasm tinygo build -o stat/mod.wasm ./stat/main.go