'''
machine_learning folder script that allows user the define feature
relationships and call all methods to train and export the ML models.

Author: Vikas katari
Date: 05/28/2025
'''

import pandas as pd
import pickle
import json
import os   
from dotenv import load_dotenv

load_dotenv('.ml_env')

# to process dataframes
import src.ml.data_processing.data_processing as dp

# to train models
import src.ml.training.training as train

training_ticker = os.getenv("TRAIN_TICKER")
label_file = os.getenv("LABEL_CONFIG_FILE")

df = dp.get_df(training_ticker) # DO NOT MODIFY FROM BELOW

label_file = "config/" + label_file
with open(label_file) as f:
    signals = json.load(f)

# stop index to exclude relationship columns
# signals[0]['name'] is just the first relationship col_name
stop = df.columns.get_loc((signals[0]['name'], ''))

model = train.model_training(df, stop) 

# export model to runtime destination
with open('src/ml/models/decider/model.pkl', 'wb') as f:
    # replace the df with the actual model when its compelte
    pickle.dump(model, f)
