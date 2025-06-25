'''
File for testing trained ML models on historical data
to validate strategies on various tickers

Author: Vikas Katari 
Date: 06/21/2025
'''

import src.ml.data_processing.data_processing as dp
import src.ml.training.training as train
import src.backtesting.account as acct
import os
import pickle
from dotenv import load_dotenv

load_dotenv('.env')



# if buy or sell use account to buy / sell 

# if sell just dump everything for now

# tally total cash or account value as is (sum value of positions + cash at end)

# return the value

def backtest(df, stop, model):
    account = acct.Account()

    val = 0

    for i in range(len(df)): # df.iloc[i, 0:stop] 
        pred = model.predict([df.iloc[i, 0:stop]])

        if pred == 1:
            account.buy(df.iloc[i, 0])
        elif pred == -1:
            account.sell(df.iloc[i, 0])


        if account.cash != val:
            print(account.cash)
            val = account.cash

    account.sell(df.iloc[len(df) - 1, 0])
    

    print(account.cash)


def run_backtest(ticker: str) -> float:
    # fit df 
    df = dp.get_df(ticker)

    # find stop val
    stop = train.find_stop(df, os.getenv("LABEL_CONFIG_FILE"))
    
    # load ML model 
    with open('src/ml/models/decider/' + os.getenv("MODEL_RUNTIME_NAME"), 'rb') as f:
        model = pickle.load(f)

    backtest(df, stop, model)

run_backtest("MSFT")