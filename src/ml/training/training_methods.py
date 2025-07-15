'''
File containing model framework agnostic methods
that are useful for training

Author: Vikas Katari 
Date: 07/11/2025
'''

import pickle
import pandas as pd

import src.ml.data_processing.user_df as ud
import src.ml.training.training as train

ML_TRAINING_DISPATCH = {
    "lightgbm": train.lightgbm_train,
    "xgboost": train.xgboost_train,
    "scikit": train.scikit_train
} 

def train_dispatch(config: ud.UserMLConfig, df: pd.DataFrame) -> None:
    '''Trains and dumps a model given user config specifications'''
    return ML_TRAINING_DISPATCH[config.config.get_model_framework()](config, df)


def dump_model(user_config: ud.UserMLConfig) -> None:
    with open('src/ml/models/decider' + user_config.config.get_model_name(), 
              'wb') as f:
        pickle.dump(f)

