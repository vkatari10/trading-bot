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
import pandas as pd



def backtest(df, stop, model):
    account = acct.Account()
    start_val = account.cash

    for i in range(len(df)): # df.iloc[i, 0:stop] 
        pred_series = df.iloc[i, 0:stop]
        pred_df = pred_series.to_frame().T
        pred = model.predict(pred_df)
        if pred == 1:
            account.buy(df.iloc[i, 0])
        elif pred == -1:
            account.sell(df.iloc[i, 0])
        # if account.cash != val:
        #     print(account.cash)
        #     val = account.cash

    account.sell(df.iloc[len(df) - 1, 0])
    print(f"{start_val:.2f} -> {account.cash:.2f}")


def run_backtest(file_name: str, ticker: str) -> float:

    print("Reading configuration")
    config = jp.UserConfig(file_name)

    print("Creating backtesting dataframe")
    df = dp.get_single_df(config, ticker)

    # return 0.0
    
    # load ML model 
    print(f"Loading {config.get_model_name()}")
    with open('src/ml/models/decider/' + config.get_model_name(), 'rb') as f:
        model = pickle.load(f)

    stop = train.find_stop(df, config)


    print(f"Starting backtest on {len(df)} rows")
    backtest(df, stop, model)

# add graph visualizer kind of like in backtrader