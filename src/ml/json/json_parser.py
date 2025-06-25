'''
JSON parser to break down user config files to their different 
components 

Author: Vikas Katari
Date: 06/25/2025
'''
import os
from dotenv import load_dotenv
from typing import Dict, List, Any
import json
load_dotenv('.env')

class UserConfig():
    '''
    Loads the MAIN_CONFIG_FILE, the JSON file declared in the 
    .env file and loads it into Python as a dict
    '''

    def __init__(self):
        with open("config/" + os.getenv("MAIN_CONFIG_FILE")) as f:
            file = json.load(f)
        self.file = file

    def get_features(self) -> List[Dict[str, Any]]:
        return self.file['features']

    def get_labels(self) -> List[Dict[str, Any]]:
        return self.file['label_logic']

    def get_training_stocks(self) -> List[str]:
        return self.file['train_stocks']

    def get_model_type(self) -> str:
        return self.file['model_type']

    def get_model_name(self) -> str:
        return self.file['model_name']

    def get_model_training_timeframe(self) -> str:
        return self.file['model_training_timeframe']

    def get_model_training_interval(self) -> str:
        return self.file['model_training_interval']