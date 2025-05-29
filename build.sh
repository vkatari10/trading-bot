#!/usr/bin/bash

# script to compile all dependencies
# set up all local servers
# and train/load models



echo -e "Starting virtual environment"
source venv/bin/activate
echo -e "Startup Sucessful"

echo -e "===================================================\n"

echo -e "Reading features and training models\n"
python3 -m src.machine_learning.pipeline
echo -e "Loaded retrained model into src/machine_learning/models/decider\n"

echo -e "===================================================\n"

echo -e "Compiling C runtime dependencies"
cd src/runtime/recomputation/csrc/
./compile.sh
cd ../../../
echo -e "Compilation Sucessful"

echo "done"
