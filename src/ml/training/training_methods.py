'''
File containing model framework agnostic methods
that are useful for training

Author: Vikas Katari 
Date: 07/11/2025
'''

import pickle
import pandas as pd

import src.ml.data_processing.user_df as ud

def dump_model(user_config: ud.UserMLConfig) -> None:
    with open('src/ml/models/decider' + user_config.config.get_model_name(), 
              'wb') as f:
        pickle.dump(f)

def find_stop(user_config: ud.UserMLConfig) -> int:
    '''Finds the index of the first label column'''
    return 5 # STUB 

