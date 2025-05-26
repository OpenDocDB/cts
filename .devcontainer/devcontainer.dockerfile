FROM golang:1.24.3

ENV GOPATH=/cache/gopath
ENV GOCACHE=/cache/gocache
ENV GOMODCACHE=/cache/gomodcache

# Install necessary tools
RUN apt-get update && apt-get install -y \
    git \
    curl \
    make \
    zip \
    unzip \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /src