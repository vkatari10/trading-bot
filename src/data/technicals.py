'''
This file contains method to compute technicals of historical stock data

Modules used
- pandas
- numpy

Author: Vikas Katari
Date: 05/12/2025
'''
import pandas as pd
import numpy as np
from pandas import DataFrame

def ema(df: DataFrame, days: int, col="Close", name="EMA") -> DataFrame:
    '''Adds an EMA column based on Close Prices to a DataFrame'''
    label = f'{name}({days})'
    df[label] = df[col].ewm(span=days, adjust=False).mean()
    return df

def sma(df: DataFrame, days: int, col="Close", name="SMA") -> DataFrame:
    '''Calcluates the SMA of Close prices in a dataframe'''
    label = f"{name}({days})"
    df[label] = df[col].rolling(days).mean()
    return df

def subtract(df: DataFrame, col1: str, col2: str,
             name="Difference") -> DataFrame:
    '''
    Returns a new data frame column where every value is
    computed where col1 - col2
    '''
    df[name] = df[col1] - df[col2]
    return df
