# Build the proxy binary.
FROM golang:1.16.6 as builder
FROM bitnami/kubectl:1.23.4 as kubectl 

WORKDIR /workspace

# Dependencies.
COPY go.mod go.mod
COPY go.sum go.sum
COPY vendor/ vendor/
COPY Makefile Makefile

# Source.
COPY test/ test/
COPY pkg/ pkg/

# Installing helm and kubectl 
COPY --from=kubectl /opt/bitnami/kubectl/bin/kubectl /usr/local/bin/
ENTRYPOINT curl https://raw.githubusercontent.com/helm/helm/master/scripts/get > get_helm.sh;chmod 700 get_helm.sh; ./get_helm.sh ; sleep 10

CMD go test ./test -timeout 99999999s -v