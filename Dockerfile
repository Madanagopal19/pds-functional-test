# Build the proxy binary.
FROM golang:1.16.6 as builder

WORKDIR /workspace

# Dependencies.
COPY go.mod go.mod
COPY go.sum go.sum
COPY vendor/ vendor/
COPY Makefile Makefile

# Source.
COPY test/ test/
COPY pkg/ pkg/

CMD go test ./test -timeout 99999999s -v