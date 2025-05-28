'''
This file contains methods to fit dataframes with technical indicators
to help prepare DataFrames for ML traning.

Modules used
- pandas

Author: Vikas Katari
Date: 05/12/2025
'''


import pandas as pd


def ema(df: pd.DataFrame, window: int, col="Close") -> pd.Series:
    '''Returns the EMA of a given column in a DataFrame'''
    emas = df[col].ewm(span=window, adjust=False).mean()
    return emas


def sma(df: pd.DataFrame, window: int, col="Close") -> pd.Series:
    '''Returns the SMA of a given column in a DataFrame'''
    smas = df[col].rolling(window).mean()
    return smas


def subtract(df: pd.DataFrame, col1: str, col2: str) -> pd.Series:
    '''
    Returns a new data frame column where every value is
    computed where col1 - col2
    '''
    diffs = df[col1] - df[col2]
    return diffs
