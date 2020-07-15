#!/usr/bin/env bash
docker run --rm -v "$GOPATH"/src:/go/src -v "$PWD":/usr/src/pinecms -w /usr/src/pinecms golang:1.14.2-alpine3.11 sh -c "apk add git && apk add gcc && apk add g++ && go build"

docker build -t xiusin/pinecms:0.0.1 .

docker run -p 2019:2019 xiusin/pinecms:0.0.1