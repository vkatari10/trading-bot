'''
JSON parser to break down user config files to their different 
components 

Author: Vikas Katari
Date: 06/25/2025
'''
from typing import Dict, List, Any
import json

class UserConfig():
    '''
    Loads the MAIN_CONFIG_FILE, the JSON file declared in the 
    .env file and loads it into Python as a dict
    '''

    def __init__(self, file_name: str):
        try :
            with open(file_name) as f:
                file = json.load(f)
            self.file = file
            self.file_name = file_name
        except Exception:
            raise ValueError("File does not exist or could not be read")

    def get_features(self) -> List[Dict[str, Any]]:
        try:
            features = self.file['features']
            return features
        except KeyError:
            raise ValueError(make_error_str("features"))

    def get_labels(self) -> List[Dict[str, Any]]:
        try:
            labels = self.file['label_logic']
            return labels
        except KeyError:
            raise ValueError(make_error_str('label_logic'))
    
    def get_training_stocks(self) -> List[str]:
        try :
            training_stocks = self.file['train_stocks']
            return training_stocks
        except KeyError:
            raise ValueError(make_error_str('train_stocks'))

    def get_model_type(self) -> str:
        try:
            model_type = self.file['ml_settings']['model_type']
            return model_type
        except KeyError:
            raise ValueError(make_error_str('model_type'))

    def get_model_name(self) -> str:
        try:
            model_name = self.file['ml_settings']['model_name']
            return model_name
        except KeyError:
            raise ValueError(make_error_str('model_name'))

    def get_model_training_timeframe(self) -> str:
        try:
            model_name = self.file['ml_settings']['model_training_timeframe']
            return model_name
        except KeyError:
            raise ValueError(make_error_str('model_training_timeframe'))

    def get_model_training_interval(self) -> str:
        try:
            model_name = self.file['ml_settings']['model_training_interval']
            return model_name
        except KeyError:
            raise ValueError(make_error_str('model_training_interval'))

    def get_backtesting_cash(self) -> float:
        try:
            cash = self.file['backtest_settings']['starting_cash']
            return cash
        except KeyError:
            raise ValueError(make_error_str('starting_cash'))
        
    def get_backtesting_commission(self) -> float:
        try:
            comish = self.file['backtest_settings']['commission']
            return comish
        except KeyError:
            raise ValueError(make_error_str('commission'))
        
    def get_backtesting_pos_size(self) -> int:
        try:
            size = self.file['backtest_settings']['position_size']
            return size
        except KeyError:
            raise ValueError(make_error_str('position_size'))    

def make_error_str(key: str) -> str:
    return f"Could not parse '{key}' properly from the " \
    "JSON config"