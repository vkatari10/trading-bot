#! /usr/bin/env python3
'''
CLI Tool to interact with various services

Author: Vikas Katari 
Date: 06/30/2025
'''
try:

    import sys # argv
    import subprocess # run shell scripts
    import src.ml.pipeline as pipe # to train models
    import src.backtesting.backtesting as bt # backtesting module

    args = sys.argv

    # TODO implement error handling

    if args[1] == "train":
        pipe.pipeline(args[2]) 
    elif args[1] == "build":
        subprocess.run(["bash", "./scripts/build.sh"])
    elif args[1] == "test":
        subprocess.run(["bash", "./scripts/test.sh"])
    elif args[1] == "run":
        file = args[2]       

        compile = subprocess.Popen(
            ["go", "build"],
            cwd="./src/runtime/go-src"
        )

        process = subprocess.Popen( # forked
            ["go", "run", ".", file],
            cwd="./src/runtime/go-src"
        )

        process.communicate()
    elif args[1] == "mlapi":
        if len(args) == 3:
            subprocess.run(["bash", "python3", "-m", "src.api.internal.model_api.model_api", args[2]])
        else: 
            print("usage: ./contrade_cli.py mlapi <PATH_TO_CONFIG_FILE>")
    elif args[1] == "backtest":
        if len(args) == 4:
            bt.run_backtest(args[2], args[3])
        else:
            print("usage: ./contrade_cli.py backtest <PATH_TO_CONFIG_FILE> <ASSET_TICKER>")
except ModuleNotFoundError:
    print("venv has been not started, run 'source venv/bin/activate'")
except KeyboardInterrupt:
    print('quit')