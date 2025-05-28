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
import json
from typing import Dict, Any, List


# Contains technicals that cannot be computed in real time
# only use for testing ML accuracy with more technicals
from finta import TA


# Python technical indicators
import src.machine_learning.data_processing.technicals as te
import src.api.external.historical_api.yfinance_api as yf # yfinance




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

    df.dropna(inplace=True)
    computing_col = "Close"

    # User defined features
    with open("src/logic/features.json") as f:
        features = json.load(f)

    # Load the features onto the DF
    load_features(df, features)

    df.dropna(inplace=True)
    return df


def load_features(df: pd.DataFrame,
                  features: List[Dict[str, Any]]) -> pd.DataFrame:
    '''
    Loads features onto the given DataFrame based on user
    defined technicals in src/logic/features.json
    '''
    for i in range(len(features)):

        if features[i]['tech'] == "SMA":
            window = features[i]['window']
            label = f"SMA_{window}"
            df[label] = te.sma(df, window)

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

