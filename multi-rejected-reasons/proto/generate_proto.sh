#!/bin/bash

version=`protoc --version`
if [[ $version < "libprotoc 3.14.0" ]]; then
   echo "Please update your protobuf to at least version 3.14.0"
   exit 1
fi

go_protobuf_version=`git --git-dir="$GOPATH/src/github.com/golang/protobuf/.git" rev-parse HEAD`
PROTO_COMMIT="347cf4a86c1cb8d262994d8ef5924d4576c5b331"

if [[ $go_protobuf_version != $PROTO_COMMIT ]]; then
   echo "Updating github.com/golang/protobuf to ${PROTO_COMMIT} commit"
   go get -d -u github.com/golang/protobuf/{proto,protoc-gen-go}
	 git --git-dir=$GOPATH/src/github.com/golang/protobuf/.git --work-tree=$GOPATH/src/github.com/golang/protobuf/  checkout $PROTO_COMMIT
	 go install github.com/golang/protobuf/{proto,protoc-gen-go}
fi

BRANCH="master"
while [[ $# -gt 0 ]]
do
key="$1"
case $key in
	-b|--branch) 
	BRANCH="$2" 
	shift
	shift
	;;
	*)
	shift
esac
done

echo BRANCH = "${BRANCH}"

cd `dirname "$0"`;

project_dir=`pwd`
echo $project_dir

echo "generating proto"
protoc -I . ./multirr/*.proto --go-grpc_out=. --go_out=.