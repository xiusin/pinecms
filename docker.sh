#!/usr/bin/env bash

COPY pinecms .

docker build -t xiusin/pinecms:0.0.1 .

docker run -p 2019:2019 xiusin/pinecms:0.0.1
