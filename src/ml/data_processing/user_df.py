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

    def __init__(self, config_file: str):
        self.config = jp.UserConfig(config_file) # JSON config itself
        self.stop = 1e-9

    def generate_model_yfinance(self) -> None:
        '''
        Generates a model using a YFinance DF from the config 
        file declared stocks
        '''
        import src.ml.training.training_methods as train # prevent partial module init
        training_df = dp.get_df(self.config)
        # self.stop = self.find_stop(training_df)
        train.train_dispatch(self, training_df)

    def generate_model_userdata(self, csv_file_path: str) -> None:
        '''Generates a model using a user CSV that contains OHLCV data'''
        import src.ml.training.training_methods as train
        ohlcv_df = pd.read_csv(csv_file_path)
        training_df = dp.process_data(ohlcv_df, self.config)
        self.stop = self.find_stop(training_df)
        # call training method here

    def generate_df_yfinance(self, dump_path: str) -> None:
        '''
        Generates a DataFrame with all user features and labels
        exports as a CSV

        Will concat all the dataframes together if multiple 
        tickers are declared in the config file
        '''
        df = dp.get_df(self.config)
        df.to_csv(dump_path)

    def generate_df_userdata(self, csv_file_path: str,
                             dump_path: str) -> None:
        '''
        Generates a DataFrame with all user features and labels
        given user data as a CSV exported as a CSV
        '''
        df = pd.read_csv(csv_file_path)
        df = dp.process_data(df, self.config)
        df.to_csv(dump_path)

    def get_single_df(self, ticker: str) -> pd.DataFrame:
        '''
        Returns a single df given a ticker with user
        features and labels
        '''
        return dp.get_single_df(self.config, ticker)
    
    def find_stop(self, df: pd.DataFrame) -> int:
        '''Find stop index where label column(s) begin (inclusive)'''
        return df.columns.get_loc(self.config.get_labels()[0]['name'])
    
