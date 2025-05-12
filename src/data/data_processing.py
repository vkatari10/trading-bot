'''
This file gets historical stock data by utilizing the yfinance API.
Also adds technical indicators and buy/sell signals.

Modules used
- pandas
- yfinance

Author: Vikas Katari
Date: 05/03/2025
'''

import pandas as pd
import backtrader as bt
import numpy as np
import src.data.technicals as te
import src.yfinance_api.yfinance_api as yf
from pandas import DataFrame

def process_data(df: DataFrame) -> DataFrame:
    '''
    Adds additional columns with technical indicators based on OHLCV columns

    Args:

    df: DataFrame with OHLCV values

    Returns:

    A modified DataFrame with various technical indicators added
    '''

    '''
    SMA 10
    SMA 30
    RSI
    MACD -> EMA 12, 26, EMA of EMA_12-EMA_26
    '''

    df.dropna(inplace=True)

    df = te.sma(df, 10) # SMA 10
    df = te.sma(df, 30) # SMA 30
    df = te.ema(df, 12) # EMA 12
    df = te.ema(df, 26) # EMA 26
    df = te.subtract(df, "EMA(12)", "EMA(26)", "EMA (12-26)")
    df = te.ema(df=df, days=9, col="EMA(12-26)", name="Signal Line")


    df.dropna(inplace=True)

    return df



def add_signals(df: DataFrame) -> DataFrame:
    df["Signal"] = np.where((df['SMA(50)'].shift(1) > df['SMA(200)'].shift(1)) & (df['SMA(50)'] < df['SMA(200)']), -1, np.where((df['SMA(50)'].shift(1) < df['SMA(200)'].shift(1)) & (df['SMA(50)'] > df['SMA(200)']),1, 0))

    df["Signal"] = np.where((df['SMA(20)'].shift(1) > df['SMA(50)'].shift(1)) & (df['SMA(20)'] <
                                                             df['SMA(50)']),
        -1, np.where((df['SMA(20)'].shift(1) < df['SMA(50)'].shift(1)) &
                     (df['SMA(20)'] > df['SMA(50)']), 1, 0)
    )

    return df

def get_df(ticker: str) -> DataFrame:
    '''
    Returns the modified dataframe of a stock with
    technicals and signals
    '''
    df = yf.get_data(ticker)
    df = process_data(df)
    df = add_signals(df)
    return df
