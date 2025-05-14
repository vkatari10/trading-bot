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
from pandas import DataFrame
from finta import TA

def process_data(df: DataFrame) -> DataFrame:
    '''
    Takes the YFinance DataFrame and fits it with our own technical
    indicators of choice, and relationships to watch for. We can then
    also add the buy or sell column

    Args:

    df (DataFrame) : DataFrame with OHLCV values

    Returns:

    A modified DataFrame with various technical indicators added
    '''
    df.dropna(inplace=True)
    ''' MY CODE VVVVVVVVVVVVVVVVVVVVVVVVVVV
    # Technical Indicators

    df = te.ema(df, 12)  # EMA(12)
    df = te.ema(df, 26)  # EMA(26)
    df = te.subtract(df, "EMA(12)", "EMA(26)", "EMA(12-26)")  # MACD
    df = te.ema(df=df, days=9, col="EMA(12-26)", name="Signal Line")
    # MACD

    # Crossovers
    # df = sg.crossover(df, "SMA(10)", "SMA(30)", "SMA Crossover")
    df = sg.crossover(df, "EMA(12-26)", "Signal Line(9)", "MACD")

    # BUY, SELL, or HOLD Signals
    # df = sg.signal(df, "SMA Crossover", name="Signal")# This doesnt work because the one below overwrites all the signals
    df = sg.signal(df, "MACD", col_name="Signal")

    '''

    # Fix multi-index columns if present
    if isinstance(df.columns, pd.MultiIndex):
        df.columns = df.columns.get_level_values(0)

        # Rename to lowercase for btalib
        df.columns = df.columns.str.lower()

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


    return df


def get_df(ticker: str) -> DataFrame:
    '''
    Returns the modified dataframe of a stock with
    technicals and signals of the specificed ticker
    '''
    df = yf.get_data(ticker)
    df = process_data(df)
    # export_df(df) uncomment this when we are off notebook
    return df

def export_df(df: DataFrame):
    df.to_csv('technical_df.csv', index=False)
