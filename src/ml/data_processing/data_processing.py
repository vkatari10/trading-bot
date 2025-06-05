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

    # print(features)
    # print(signals)

    # put featurs on the training dataframe
    df = OHCLV_diffs(df)
    df = load_features(df, features)
    df = relationships(df, signals)

    # print("Number of things that are not hold")
    # print(len(df[df['final_signal'] != 0]))

    df.dropna(inplace=True)
    return df


def load_features(df: pd.DataFrame,
                  features: List[Dict[str, Any]]) -> pd.DataFrame:
    '''
    Loads user defined technical indicators to determine buy/sell
    signals based on the definitions in src/logic/features.json
    '''
    for i in range(len(features)):

        # Guranteed for each object
        name = features[i]['name']
        tech = features[i]['tech']

        if tech == "SMA":
            window = features[i]['window']
            df[name] = te.sma(df, window)
        elif tech == "EMA":
            window = features[i]['window']
            df[name] = te.ema(df, window)
        elif tech in ("delta", "diff"):
            col1 = features[i]["col1"]
            col2 = features[i]["col2"]
            df[name] = handle_relations(df, tech, col1, col2)

    return df


def handle_relations(df: pd.DataFrame, tech: str, col1: str, 
                     col2: str) -> pd.Series:
    '''
    Handles the user features.json file when the user delcares 
    an object with an "tech" value of "delta" or "diff"
    '''

    result = None

    if not col2 and tech == "delta":
        result = te.delta(df, col1)
    elif col1 and col2 and tech == "delta":
        result = te.delta_diff(df, col1, col2)
    elif col1 and col2 and tech == "diff":
        result = te.diff(df, col1, col2)

    return result


def relationships(df: pd.DataFrame,
                  signals: List[Dict[str, Any]]) -> pd.DataFrame:
    '''
    Loads user defined relationships to determine buy/sell signals
    based on the definitions in src/logic/signals.json
    '''

    # DO NOT INCLUDE RELATIONSHIPS FOR TRAINING PURPOSES
    # REMOVE THIS LATER WHEN WE TEST THE OTHER TRANING
    # PROCESS (stop in training.py excludes these)
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


def OHCLV_diffs(df: pd.DataFrame) -> pd.DataFrame:
    '''Puts the difference cols of the OHCLV data from Yfinance'''
    yf_cols = ['Close', 'High', 'Low', 'Open', 'Volume']

    for col in yf_cols:
        col_name = col + "_delta"
        df[col_name] = te.delta(df, col)

    return df


def get_df(ticker: str) -> pd.DataFrame:
    '''
    Returns the modified dataframe of a stock with
    technicals and signals of the specificed ticker
    '''
    df = yf.get_data(ticker)
    df = process_data(df)
    return df