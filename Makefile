dev:
	GOOS=js GOARCH=wasm tinygo build -o stat/mod.wasm ./stat/main.go

prod:
	GOOS=js GOARCH=wasm tinygo build -o stat/mod.wasm -no-debug ./stat/main.go