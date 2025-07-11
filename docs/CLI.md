# ConTrade CLI 

To get quick jobs done rom CLI, you can call the `./contrade_cli.py` tool. 

## Commands

| Command | Description | Usage |
|:-------:|--------|------|
| build | Compiles `Go` source code and `C` source code if applicable | `./contrade_cli.py build` |
| mlapi | Starts ML API Server, should be called before the `run` command to ensure inferences can be made | `./contrade_cli.py mlapi <CONFIG_FILE_PATH>` | 
| run | Starts Live execution with a given JSON configuration file, the path of the config file should be relative from the root, you will need to call `mlapi` command as well from another terminal to ensure that the ML model can make live inferences. **NOTE:** make sure you are using the same config file as the one used to call `mlapi` | `./contrade_cli.py run <CONFIG_FILE_PATH>` | 
| test | Runs all unit tests for both `Python` and `Go` source files | `./contrade_cli.py test` |
| train | Trains ML model given a JSON configuration file, the path of the config file should be relative from the root | `./contrade_cli.py train <CONFIG_FILE_PATH>` | 