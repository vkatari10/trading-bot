'''
File for testing trained ML models on historical data
to validate strategies on various tickers

Author: Vikas Katari 
Date: 06/21/2025
'''

import src.ml.data_processing.data_processing as dp
import src.ml.json.json_parser as jp
import src.ml.training.training as train
import src.backtesting.account as acct
import os
import pickle
import matplotlib.pyplot as plt
import numpy as np
import pandas as pd


# terminal formatting
from rich.progress import Progress
from rich.console import Console
from datetime import datetime


def log(message: str, style="bold white"):
    console = Console()
    now = datetime.now().strftime("%H:%M:%S")
    console.print(f"[{now}] [{style}]{message}[/{style}]")


def backtest(df: pd.DataFrame, stop: int, model, 
             settings: jp.UserConfig) -> None:
    account = acct.Account(
        cash=settings.get_backtesting_cash(),
        commission=settings.get_backtesting_commission(),
        batch_size=settings.get_backtesting_pos_size()
    )
    start_val = account.cash

    bs_list = []

    with Progress() as progress:

        task = progress.add_task(f"{datetime.now().strftime("%H:%M:%S")} Progress", total=len(df))

        for i in range(len(df)): # df.iloc[i, 0:stop] 
            pred_series = df.iloc[i, 0:stop]
            pred_df = pred_series.to_frame().T
            pred = model.predict(pred_df)
            bs_list.append(int(pred))
            if pred == 1:
                account.buy(df.iloc[i, 0])
            elif pred == -1:
                account.sell(df.iloc[i, 0])
            progress.update(task, advance=1)
            # if account.cash != val:
            #     print(account.cash)
            #     val = account.cash


    account.sell(df.iloc[len(df) - 1, 0])
    print(f"{start_val:.2f} -> {account.cash:.2f}")

    bs_series = pd.Series(bs_list)
    df['bf_col'] = bs_series
    
    # plt.plot(df.index, df.iloc[:, 0])
    # plt.title('Results')
    # plt.xlabel("date")
    # plt.ylabel("price")
    # plt.show()


def run_backtest(file_name: str, ticker: str) -> None:


    
    config = jp.UserConfig(file_name)
    log(f"Read {config.file_name}", style="green")

    if ticker in config.get_training_stocks():
        log(f"Note: {config.get_model_name()} was trained on ticker {ticker}", 
            style="yellow")

    df = dp.get_single_df(config, ticker)
    log(f"Testing data contains {len(df)} rows", style="green")

    # return 0.0
    
    # load ML model 

    with open('src/ml/models/decider/' + config.get_model_name(), 'rb') as f:
        model = pickle.load(f)
    log(f"Loaded {config.get_model_name()}", style="green")

    stop = train.find_stop(df, config)

    backtest(df, stop, model, config)

# add graph visualizer kind of like in backtrader

