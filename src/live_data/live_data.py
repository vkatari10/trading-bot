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

'''
export the dataframe as a CSV, we can then load it into here
and then keep appending the new data

take the dataframe, take the close and convert it to a series

-> take the series
-> append the new data from finnhub
-> recompute using finta
-> get new techicals
-> reappend back to the original dataframe
-> make model predict on the last added row if buy/sell/hold

Execution:
-> bring the dataframe into the driver file
-> get quote update from finnhub
->
'''
import pandas as pd
import numpy as np
from finta import TA
import src.data.technicals as te

def append_new(df: pd.DataFrame, current: float,
               high: float, low: float, _open: float) -> pd.DataFrame:
    '''
    Appends new price to the 'close' pd.Series obtained from
    get_close() along with the recomputated technicals
    '''
    new_row = pd.DataFrame({
        'close': [current],
        'high': [high],
        'low': [low],
        'open': [_open],
        'RSI': [0],
        'BB_UPPER': [0],
        'BB_LOWER': [0],
        'MACD': [0],
        'SIGNAL': [0],
        'SMA(10)': [0],
        'SMA(30)': [0],
        'ACTION': [0]
    })

    new_row.index = [len(df)]

    _new = pd.concat([df, new_row], axis=0)

    rsi = TA.RSI(df)
    bbands = TA.BBANDS(df)
    macd = TA.MACD(df)
    sma30 = te.sma(df, days=30, col='close')
    sma10 = te.sma(df, days=10, col='close')


    _new.iat[-1, 4] = rsi.iat[-1]
    _new.iat[-1, 5] = bbands['BB_LOWER'].iat[-1]
    _new.iat[-1, 6] = bbands['BB_UPPER'].iat[-1]
    _new.iat[-1, 7] = macd['MACD'].iat[-1]
    _new.iat[-1, 8] = macd['SIGNAL'].iat[-1]
    _new.iat[-1, 9] = sma10.iat[-1, 0]
    _new.iat[-1, 10] = sma30.iat[-1, 0]
    _new.iat[-1, 11] = 0

    return _new


def make_csv(file_name: str, df: pd.DataFrame) -> None:
    '''Takes the runtime df with recomputed technicals and saves it'''
    df.to_csv(file_name)
