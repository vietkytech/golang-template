ARG GO_VERSION=1.15.7-alpine
FROM docker.chotot.org/golang-builder:$GO_VERSION as builder
WORKDIR /go/src/github.com/vietkytech/golang-template
ENV GOPRIVATE "git.chotot.org/*"
COPY go.mod go.sum /go/src/github.com/vietkytech/golang-template
RUN go mod download
COPY . .