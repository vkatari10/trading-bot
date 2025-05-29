'''
This file contains methods to add signals to Pandas DataFrames that
contain technical indicators

Modules Used
- pandas
- numpy

Author: Vikas Kataru
Date: 05/15/2025
'''
import pandas as pd
import numpy as np

def crossover(df: pd.DataFrame, col1: str, col2: str) -> pd.Series:
    '''
    Finds the crossover between two indicators contained in two
    columns of a DataFrame.

    Where col1 crosses above col2 -> 1
    Where col2 crosses below col1 -> -1
    Where no crossing occurs -> 0
    '''
    crosses = np.where(
        (df[col1].shift(1) < df[col2].shift(1)) &
        (df[col1] > df[col2]),
        1,
        np.where((df[col1].shift(1) > df[col2].shift(1))
                 & (df[col1] < df[col2]), -1, 0))
    return crosses


def above(df: pd.DataFrame, col1: str, col2: str,
          col_name: str) -> pd.DataFrame:
    '''
    Finds if col1 is above col2, where true the value will be
    1, else 0
    '''
    df[col_name] = np.where((df[col1] > df[col2]), 1, 0)
    return df


def below(df: pd.DataFrame, col1: str, col2: str,
          col_name: str) -> pd.DataFrame:
    '''
    Finds if col1 is below col2, where true the value will be
    1, else 0
    '''
    df[col_name] = np.where((df[col1] > df[col2]), 0, 1)
    return df

def rsi_signal(df: pd.DataFrame, rsi_col: str, col_name: str,
               bottom=30, top=60) -> pd.DataFrame:
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

def sum_to_sigs(df: pd.DataFrame, start_col: int) -> pd.Series:
    '''
    Takes relationship columns after the indicators have been declared
    and sums up their signs to produce the signal column where >0
    signifies a buy signal, =0 singifies a hold signal, and <0
    signifies a sell signal.

    Args:

    df (pd.DataFrame): the dataframe containing the relationship values
    between technical indicators with values of {-1, 0, 1} to determine
    the relationship
    start_col (int): the column inclusive index that contains the first column
    of relationship information

    Return:

    Series representing the sum of technical relationships to signify
    a buy, sell, or hold action
    '''
    cols = len(df.columns)
    return df.iloc[:, start_col:cols].sum(axis = 1)
