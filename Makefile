wasm:
	tinygo build -o ./wasm/main.wasm -scheduler=none -target=wasi ./wasm/wasm.go

envoy:
	envoy -c envoy.yaml -l debug
