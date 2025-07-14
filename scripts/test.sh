#! /usr/bin/bash

printf "%s\n" "Running Python Tests"

PYTHONPATH=. pytest

printf "%s" "============================"
printf "%s\n" "Running Golang Tests"
cd src/runtime/go-src/tests
go test