'''
This file gets historical stock data by utilizing the yfinance API

Modules used
- pandas
- yfinance

Author: Vikas Katari
Date: 05/03/2025
'''

import yfinance as yf
import pandas as pd
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


def process_data(df: DataFrame) -> None:
    '''
    Adds additional columns with technical indicators based on OHLCV columns

    Args:

    df: DataFrame with OHLCV values

    Returns:

    A modified DataFrame with various technical indicators added
    '''


    '''
    Moving Average 14
    Moving Average 20
    Moving Average 50
    Moving Average 200

    '''




    return
