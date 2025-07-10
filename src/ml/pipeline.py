'''
Script that will train a ML model against user config specifications, 
for all training stocks listed, and dumped using pickle

Author: Vikas Katari
Date: 05/28/2025
'''
import pickle
import logging

from rich.console import Console
from datetime import datetime

def log(message: str, style="bold white"):
    console = Console()
    now = datetime.now().strftime("%H:%M:%S")
    console.print(f"[{now}] [{style}]{message}[/{style}]")

# to process dataframes
import src.ml.data_processing.data_processing as dp

# to train models
import src.ml.training.training as train

# user config file 
import src.ml.json.json_parser as jp

def pipeline(file: str) -> None:

    # user config file
    config = jp.UserConfig(file)    
    log(f"Read {file}", style="green")

    # get all training DFs w/ user features and labels

    df = dp.get_df(config)
    log(f"Created all ({len(config.get_training_stocks())}) training DataFrames",
        style="green")
    # find the stop column to prevent training on cols that are not features
    stop = train.find_stop(df, config)

    # train model
    model = train.model_training(df, stop) 
    log(f"Trained ML model: {config.get_model_name()}", style="green")

    # export model 
    with open("src/ml/models/decider/" + config.get_model_name(), 'wb') as f:
        pickle.dump(model, f)
    log(f"Dumped ML model: {config.get_model_name()}", style="green")

