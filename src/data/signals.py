'''
This file contains methods to add signals to Pandas DataFrames that
contain technical indicators

Modules Used
- pandas
- numpy
'''
import pandas as pd
import numpy as np
from pandas import DataFrame


def crossover(df: DataFrame, col1: str, col2: str,
              col_name: str) -> DataFrame:
    '''
    Finds the crossover between two indicators contained in two
    columns of a DataFrame.

    Where col1 crosses above col2 -> 1
    Where col2 crosses below col1 -> -1
    Where no crossing occurs -> 0
    '''
    df[col_name] = np.where(
        (df[col1].shift(1) < df[col2].shift(1)) &
        (df[col1] > df[col2]),
        1,
        np.where((df[col1].shift(1) > df[col2].shift(1))
                 & (df[col1] < df[col2]), -1, 0))
    return df


def above(df: DataFrame, col1: str, col2: str,
          col_name: str) -> DataFrame:
    '''
    Finds if col1 is above col2, where true the value will be
    1, else 0
    '''
    df[col_name] = np.where((df[col1] > df[col2]), 1, 0)
    return df


def below(df: DataFrame, col1: str, col2: str,
          col_name: str) -> DataFrame:
    '''
    Finds if col1 is below col2, where true the value will be
    1, else 0
    '''
    df[col_name] = np.where((df[col1] > df[col2]), 0, 1)
    return df


def signal(df: DataFrame, col: str, col_name="SIGNAL") -> DataFrame:
    '''
    Determines if the signal is "BUY" or "SELL" if the value is "1"
    or "0", respectivley
    '''
    df[col_name] = np.where(df[col] == 1, "BUY",
                            np.where(df[col] == -1, "SELL", "HOLD")
)
    return df

def rsi_signal(df: DataFrame, rsi_col: str, col_name: str,
               bottom=30, top=60) -> DataFrame:
    '''
    Finds the crossover between two indicators contained in two
    columns of a DataFrame.

    Where col1 crosses above col2 -> 1
    Where col2 crosses below col1 -> -1
    Where no crossing occurs -> 0
    '''
    df[col_name] = np.where(df[rsi_col] < 30, 1,
                            np.where(df[rsi_col] > 70, -1, 0))

    return df
