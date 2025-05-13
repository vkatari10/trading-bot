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
import pickle
from sklearn.model_selection import train_test_split
from sklearn.ensemble import RandomForestClassifier
from sklearn.metrics import accuracy_score, classification_report
from imblearn.under_sampling import RandomUnderSampler
import warnings

def train(df: DataFrame) -> RandomForestClassifier:
    X = df.iloc[:, [0, 4]] # All columns except the signal (feature)
    y = df.iloc[:, 10] # Just the signal column (label)

    rus = RandomUnderSampler(random_state=42)
    X_res, y_res = rus.fit_resample(X, y)

    X_train, X_test, y_train, y_test = train_test_split(X_res, y_res, test_size=0.2, random_state=42)

    # Initialize RandomForestClassifier
    rf_classifier = RandomForestClassifier(n_estimators=100, random_state=42)

    # Fit the classifier to the training data
    rf_classifier.fit(X_train, y_train)

    # Make predictions
    y_pred = rf_classifier.predict(X_test)

    return rf_classifier
