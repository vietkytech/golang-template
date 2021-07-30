ARG GO_VERSION=1.15.7-alpine
FROM docker.chotot.org/golang-builder:$GO_VERSION as builder
WORKDIR /go/src/git.chotot.org/fse/multi-rejected-reasons
ENV GOPRIVATE "git.chotot.org/*"
COPY go.mod go.sum /go/src/git.chotot.org/fse/multi-rejected-reasons
RUN go mod download
COPY . .
# RUN go build -o ./dist/multi-rejected-reasons
