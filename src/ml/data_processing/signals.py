'''
This file contains methods to add signals to Pandas DataFrames that
contain technical indicators

Author: Vikas Katari
Date: 05/15/2025
'''
import pandas as pd
import numpy as np

import src.ml.json.json_parser as jp

def crossover(df: pd.DataFrame, col1: str, col2: str, persist: int) -> pd.Series:
    '''
    Finds the crossover between two indicators contained in two
    columns of a DataFrame.

    Where col1 crosses above col2 -> 1
    Where col2 crosses below col1 -> -1
    Where no crossing occurs (i.e. equal values) -> 0

    *persist not used, only to fit dispatcher signature
    '''
    crosses = np.where(
        (df[col1].shift(1) < df[col2].shift(1)) &
        (df[col1] > df[col2]),
        1,
        np.where((df[col1].shift(1) > df[col2].shift(1))
                 & (df[col1] < df[col2]), -1, 0))
    return crosses

# TODO: Redo methods below 

def above_sig(df: pd.DataFrame, col1: str, col2: str, persist: int) -> pd.Series:
    """
    Signal logic:
    - 1 if col1 > col2 for `persist` consecutive rows
    - -1 if col1 < col2 (immediate)
    - 0 otherwise
    """
    gt = df[col1] > df[col2]
    lt = df[col1] < df[col2]

    result = np.zeros(len(df), dtype=int)
    above_count = 0

    for i in range(len(df)):
        if lt.iloc[i]:
            result[i] = -1
            above_count = 0  # reset streak
        elif gt.iloc[i]:
            above_count += 1
            if above_count >= persist:
                result[i] = 1
            else:
                result[i] = 0
        else:
            # equal — reset streak, emit 0
            above_count = 0
            result[i] = 0

    return result

def below_sig(df: pd.DataFrame, col1: str, col2: str, persist: int) -> pd.Series:
    """
    Signal logic:
    - 1 if col1 < col2 for `persist` consecutive rows
    - -1 if col1 > col2 (immediate)
    - 0 otherwise
    """
    lt = df[col1] < df[col2]
    gt = df[col1] > df[col2]

    result = np.zeros(len(df), dtype=int)
    below_count = 0

    for i in range(len(df)):
        if gt.iloc[i]:
            result[i] = -1
            below_count = 0  # reset
        elif lt.iloc[i]:
            below_count += 1
            if below_count >= persist:
                result[i] = 1
            else:
                result[i] = 0
        else:
            below_count = 0
            result[i] = 0

    return result

def final_sig(uc: jp.UserConfig, df: pd.DataFrame, start_col: int) -> pd.Series:
    '''
    Calculates the final signal for a given row 

    1. Multiplies relationship value against it defined weight in the config
    2. Take the sum of all relationship values multiplied against their weight
    3. Place that sum as the final signal value

    let sig be the final signal value

    sig > 0 = BUY
    sig < 0 = SELL
    sig = 0 = HOLD

    0 can also be replaced with a threshold/buffer should a user want a higher value than 0 
    to represent buy / sell signals and maintain hold more often
    '''
    # but we need to make sure that based on the weight that the value
    # is a discerete conitinious value integer like -1, 0, 1 we should
    # not be passing raw float values 
    weights = uc.get_all_weights()

    final_sigs = []

    for i in range(len(df)):
        row = df.iloc[i, start_col:]
        weighted_row = row.values * weights
        final_sum = weighted_row.sum()

        sig = 0

        if final_sum > 0:
            sig = 1
        elif final_sum < 0:
            sig = -1
        else:
            sig = 0

        final_sigs.append(sig)
    
    return final_sigs


def above(df: pd.DataFrame, col1: str, col2: str) -> pd.Series:
    '''
    Finds if col1 is above col2, where true the value will be
    1, else 0
    '''
    above = np.where((df[col1] > df[col2]), 1, 0)
    return above


def below(df: pd.DataFrame, col1: str, col2: str) -> pd.Series:
    '''
    Finds if col1 is below col2, where true the value will be
    1, else 0
    '''
    below = np.where((df[col1] < df[col2]), 0, 1)
    return below

def sum_to_sigs(df: pd.DataFrame, start_col: int) -> pd.Series:
    '''
    Takes relationship columns after the indicators have been declared
    and sums up their signs to produce the signal column where >0
    signifies a buy signal, =0 signifies a hold signal, and <0
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
