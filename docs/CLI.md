# StratForge CLI 

To get quick jobs done rom CLI, you can call the `./contrade_cli.py` tool. 

## Commands

| Command | Description | Usage |
|:-------:|--------|------|
| run | Starts Live execution with a given JSON configuration file, the path of the config file should be relative from the root | `./contrade_cli.py run <CONFIG_FILE_PATH>` | 
| test | Runs all unit tests for both `Python` and `Go` source files | `./contrade_cli.py test` |
| train | Trains ML model given a JSON configuration file, the path of the config file should be relative from the root | `./contrade_cli.py train <CONFIG_FILE_PATH>` | 