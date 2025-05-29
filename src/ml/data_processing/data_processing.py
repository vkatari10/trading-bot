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
import src.ml.data_processing.technicals as te
import src.api.external.historical_api.yfinance_api as yf # yfinance
import src.ml.data_processing.signals as sig


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

    with open("src/logic/signals.json") as f:
        signals = json.load(f)

    # Load the features onto the DF
    df = load_features(df, features)
    df = relationships(df, signals)

    df.dropna(inplace=True)
    return df


def load_features(df: pd.DataFrame,
                  features: List[Dict[str, Any]]) -> pd.DataFrame:
    '''
    Loads user defined technical indicators to determine buy/sell
    signals based on the definitions in src/logic/features.json
    '''
    for i in range(len(features)):

        if features[i]['tech'] == "SMA":
            window = features[i]['window']
            name = features[i]['name']
            df[name] = te.sma(df, window)
        elif features[i]['tech'] == "EMA":
            window = features[i]['window']
            name = features[i]['name']
            df[name] = te.ema(df, window)

    return df

def relationships(df: pd.DataFrame,
                  signals: List[Dict[str, Any]]) -> pd.DataFrame:
    '''
    Loads user defined relationships to determine buy/sell signals
    based on the definitions in src/logic/signals.json
    '''

    # DO NOT INCLUDE RELATIONSHIPS FOR TRAINING PURPOSES
    # REMOVE THIS LATER WHEN WE TEST THE OTHER TRANING
    # PROCESS
    stop_col = signals[0]['name']

    for i in range(len(signals)):

        relationship = signals[i]['sig']
        new_name = signals[i]['name']
        col1 = signals[i]['col1']
        col2 = signals[i]['col2']

        if relationship == "crossover":
            df[new_name] = sig.crossover(df, col1, col2)
        elif relationship == "above":
            df[new_name] = sig.above(df, col1, col2)
        elif relationship == "below":
            df[new_name] = sig.below(df, col1, col2)

    # ==== replace index tuple with first relationship defined ====
    index = df.columns.get_loc((stop_col,''))

    # ==== do not modify ====
    df['final_signal'] = sig.sum_to_sigs(df, index)
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

