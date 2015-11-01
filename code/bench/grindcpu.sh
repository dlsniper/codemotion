#!/usr/bin/env bash

set -e

target=${1}
time=${2}
name=${3}

cd ${target}
go test -cpuprofile=cpu.out -benchtime ${time}s -bench=${name}
go tool pprof -callgrind -alloc_space ${target}.test cpu.out > cpu.grind
kcachegrind cpu.grind
