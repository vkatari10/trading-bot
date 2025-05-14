'''
This file contains the runtime environment to run ML models
on real time data and place buy/sell orders

Modules Used
- Pickle
- Time
- Pandas

Author: Vikas Katari
Date: 05/13/25
'''
import pickle as pkl
import time
import pandas as pd

import src.finnhub_api.market_data as fb
import src.alpaca_api.ordering as ao
import src.live_data.live_data as ld

model_file = "src/models/model_v1.pkl"

with open(model_file, 'rb') as f:
    pkl.load(f)

df = pd.read_csv("notebooks/refined_df.csv")
print(df)

ticker = "AAPL"

while (True):
    try:
        stock_price = fb.get_quote(ticker)
        print(f"Quote ({ticker}): {stock_price['c']}")

        # ADD this to the historical data chart to get new indicators

        # Feed new technicals to ML bot

        # make predict

        # based on predict call alpaca to buy or sell

        '''
        recompute(stock_price, df)
        '''



        time.sleep(60)
    except KeyboardInterrupt:
        print("exit")
        break

# make_csv(df)
