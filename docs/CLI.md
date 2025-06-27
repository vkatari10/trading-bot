# StratForge CLI 

To get quick jobs done without having the start the TUI, simply use the the CLI. 

## Starting

Simply run `./cli` from the root of the project

## Commands

| Command | Action | Args |
|:-------:|--------|------|
| build | Trains and dumps an ML model based on config file, compiles all dependencies |none |
| env | Starts the `venv` even if already started, updates the `.env` if any changes were made | none | 
| mlapi | Starts the ML API Sever | none | 
| test | Runs tests for both `Python` and `Go` source code | none |
| train | Trains and dumps an ML model based on config file | none | 