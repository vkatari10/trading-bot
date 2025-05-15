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

import src.finnhub_api.market_data as fh
import src.alpaca_api.ordering as ao
import src.live_data.live_data as ld
from sklearn.ensemble import RandomForestClassifier


print("Opening ML model")

model_file = "src/models/model_v1.pkl"

with open(model_file, 'rb') as f:
    model = pkl.load(f)

print("Opening Historical Data")

# when finished rename the exported df to this file name below
df = pd.read_csv("src/engine/refined_df.csv")

ticker = "AAPL"

while (True):
    try:
        price = fh.get_quote(ticker)
        print(f"New Quote ({ticker}): {price['c']}")
        df = ld.append_new(df, price['c'], price['h'],
                           price['l'], price['o'])

        X = df.iloc[[-1]].drop(columns="ACTION")
        prediction = model.predict(X)
        action = prediction[0]

        if action > 0:
            print("BUY SIG") # hook alpaca orders here
            ao.place_market_order(ticker, 1, "buy")
        elif action < 0:
            ao.place_market_order(ticker, 1, "sell")
        else:
            print("HOLD SIG")

        time.sleep(60)
    except KeyboardInterrupt:
        print("saving collected data")
        # rename this to the thing above when opening the df
        # later we can just name the original file name to overwrite it
        ld.make_csv('src/engine/runtime_df.csv', df)
        print("saved")
        print("done")
        break
