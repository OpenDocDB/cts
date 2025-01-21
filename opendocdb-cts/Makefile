all: test build run lint

init:
	go mod tidy && go mod verify
	cd tools && go mod tidy && go mod verify && go generate

fmt:
	../bin/goimports -local=github.com/OpenDocDB/cts/opendocdb-cts -w .
	../bin/gofumpt -w .

test: fmt
	go test -count=1 ./...

build: fmt
	go run main.go --help

run:
	go run main.go --dir=../cts fmt
	go run main.go --dir=../cts run --uri=mongodb://127.0.0.1:27001/qklck6njgh9o --golden
	go run main.go --dir=../cts run --uri=mongodb://127.0.0.1:27001/cts

lint: fmt
	../bin/go-consistent -pedantic ./internal/...
	../bin/golangci-lint run
	../bin/govulncheck -test -show=verbose,color ./...

docker-init:
	-docker buildx create --name=opendocdb-cts --driver=docker-container --bootstrap=true --use=false

docker-local:
	docker buildx build --builder=opendocdb-cts --tag=opendocdb-cts --load ..
