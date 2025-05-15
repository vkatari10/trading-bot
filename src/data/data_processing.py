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
import src.data.signals as sg
import src.yfinance_api.yfinance_api as yf

from finta import TA


def process_data(df: pd.DataFrame) -> pd.DataFrame:
    '''
    Takes the YFinance DataFrame and fits it with our own technical
    indicators of choice, and relationships to watch for. We can then
    also add the buy or sell column

    Args:

    df (DataFrame) : DataFrame with OHLCV values

    Returns:

    A modified DataFrame with various technical indicators added
    '''

    # for finta
    if isinstance(df.columns, pd.MultiIndex):
        df.columns = df.columns.get_level_values(0)

        df.columns = df.columns.str.lower()
    '''
    Remember that indexes 0-4 from yfinance API are:
    close
    high
    low
    open
    volume (optional, usually excluded)
    '''

    df.dropna(inplace=True)
    # Finta methods
    df['RSI'] = TA.RSI(df)
    bbands = TA.BBANDS(df)
    df['BB_LOWER'] = bbands['BB_LOWER']
    df['BB_UPPER'] = bbands['BB_UPPER']
    macd = TA.MACD(df)
    df['MACD'] = macd['MACD']
    df['SIGNAL'] = macd['SIGNAL']

    # my own methods
    df = te.sma(df, 10, col='close')  # SMA(10)
    df = te.sma(df, 30, col='close')  # SMA(30)

    df.dropna(inplace=True)

    df = sg.crossover(df, 'SMA(10)', 'SMA(30)', col_name='SMA_CROSS')
    df = sg.crossover(df, 'MACD', 'SIGNAL', col_name='MACD_CROSS')
    df = sg.above(df, 'close', 'BB_UPPER', col_name='BB_SELL')
    df = sg.below(df, 'close', 'BB_LOWER', col_name='BB_BUY')

    df = sg.rsi_signal(df, 'RSI', col_name='RSI_SIG')

    df['signal_final'] = sg.sum_to_sigs(df, 12)

    return df


def get_df(ticker: str) -> pd.DataFrame:
    '''
    Returns the modified dataframe of a stock with
    technicals and signals of the specificed ticker
    '''
    df = yf.get_data(ticker)
    df = process_data(df)
    # export_df(df) uncomment this when we are off notebook
    return df


def export_df(df: pd.DataFrame):
    df.to_csv('src/data/technical_df.csv', index=False)


df = get_df("AAPL")
print(df)
