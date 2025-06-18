# Scripts

This folder contains scripts to start various services. They should be called from the root of this project directory like `./scripts/x.sh` unless noted.<br>

Purpose:
- `build.sh`
    - Trains and Dumps ML model based on defined JSON in `config/` and settings in `.ml_env`
    - Compiles `C` and `Go` source code
- `env.sh`
    - Starts the virtual environment assuming the virtual environment for `Python` packages is named `venv`
    - NOTE: You need to `source` this file rather than executing it
- `mlapi.sh`
    - This will start the ML API Server that the runtime engine can then interact with 
- `test.sh`
    - This will execute all tests for `Python` and `Go` source code
- `train.sh`
    - Trains and Dumps the ML model based on defined JSON in `config/` and setttings in `.ml_env`