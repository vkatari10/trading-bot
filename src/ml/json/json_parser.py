'''
JSON parser to break down user config files to their different 
components 

Author: Vikas Katari
Date: 06/25/2025
'''
from typing import Dict, List, Any, Optional
import json

class UserConfig():
    '''
    Represents the JSON config file to easily get
    grab data with wrapper methods
    '''

    def __init__(self, file_name: str):
        try :
            with open(file_name) as f:
                file = json.load(f)
            self.file = file
            self.file_name = file_name
        except Exception:
            raise ValueError("File does not exist or could not be read")
        
    def get(self, *args: str) -> Any:
        f = self.file
        res = None
        for arg in args:
            f = f.get(arg)
        res = f
        if res is None:
            raise ValueError(f"{args} could not be parsed")
        return res
    
    # training/stock data 

    def get_features(self) -> List[Dict[str, Any]]:
        return self.get('features')

    def get_labels(self) -> List[Dict[str, Any]]:
        return self.get('label_logic')
    
    def get_training_stocks(self) -> List[str]:
        return self.get('train_stocks')
        
    def get_live_stocks(self) -> List[str]:
        return self.get('live_trade_stocks')
    
    # ML settings

    def get_model_type(self) -> str:
        return self.get('ml_settings', 'scikit_model_type')
    
    def get_model_framework(self) -> str:
        return self.get('ml_settings', 'model_framework')
    
    def get_model_name(self) -> str:
        return self.get('ml_settings', 'model_name')

    def get_model_training_timeframe(self) -> str:
        return self.get('ml_settings', 'model_training_timeframe')

    def get_model_training_interval(self) -> str:
        return self.get('ml_settings', 'model_training_interval')
    
    def get_ml_settings(self) -> Dict[str, Any]:
        return self.get('ml_settings')
        
    def get_OHLCV_diffs_setting(self) -> bool:
        return self.get('ml_settings', 'use_OHLCV_diffs')

    def get_hyperparameters(self) -> Dict[str, Any]:
        return self.get('ml_settings', "hyperparameters")

    # Backtest settings

    def get_backtesting_cash(self) -> float:
        return self.get('backtest_settings', 'starting_cash')

    def get_backtesting_commission(self) -> float:
        return self.get('backtest_settings', 'commission')
    
    def get_backtesting_pos_size(self) -> int:
        return self.get('backtest_settings', 'position_size')

   