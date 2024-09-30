dev:
	GOOS=js GOARCH=wasm tinygo build -o stat/mod.wasm -no-debug ./stat/main.go
	GOOS=js GOARCH=wasm tinygo build -o regr/mod.wasm  ./regr/main.go
	GOOS=js GOARCH=wasm go build -o plot/mod.wasm ./plot/main.go

prod:
	GOOS=js GOARCH=wasm tinygo build -o stat/mod.wasm -no-debug ./stat/main.go
	GOOS=js GOARCH=wasm tinygo build -o plot/mod.wasm -no-debug ./plot/main.go