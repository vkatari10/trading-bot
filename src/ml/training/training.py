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
import src.ml.json.json_parser as jp
from sklearn.model_selection import train_test_split
from sklearn.ensemble import RandomForestClassifier
from sklearn.metrics import accuracy_score, classification_report
from imblearn.under_sampling import RandomUnderSampler

# Note we can use different training models

def find_stop(df: pd.DataFrame, uc: jp.UserConfig) -> int:
    '''
    Finds Stop index for features
    '''
    return df.columns.get_loc((uc.get_labels()[0]['name']))


def model_training(df: pd.DataFrame, to_col: int,
                   *args) -> RandomForestClassifier:
    '''
    Method that will train an ML models on all DF contained
    in the JSON config file
    '''
    cols = [i for i in range(to_col)]

    for arg in args:
        del cols[arg]

    # Model
    rf_classifier = RandomForestClassifier(n_estimators=100, random_state=42)

    X = df.iloc[:, cols] # All columns with technical indicators
    y = df.iloc[:, -1] # Just the signal column (label)

    rus = RandomUnderSampler(random_state=42) 
    X_res, y_res = rus.fit_resample(X, y) 

    X_train, X_test, y_train, y_test = train_test_split(X_res, y_res,
                                                        test_size=0.2,
                                                        random_state=42)

    rf_classifier.fit(X_train, y_train)
    return rf_classifier
