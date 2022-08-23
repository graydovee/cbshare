.PHONY: generate
generate:
	protoc --go_out=./pkg --go_opt=paths=source_relative  --go-grpc_out=./pkg --go-grpc_opt=paths=source_relative proto/*.proto

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: build
build: generate vet fmt
	GOOS=darwin GOARCH=amd64 go build -o bin/amd64/darwin/cbshare
	GOOS=darwin GOARCH=arm64 go build -o bin/arm64/darwin/cbshare
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o bin/amd64/linux/cbshare
	CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build -o bin/arm64/linux/cbshare
	GOOS=windows GOARCH=amd64 go build -o bin/amd64/windows/cbshare.exe
	GOOS=windows GOARCH=arm64 go build -o bin/arm64/windows/cbshare.exe