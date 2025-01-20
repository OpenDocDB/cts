all: test build run lint

init:
	go mod tidy
	go mod verify

fmt:
	go tool goimports -local=github.com/OpenDocDB/cts/opendocdb-cts -w .
	go tool gofumpt -w .

test: fmt
	go test -count=1 ./...

build: fmt
	go run main.go --help

run:
	go run main.go --dir=../cts fmt
	go run main.go --dir=../cts run --uri=mongodb://127.0.0.1:27001/qklck6njgh9o --golden
	go run main.go --dir=../cts run --uri=mongodb://127.0.0.1:27001/cts

lint: fmt
	go tool go-consistent -pedantic ./...
	go tool golangci-lint run
	go tool govulncheck -test -show=verbose,color ./...

docker-init:
	-docker buildx create --name=opendocdb-cts --driver=docker-container --bootstrap=true --use=false

docker-local:
	docker buildx build --builder=opendocdb-cts --tag=opendocdb-cts --load ..
