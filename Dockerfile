ARG GO_VERSION=1.15.15-buster
FROM docker.chotot.org/golang-builder:$GO_VERSION as builder
WORKDIR /go/src/github.com/vietkytech/golang-template/
ENV GOPRIVATE "git.chotot.org/*"
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /go/src/github.com/vietkytech/golang-template/gotemplate
RUN go build -o ./dist/gotemplate

FROM debian:buster-20210816-slim
RUN apt-get update
RUN apt-get install -y ca-certificates
RUN apt-get install -y tzdata && \
  cp -f /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime && \
  apt-get remove -y tzdata
WORKDIR /app
EXPOSE 8080
COPY --from=builder /go/src/github.com/vietkytech/golang-template/gotemplate/dist/gotemplate .
COPY --from=builder /go/src/github.com/vietkytech/golang-template/gotemplate/config ./config/
COPY --from=builder /go/src/github.com/vietkytech/golang-template/gotemplate/templates ./templates/
COPY --from=builder /go/src/github.com/vietkytech/golang-template/run_all.sh .
COPY ./run_all.sh /usr/bin/
CMD ["sh","/usr/bin/run_all.sh", ""]