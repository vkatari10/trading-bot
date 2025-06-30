'''
Script that will train a ML model against user config specifications, 
for all training stocks listed, and dumped using pickle

Author: Vikas Katari
Date: 05/28/2025
'''
import pickle
import logging

logging.basicConfig(level=logging.INFO)

# to process dataframes
import src.ml.data_processing.data_processing as dp

# to train models
import src.ml.training.training as train

# user config file 
import src.ml.json.json_parser as jp

# user config file
logging.info("Loading user config file")
config = jp.UserConfig()    

# get all training DFs w/ user features and labels
logging.info("Building training dataframe")
df = dp.get_df(config)

# find the stop column to prevent training on cols that are not features
stop = train.find_stop(df, config)

# train model
logging.info("Training model")
model = train.model_training(df, stop) 

# export model 
logging.info("Dumping model")
with open("src/ml/models/decider/" + config.get_model_name(), 'wb') as f:
    pickle.dump(model, f)

logging.info("Training done")
