'''
Class implementation representing a user DataFrame
given their configuration specs

Author: Vikas Katari
Date: 07/11/2025
'''

import src.ml.json.json_parser as jp # user config
import src.ml.data_processing.data_processing as dp
import pandas as pd 

class UserMLConfig():

    def __init__(self, user_config: jp.UserConfig):
        # df settings
        self.features = user_config.get_features()
        self.label_logic = user_config.get_labels()
        self.tickers = user_config.get_training_stocks()

        # all other settings
        self.config = user_config

    '''
    Put user defined features and labels and final label(s)

    Put OHLCV diffs if they want that (optional)

    Add method to just train the ML model and then dump it if the df exists ig
    '''

    def get_training_df(self) -> pd.DataFrame:
        return dp.get_df(self.config)
    
    def get_single_df(self, ticker: str) -> pd.DataFrame:
        return dp.get_single_df(self.config, ticker)
    
    def train_dump_model(self) -> None:
        # call training method based on sci kit learn or not
        # training method should dump model as well 

        # specs
        # takes in DF    -> return None (dump model)
        
        pass

    

