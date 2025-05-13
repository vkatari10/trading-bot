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

    # Technical Indicators
    df = te.sma(df, 10)  # SMA(10)
    df = te.sma(df, 30)  # SMA(30)
    df = te.ema(df, 12)  # EMA(12)
    df = te.ema(df, 26)  # EMA(26)
    df = te.subtract(df, "EMA(12)", "EMA(26)", "EMA(12-26)")  # MACD
    df = te.ema(df=df, days=9, col="EMA(12-26)", name="Signal Line")
    # MACD

    # Crossovers
    df = sg.crossover(df, "SMA(10)", "SMA(30)", "SMA Crossover")
    df = sg.crossover(df, "EMA(12-26)", "Signal Line(9)", "MACD")

    # BUY, SELL, or HOLD Signals
    df = sg.signal(df, "SMA Crossover", name="Signal")
    df = sg.signal(df, "MACD", name="Signal")

    df.dropna(inplace=True)

    return df


def get_df(ticker: str) -> DataFrame:
    '''
    Returns the modified dataframe of a stock with
    technicals and signals of the specificed ticker
    '''
    df = yf.get_data(ticker)
    df = process_data(df)
    return df
