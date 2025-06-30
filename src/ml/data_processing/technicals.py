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

# all method default on close values right now for computations

def ema(df: pd.DataFrame, json: Dict[str, Any], 
        col="Close") -> pd.Series: # default on close values 
    '''Returns the EMA of a given column in a DataFrame'''
    window = json["window"] # does not account for smoothing values yet 
    emas = df[col].ewm(span=window, adjust=False).mean()
    return emas

def sma(df: pd.DataFrame, json: Dict[str, Any], 
        col="Close") -> pd.Series: 
    '''Returns the SMA of a given column in a DataFrame'''
    window = json["window"]
  
    smas = df[col].rolling(window).mean()
    return smas

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

