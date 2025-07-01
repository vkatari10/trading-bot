#!/usr/bin/bash

# script to compile all dependencies
# set up all local servers
# and train/load models

echo -e "Starting virtual environment"
source venv/bin/activate

echo -e "==================================================="

echo -e "Compiling C dependencies"
cd src/runtime/c-src/
./compile.sh
cd ../../..

echo -e "==================================================="

echo -e "Compiling Go dependencies"
cd src/runtime/go-src/
go build
cd ../../..

echo -e "==================================================="

echo "Build Stage Done"