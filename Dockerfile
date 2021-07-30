ARG GO_VERSION=1.15.7-alpine
FROM docker.chotot.org/golang-builder:$GO_VERSION as builder
WORKDIR /go/src/git.chotot.org/fse/multi-rejected-reasons/
ENV GOPRIVATE "git.chotot.org/*"
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /go/src/git.chotot.org/fse/multi-rejected-reasons/multi-rejected-reasons
RUN go build -o ./dist/multi-rejected-reasons

FROM alpine:3.11
RUN apk add --update ca-certificates
RUN apk add --no-cache tzdata && \
  cp -f /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime && \
  apk del tzdata
WORKDIR /app
EXPOSE 8080
COPY --from=builder /go/src/git.chotot.org/fse/multi-rejected-reasons/multi-rejected-reasons/dist/multi-rejected-reasons .
COPY --from=builder /go/src/git.chotot.org/fse/multi-rejected-reasons/multi-rejected-reasons/config ./config/
COPY --from=builder /go/src/git.chotot.org/fse/multi-rejected-reasons/run_all.sh .
COPY ./run_all.sh /usr/bin/
CMD ["sh","/usr/bin/run_all.sh"]