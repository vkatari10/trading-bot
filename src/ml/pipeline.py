'''
machine_learning folder script that allows user the define feature
relationships and call all methods to train and export the ML models.

Author: Vikas katari
Date: 05/28/2025
'''

import pandas as pd
import pickle
import json

# to process dataframes
import src.ml.data_processing.data_processing as dp

# to train models
import src.ml.training.training as train

training_ticker = "AAPL"

df = dp.get_df(training_ticker) # DO NOT MODIFY


with open("src/logic/signals.json") as f:
    signals = json.load(f)

stop = df.columns.get_loc((signals[0]['name'], ''))

model = train.model_training(df, stop, 4)

# export model to runtime destination
with open('src/ml/models/decider/model.pkl', 'wb') as f:
    # replace the df with the actual model when its compelte
    pickle.dump(model, f)
