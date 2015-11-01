#!/usr/bin/env bash

set -e

target=${1}
time=${2}
name=${3}

cd ${target}
go test -benchmem -memprofile=mem.out -test.memprofilerate=1 -benchtime ${time}s -bench=${name}
go tool pprof -callgrind ${target}.test mem.out > mem.grind
kcachegrind mem.grind
