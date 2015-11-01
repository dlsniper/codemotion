#!/usr/bin/env bash

target=${1}

cd ${target}
go build -gcflags -S bench${target}.go