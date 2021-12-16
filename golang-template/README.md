# multi-rejected-reasons

Multi RR

# run
- `export GOPRIVATE=git.chotot.org/*`
- run server `go run multi-rejected-reasons/main.go`
- run client `go run multi-rejected-reasons/main.go client --address=localhost:8081`

# generating proto
- `bash multi-rejected-reasons/proto/generate_proto.sh`
- run `bash multi-rejected-reasons/proto/install_protoc.sh` if there is an error asking for protoc version
