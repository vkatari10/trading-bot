'''
This file contains the metohd to get Panda DataFrames using the YFinance Api.

Modules
- yfinance
- pandas

Author: Vikas Katari
Date: 05/12/2025
'''

import yfinance as yf
from pandas import DataFrame

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
