wasm:
	tinygo build -o ./wasm/main.wasm -scheduler=none -target=wasi ./wasm/wasm.go
.PHONY: wasm

envoy:
	envoy -c ./envoy.yml -l debug
