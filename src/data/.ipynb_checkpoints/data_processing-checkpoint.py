'''
This file gets historical stock data by utilizing the yfinance API.
Also adds technical indicators and buy/sell signals.

Modules used
- pandas
- yfinance

Author: Vikas Katari
Date: 05/03/2025
'''

import yfinance as yf
import pandas as pd
import numpy as np
from pandas import DataFrame  # To declare return type


def get_data(ticker: str) -> DataFrame:
    '''
    Returns all historical data given an ticker as a Pandas DataFrame

    Args:

    ticker: the ticker of the asset

    Raises:

    AttributeError: if the ticker does not exist

    Returns:

    A Pandas DataFrame with OHLCV as columns from the stocks inception date
    '''
    df = yf.download(tickers=ticker)
    return df


def process_data(df: DataFrame) -> DataFrame:
    '''
    Adds additional columns with technical indicators based on OHLCV columns

    Args:

    df: DataFrame with OHLCV values

    Returns:

    A modified DataFrame with various technical indicators added
    '''

    '''
    CURRENT TECHNICALS ADDED:

    Moving Average 20
    Moving Average 50
    Moving Average 200

    FUTURE TO ADD:

    Exponential Moving Average
    RSI
    MACD
    '''

    df['SMA(20)'] = df.Close.rolling(20).mean()
    df['SMA(50)'] = df.Close.rolling(50).mean()
    df['SMA(200)'] = df.Close.rolling(200).mean()


    return df


def add_signals(df: DataFrame) -> DataFrame:
    df["Signal"] = np.where((df['SMA(50)'].shift(1) > df['SMA(200)'].shift(1)) & (df['SMA(50)'] < df['SMA(200)']), -1, np.where((df['SMA(50)'].shift(1) < df['SMA(200)'].shift(1)) & (df['SMA(50)'] > df['SMA(200)']),1, 0))

    df["Signal"] = np.where((df['SMA(20)'].shift(1) > df['SMA(50)'].shift(1)) & (df['SMA(20)'] <
                                                             df['SMA(50)']),
        -1, np.where((df['SMA(20)'].shift(1) < df['SMA(50)'].shift(1)) &
                     (df['SMA(20)'] > df['SMA(50)']), 1, 0)
    )

    return df
