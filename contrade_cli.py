#! /usr/bin/env python3
'''
CLI Tool to interact with various services

Author: Vikas Katari 
Date: 06/30/2025
'''
import argparse
import subprocess

def train(args):
    import src.ml.pipeline as pipe
    pipe.pipeline_yfinance(args.config)

def build(args):
    subprocess.run(["bash", "./scripts/build.sh"])

def test(args):
    subprocess.run(["bash", "./scripts/test.sh"])

def run(args):
    subprocess.Popen(["go", "build"], cwd="./src/runtime/go-src").wait()
    process = subprocess.Popen(
        ["go", "run", ".", args.config],
        cwd="./src/runtime/go-src"
    )
    process.communicate()

def mlapi(args):
    subprocess.run(
        ["python", "-m", "src.api.internal.model_api.fast_model_api", args.config]
    )

def backtest(args):
    import src.backtesting.backtesting as bt
    bt.run_backtest(args.config,
                    args.asset,
                    args.cash,
                    args.commission,
                    args.slippage,
                    args.size)

# --- parser setup ---
parser = argparse.ArgumentParser(prog="contrade")
subparsers = parser.add_subparsers(dest="command")

# train
p_train = subparsers.add_parser("train", help="Train a model using a config file")
p_train.add_argument("config", help="Path to config file (e.g. config/xyz.json)")
p_train.set_defaults(func=train)

# build
p_build = subparsers.add_parser("build", help="Run the build script")
p_build.set_defaults(func=build)

# test
p_test = subparsers.add_parser("test", help="Run the test script")
p_test.set_defaults(func=test)

# run
p_run = subparsers.add_parser("run", help="Run Go execution engine with config")
p_run.add_argument("config", help="Path to config file (e.g. config/xyz.json)")
p_run.set_defaults(func=run)

# mlapi
p_mlapi = subparsers.add_parser("mlapi", help="Run fast ML API with config")
p_mlapi.add_argument("config", help="Path to config file (e.g. config/xyz.json)")
p_mlapi.set_defaults(func=mlapi)

# backtest
p_backtest = subparsers.add_parser("backtest", help="Run backtest with config and asset")
p_backtest.add_argument("config", help="Path to config file (e.g. config/xyz.json)")
p_backtest.add_argument("asset", help="Asset ticker symbol (e.g. AAPL)")
p_backtest.set_defaults(func=backtest)

# optional args
p_backtest.add_argument("--cash", type=float, default=None, help="Initial cash balance")
p_backtest.add_argument("--commission", type=float, default=None, help="Commission per trade")
p_backtest.add_argument("--slippage", type=float, default=None, help="Random range where buy/sell trades could be adjusted by")
p_backtest.add_argument("--size", type=int, default=None, help="How many shares should be bought on every trade")

try:
    args = parser.parse_args()
    if hasattr(args, "func"):
        args.func(args)
    else:
        parser.print_help()
except ModuleNotFoundError:
    print("venv has not been started, run 'source venv/bin/activate'")
except KeyboardInterrupt:
    print("quit")
