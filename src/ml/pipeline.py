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

load_dotenv('.env')

# to process dataframes
import src.ml.data_processing.data_processing as dp

# to train models
import src.ml.training.training as train

# get DF w/ user features and labels
df = dp.get_df(os.getenv("TRAIN_TICKER"))

# Load features
with open("config/" + os.getenv("LABEL_CONFIG_FILE")) as f:
    signals = json.load(f)

# stop index to exclude relationship columns
# signals[0]['name'] is just the first relationship col_name
stop = df.columns.get_loc((signals[0]['name'], ''))

# train model
model = train.model_training(df, stop) 

# export model to runtime destination
with open("src/ml/models/decider/" + os.getenv("MODEL_DUMP_NAME"), 'wb') as f:
    pickle.dump(model, f)
