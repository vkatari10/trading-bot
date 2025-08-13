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
from typing import Dict, Any, List


# Python technical indicators
import src.api.external.historical_api.yfinance_api as yf
import src.ml.data_processing.signals as sig
import src.ml.data_processing.dispatcher as dp
import src.ml.data_processing.technicals as te

# User Config JSON parser
import src.ml.json.json_parser as jp

# User ML settings
import src.ml.data_processing.user_df as ud


def process_data(df: pd.DataFrame, ud: ud.UserMLConfig) -> pd.DataFrame:
    '''
    Modifies a YFinance DataFrame to specifications used by the jp.UserConfig
    features and label configuration

    Args:

    df (pd.DataFrame): DataFrame from YFinance (no multiindex cols)
    config (jp.UserConfig): UserConfig object containing user features and labels

    Returns: 

    A pd.DataFrame with all user defined features and labels
    '''
    df.columns = [col.lower() for col in df.columns] # rename for TA-lib
    df.dropna(inplace=True)

    # put featurs on the training dataframe
    df = load_features(df, ud.get_features())
    df = put_relationships(ud, df, ud.get_labels())

    # print("Number of things that are not hold")
    # print(len(df[df['final_signal'] != 0]))
    df.dropna(inplace=True)
    return df


def load_features(df: pd.DataFrame,
                  features: List[Dict[str, Any]]) -> pd.DataFrame:
    '''
    Loads user defined technical indicators (features) on to the 
    training dataframe

    Args:

    df (pd.DataFrame): DataFrame from YFinance (no multiindex cols)
    features (List[Dict[str, Any]]): List of JSON objects from the user config 
    file 

    Returns: 

    The modified dataframe with the user defined features
    '''
    for i in range(len(features)):

        tech = features[i]['tech']

        if tech == 'delta' or tech == 'diff':
            df[features[i]['name']] = dp.dispatch_feature(features[i]['tech'],
                                                          df, features[i])
        else:
            df[features[i]['name']] = te.put_technical(features[i]['tech'], 
                                                       df, features[i]['args'])

    return df


def put_relationships(uc: jp.UserConfig, df: pd.DataFrame, 
                      labels: List[Dict[str, Any]]) -> pd.DataFrame:
    """Puts user defined labelling logic onto the training DataFrame"""
    for label in labels:
        df[label['name']] = dp.dispatch_signal(label['sig'],
                                               df,
                                               label['col1'],
                                               label['col2'],
                                               label.get('persist'))

    # put final signal values based on relationship values / weights
    df['final_sig'] = sig.final_sig(uc, df, df.columns.get_loc(labels[0]['name']))  
    return df

def get_df(uc: jp.UserConfig, concat=True) -> pd.DataFrame:
    '''
    Returns dataframe(s) with all user defined features and labels,

    Args:

    uc (jp.UserConfig): Object representing JSON config file and wanted 
    features
    cocnat (bool): True if all dataframes for all tickers should be returned
    as one, else False to recieve a list of the DataFrames in the same order
    as configured in the JSON config file

    Returns:

    A concated DataFrame of all tickers if concat is true from the config file, 
    else a list of DataFrames of each ticker if concat is false
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
    
def get_single_df(ud: ud.UserMLConfig, ticker: str) -> pd.DataFrame:
    '''
    Returns a single dataframe with all user defined features and labels,

    Args:

    uc (jp.UserConfig): Object representing JSON config file and wanted 
    features
    ticker (str): Valid stock ticker 
    
    Returns:

    A single dataframe with the user defined features and labels from the JSON 
    config file 
    '''
    df = yf.get_data(ticker, ud.get_model_training_interval(), ud.get_model_training_timeframe())
    df = process_data(df, ud)
    return df

def rebalance_df(ud: ud.UserMLConfig, 
                 df: pd.DataFrame) -> List[pd.DataFrame | pd.Series]:
    '''Rebalance dataframe by features and labelling logic, useful for all models'''
    from imblearn.under_sampling import RandomUnderSampler

    cols = [i for i in range(len(df.columns) - 1)]

    X = df.iloc[:, cols]
    y = df.iloc[:, -1] # TODO accept multiple back label columns

    rus = RandomUnderSampler(random_state=42)
    X_res, y_res = rus.fit_resample(X, y)

    return [X_res, y_res]

def normalize_df(ud: ud.UserMLConfig, df: pd.DataFrame) -> pd.DataFrame:
    '''Normalizes the data in a DF, not needed for tree based models'''
    pass # TODO implement
