'''
This file contains methods to fit dataframes with technical indicators
to help prepare DataFrames for ML traning.

Modules used
- pandas

Author: Vikas Katari
Date: 05/12/2025
'''
import pandas as pd
from typing import Dict, Any
from talib import abstract

def delta(df: pd.DataFrame, json: Dict[str, Any], 
          col=None) -> pd.Series: # col unused -- for dispatcher
    '''Returns the delta or difference of deltas for a col'''
    col1 = json["col1"] # required

    if json["col2"] == None: # delta of single col
        deltas = df[col1].diff()
        return deltas
    else: # differences of deltas
        col2 = json["col2"]
        delta_diffs = (df[col1] - df[col2]).diff()
        return delta_diffs
    
def diff(df: pd.DataFrame, json: Dict[str, Any], 
         col=None) -> pd.Series: # col unused -- for dispatcher
    '''Returns the difference of two columns'''
    col1 = json["col1"]
    col2 = json["col2"]
    diffs = df[col1] - df[col2]
    return diffs

def put_technical(tech: str, df: pd.DataFrame, 
                  args: Dict[str, Any]) -> pd.Series:
    '''Calls TA-lib wrapper given a technical value, df, and args'''
    try:
        func = abstract.Function(tech) # <- TA-lib internal dispatch table 
        res = func(df, **args) 
        return res  
    except Exception:
        raise ValueError(f"\"{tech}\": not a valid technical, check config file")