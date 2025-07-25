'''
File for testing trained ML models on historical data
to validate strategies on various tickers

Author: Vikas Katari 
Date: 06/21/2025
'''

import src.ml.data_processing.data_processing as dp # DF fitting
import src.ml.json.json_parser as jp # user config
import src.backtesting.account as acct # account 
import pickle
import pandas as pd
import random

# terminal formatting
from rich.progress import Progress
from rich.console import Console
from datetime import datetime

def log(message: str, style="bold white"):
    console = Console()
    now = datetime.now().strftime("%H:%M:%S")
    console.print(f"[{now}] [{style}]{message}[/{style}]")

def do_backtest(
        df: pd.DataFrame,
        model,
        cash: float,
        commission: float,
        batch_size: int,
        slippage: float) -> None:
    
    account = acct.Account(
        cash=cash,
        commission=commission,
        batch_size=batch_size
    )

    start_val = cash

    bs_list = [] # track buy/sell

    with Progress() as progress:

        task = progress.add_task(f"[[bold green]{datetime.now().strftime("%H:%M:%S")}[/bold green]] Progress", total=len(df))
        for i in range(len(df)): 
    
            # price point
            pred_series = df.iloc[i, :-1] # assuming one label
            pred_df = pred_series.to_frame().T

            # get predictions
            pred = model.predict(pred_df)
            bs_list.append(int(pred))

            if pred == 1:
                account.buy(df.iloc[i, 0] + random.uniform(-slippage, slippage))
            elif pred == -1:
                account.sell(df.iloc[i, 0] + random.uniform(-slippage, slippage))
            progress.update(task, advance=1)


    account.sell(df.iloc[len(df) - 1, 0] + random.uniform(-slippage, slippage))
    print(f"{start_val:.2f} -> {account.cash:.2f}")

    bs_series = pd.Series(bs_list)
    df['bf_col'] = bs_series

def run_backtest(file_name: str, 
                 ticker: str,
                 cash, 
                 commission,
                 slippage, 
                 batch_size) -> None:

    config = jp.UserConfig(file_name)
    log(f"Read {config.file_name}", style="green")

    if ticker in config.get_training_stocks():
        log(f"Note: {config.get_model_name()} was trained on ticker {ticker}", 
            style="yellow")

    df = dp.get_single_df(config, ticker)
    log(f"Testing data contains {len(df)} rows", style="green")

    with open('models/' + config.get_model_name(), 'rb') as f:
        model = pickle.load(f)
    log(f"Loaded {config.get_model_name()}", style="green")

   
    # add default args here and check 
    if not cash:
        user_cash = config.get_backtesting_cash()
        cash = user_cash

    if not commission:
       user_commission = config.get_backtesting_commission()
       commission = user_commission

    if not slippage:
        user_slippage = config.get_backtesting_slippage()
        slippage = user_slippage
    
    
    if not batch_size:
        user_pos_size = config.get_backtesting_pos_size()
        batch_size = user_pos_size
    


    do_backtest(df, model, cash, commission, batch_size, slippage)

# add graph visualizer kind of like in backtrader

