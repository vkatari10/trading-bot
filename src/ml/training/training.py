'''
File containing methods to train ML models on DataFrames with
historical stock data with all features and signal labels present

Modules Used
- pandas
- sklearn
- pickle

Author: Vikas Katari
Date: 05/13/2025
'''
import pandas as pd
import pickle
from sklearn.model_selection import train_test_split
from sklearn.ensemble import RandomForestClassifier
from sklearn.metrics import accuracy_score, classification_report
from imblearn.under_sampling import RandomUnderSampler


# Note we can use different training models


def model_training(df: pd.DataFrame, to_col: int,
                   *args) -> RandomForestClassifier:
    '''
    Method to train models given the features (as historical OHLC)
    data that finta expects plus technical indicators.

    This method expects that the label signaling to "buy" (>0),
    "sell" (<0), or hold (=0) are all contained in the last
    column of the data frame, i.e df.iloc[:, -1]

    Args:

    df (pandas.DataFrame): df containing the OHLC and other
    technical indicators added by the user
    to_col (int): the column containing the last technical
    indicator (we do not incldue the relationships we
    defined for since we need to ml model to figure those
    relationships as well -- this should be the actual number
    of the column -- not the index.
    *args (int): columns to exclude before the relationships
    are defined, usually volume (4)
    '''
    cols = [i for i in range(to_col)]

    for arg in args:
        del cols[arg]

    X = df.iloc[:, cols] # All columns with technical indicators
    y = df.iloc[:, -1] # Just the signal column (label)

    rus = RandomUnderSampler(random_state=42) 
    X_res, y_res = rus.fit_resample(X, y) 

    X_train, X_test, y_train, y_test = train_test_split(X_res, y_res,
                                                        test_size=0.2,
                                                        random_state=42)

    # Model
    rf_classifier = RandomForestClassifier(n_estimators=100, random_state=42)

    rf_classifier.fit(X_train, y_train)

   

    return rf_classifier
