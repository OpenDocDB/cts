FROM --platform=$BUILDPLATFORM golang:1.24.3

ENV GOPATH=/cache/gopath
ENV GOCACHE=/cache/gocache
ENV GOMODCACHE=/cache/gomodcache

WORKDIR /src/opendocdb-cts

RUN --mount=target=/src --mount=type=cache,target=/cache <<EOF
set -ex

go mod download
go mod verify
EOF

RUN --mount=target=/src --mount=type=cache,target=/cache <<EOF
set -ex

export GOOS=$TARGETOS
export GOARCH=$TARGETARCH
export CGO_ENABLED=0

go build -v -o /opendocdb-cts
/opendocdb-cts --help
EOF

WORKDIR /src

ENTRYPOINT [ "/opendocdb-cts", "run", "--uri=mongodb://mongodb:27001/cts"]