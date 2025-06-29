'''
This file gets historical stock data by utilizing the yfinance API.
Also adds technical indicators and buy/sell signals.

Modules used
- pandas
- yfinance
- os/dotenv

Author: Vikas Katari
Date: 05/03/2025
'''
import pandas as pd
from typing import Dict, Any, List


# Python technical indicators
import src.ml.data_processing.technicals as te
import src.api.external.historical_api.yfinance_api as yf # yfinance
import src.ml.data_processing.signals as sig
import src.ml.data_processing.dispatcher as dp

# User Config JSON parser
import src.ml.json.json_parser as jp


def process_data(df: pd.DataFrame, config: jp.UserConfig) -> pd.DataFrame:
    '''
    Modifies a YFinance DataFrame to specifications used by the jp.UserConfig
    features and label configuration

    Args:

    df (pd.DataFrame): DataFrame from YFinance (with Multiindex Cols)
    config (jp.UserConfig): UserConfig object containing user features and labels

    Returns: 

    A pd.DataFrame with all user defined features and labels
    '''
    df.dropna(inplace=True)

    # put featurs on the training dataframe
    df = OHCLV_diffs(df)
    df = load_features(df, config.get_features())
    df = relationships(df, config.get_labels())

    # print("Number of things that are not hold")
    # print(len(df[df['final_signal'] != 0]))
    df.dropna(inplace=True)
    return df


# TODO make these dispatch tables somehow
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

    Loads user defined 
    '''
    # DO NOT INCLUDE RELATIONSHIPS FOR TRAINING PURPOSES
    # REMOVE THIS LATER WHEN WE TEST THE OTHER TRANING
    # PROCESS (stop in training.py excludes these)
    stop_col = signals[0]['name']

    for i in range(len(signals)): # Call Dispatcher for label logic methods 
        df[signals[i]['name']] = dp.dispatch_label(signals[i]['sig'], df, 
                                                   signals[i]['col1'], 
                                                   signals[i]['col2'])

    # ==== replace index tuple with first relationship defined ====
    index = df.columns.get_loc((stop_col))

    # put the final sum of signals to indicate buy/sell
    df['final_signal'] = sig.sum_to_sigs(df, index)
    return df


def OHCLV_diffs(df: pd.DataFrame) -> pd.DataFrame:
    '''Puts the difference cols of the OHCLV data from Yfinance'''
    yf_cols = ['Close', 'High', 'Low', 'Open', 'Volume']

    for col in yf_cols:
        col_name = col + "_delta"
        df[col_name] = te.delta(df, col)

    return df


def get_df(uc: jp.UserConfig, concat=True) -> pd.DataFrame:
    '''
    Returns a dataframe with all user defined features and labels,
    for multiple dataframes that are concatenated and the then returned,
    else if concat is False then they are returned in array in the same order
    as the JSON config delcares on 'train_tickers'
    '''

    training_df = []
    training_stocks = uc.get_training_stocks() # stocks to train on

    if len(training_stocks) == 0:
        print(f"Must have >= 1 training stock declared in JSON config")

    for stock in training_stocks:   
        # download DF
        df = yf.get_data(stock, uc.get_model_training_interval(), 
                        uc.get_model_training_timeframe())

        # put user config data ontop
        df = process_data(df, uc) 

        training_df.append(df)
        
    if concat:
        final_df = pd.concat(training_df, ignore_index=False)
        return final_df
    else:
        return training_df
    
def get_single_df(uc: jp.UserConfig, ticker: str, period, timeframe) -> pd.DataFrame:
    '''
    Returns a single Dataframe given a ticker with all user defined features and labels
    '''
    df = yf.get_data(ticker, period, timeframe)
    df = process_data(df, uc)
    return df
