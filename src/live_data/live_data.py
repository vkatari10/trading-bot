'''
This file contains methods to recompute technicals in real time by
appending the close data to insert new prices

Modules Used
- pandas
- numpy
- finta

Author: Vikas Katari
Date: 05/13/25
'''
import pandas as pd
import numpy as np
from finta import TA
import src.data.technicals as te
from typing import List
import datetime

import example

example.print()


def initialize_row(df: pd.DataFrame, current: float, high: float,
                   low: float, _open:float) -> pd.DataFrame:
    '''
    Insert a new row in to the dataframe with the current, high,
    low and open values to recompute technical indicators with
    the newest updated price

    Args:

    df (pd.DataFrame): dataframe that contains a currrent, high
    low, _open values from index 0-3 inclusive, default from
    the yfinance API
    current (float): the current price column
    high (float): the high price column
    low (float): the low price column
    _open (float): the open price column

    ReturnL

    A dataframe contianing the current, high, low, _open
    values and technical indicators. Sets all others
    values in all other columns to 0.
    '''
    cols_list = df.columns

    init_data = dict()

    # intialize new row
    init_data[cols_list[0]] = [current]
    init_data[cols_list[1]] = [high]
    init_data[cols_list[2]] = [low]
    init_data[cols_list[3]] = [_open]

    for i in range(5, len(cols_list)):
        init_data[cols_list[i]] = [float(0.0)]


    new_row = pd.DataFrame(init_data)
    new_row.index = datetime.datetime.now()

    _new = pd.concat([df, new_row], axis=0)

    return _new


def append_new(df: pd.DataFrame, current: float,
               high: float, low: float, _open: float) -> pd.DataFrame:
    '''
    Appends new price to the 'close' pd.Series obtained from
    get_close() along with the recomputated technicals
    '''

    # insert bottom row into DF
    _new = initialize_row(df, current, high, low, _open)

    # recompute techncials for the last rows
    rsi = TA.RSI(_new)
    bbands = TA.BBANDS(_new)
    macd = TA.MACD(_new)
    sma30 = te.sma(_new, days=30, col='close')
    sma10 = te.sma(_new, days=10, col='close')

    # manually inesrt the technical recoputation into the last row
    _new.iat[-1, 5] = rsi.iat[-1]
    _new.iat[-1, 6] = bbands['BB_LOWER'].iat[-1]
    _new.iat[-1, 7] = bbands['BB_UPPER'].iat[-1]
    _new.iat[-1, 8] = macd['MACD'].iat[-1]
    _new.iat[-1, 9] = macd['SIGNAL'].iat[-1]
    _new.iat[-1, 10] = sma10.iat[-1, 0]
    _new.iat[-1, 11] = sma30.iat[-1, 0]


    _new.iat[-1, -1] = 0 # Leave 0 to let ML predict action

    return _new


def make_csv(file_name: str, df: pd.DataFrame) -> None:
    '''Takes the runtime df with recomputed technicals and saves it'''
    df.to_csv(file_name)
