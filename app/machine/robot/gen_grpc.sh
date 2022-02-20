#!/bin/sh

grpc_cpp_plugin_location=$(which grpc_cpp_plugin)

protoc --proto_path=./protos/ --grpc_out=./command_service --plugin=protoc-gen-grpc="$grpc_cpp_plugin_location" command.proto
printf "Generate 'command.proto' message classes finished\n"

protoc --proto_path=./protos/ --cpp_out=./command_service command.proto
printf "Generate 'command.proto' service classes finished\n"
