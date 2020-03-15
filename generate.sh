#!/bin/bash

# generating protobuf files of greeting service
protoc greet/greet.proto --go_out=plugins=grpc:.

# generating protobuf files of calculator service
protoc calculator/calculator.proto --go_out=plugins=grpc:.