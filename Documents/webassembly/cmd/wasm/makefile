
bld: 
	GOOS=js GOARCH=wasm go build -o ../../assets/fileprocessor.wasm
	
run: 
	rm ../../assets/fileprocessor.wasm
	GOOS=js GOARCH=wasm go build -o ../../assets/fileprocessor.wasm
	go run ../server/main.go