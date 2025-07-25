# ConTrade CLI 

To get quick jobs done rom CLI, you can call the `./contrade_cli.py` tool. 

## Training

`./contrade_cli.py train <PATH_TO_CONFIG_FILE>`

The train command will parse, train, and dump the ML model according to your specifications

## Backtesting

`./contrade_cli.py backtest <PATH_TO_CONFIG_FILE> <ASSET_TICKER> --cash 10000 --commission 0.00 --slippage 0.03 --size 1`

The backtest command will backtest your model against the specified granularity/timeframe given a valid Ticker to backtest on.

All arguments starting with `--` are optional and **do not** need to be provided in your JSON config already contains default settings for backtest settings. In this case you just need to specify the asset ticker.

If your JSON config **does not** contain default settings then you **must** provide all optional arguments then. If you have default arguments in your JSON, you can provide **any** number of optional arguments and override them from the CLI.

## Run Live

You will **need** to use 2 terminals, and do them **in order** as shown.

1. `./contrade_cli.py mlapi <PATH_TO_CONFIG_FILE>`
2. `./contrade_cli.py run <PATH_TO_CONFIG_FILE>`

The first command will load the model onto its own server, then the run command will connect to the ML model server. The config file **must** be the same when using both commands, or undefined behavior will occur. 

## Build Dependencies

`./contrade_cli.py build`

The build command just compiles all required `Go` source code, and `C` code if applicable. Should only be used during the setup phase

## Run Tests

`./contrade_cli.py test`

The test command will run test for both `Python` and `Go` source code. CI pipelines already exist to automate these tests on push and pull requests, however they take around 2 minutes to complete, since it needs to set up the environment from scratch every time. 