'''
File containing methods to train ML models on DataFrames with
historical stock data with all features and labels present

Modules Used
- pandas
- sklearn
- pickle

Author: Vikas Katari
Date: 05/13/2025
'''
import pandas as pd
from pandas import DataFrame
from sklearn.model_selection import train_test_split
from sklearn.ensemble import RandomForestClassifier
from sklearn.metrics import accuracy_score, classification_report
import warnings
import pickle


# ====Pickle Usage====
# filename = xyz
# pickle.dump(model, open(filename, 'wb))
#
# load = pickle.load(open(filename, 'rb'))


def train(df: DataFrame):
    '''
    Trains a model on a given DataFrame
    '''

    '''
    use randomforestclassifier() first
    '''
