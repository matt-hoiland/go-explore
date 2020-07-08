#!/bin/bash

protoc \
    --go_out=plugins=grpc:. \
    --descriptor_set_out=echo/echopb/echo.protoset \
    echo/echopb/echo.proto
