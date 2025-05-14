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
import finta as ft


def get_close(df: pd.DataFrame) -> pd.Series:
    '''Gets close data back a Pandas Series'''
    return df['close']

def append_new(df: pd.Series, quote_price: float) -> pd.Series:
    '''
    Appends new price to the 'close' pd.Series obtained from
    get_close()
    '''
    pass


def recompute(df: pd.DataFrame, quote_price: float) -> None:
    '''
    Recomputes the technicals in real time by modifying the selected
    technicals and appending to the bottom of the DataFrame
    '''
    pass

def make_csv(file_name: str, df: pd.DataFrame) -> None:
    '''Takes the runtime df with recomputed technicals and saves it'''
    df.to_csv('runtime_df.csv')
