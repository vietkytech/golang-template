#!/bin/bash

PB_REL="https://github.com/protocolbuffers/protobuf/releases"
PROTOC_VERSION=3.14.0
curl -LO $PB_REL/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip

unzip protoc-${PROTOC_VERSION}-linux-x86_64.zip -d $HOME/.local

# go get -u github.com/carousell/Orion/protoc-gen-orion

export PATH="$PATH:$HOME/.local/bin"

