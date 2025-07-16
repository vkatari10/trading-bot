'''
File containing methods to train ML models on DataFrames with
historical stock data with all features and signal labels present

Author: Vikas Katari
Date: 05/13/2025
'''
# data/eval imports
import pandas as pd
from sklearn.model_selection import train_test_split
from sklearn.metrics import accuracy_score, classification_report
from imblearn.under_sampling import RandomUnderSampler
import pickle

# project files
import src.ml.data_processing.user_df as ud
import src.ml.data_processing.data_processing as dp
import src.ml.json.json_parser as jp


# model imports
from sklearn.ensemble import (
    RandomForestClassifier,
    GradientBoostingClassifier,
    AdaBoostClassifier,
    ExtraTreesClassifier,
)
from sklearn.neighbors import KNeighborsClassifier
from sklearn.linear_model import LogisticRegression, RidgeClassifier, SGDClassifier, PassiveAggressiveClassifier
from sklearn.svm import SVC
from sklearn.tree import DecisionTreeClassifier
from sklearn.neural_network import MLPClassifier

import lightgbm as lgb
import xgboost as xgb

MODEL_DISPATCH = {
    "RandomForrestClassifier": RandomForestClassifier,
    "gradient_boosting": GradientBoostingClassifier,
    "ada_boost": AdaBoostClassifier,
    "extra_trees": ExtraTreesClassifier,
    "knn": KNeighborsClassifier,
    "logistic_regression": LogisticRegression,
    "ridge_classifier": RidgeClassifier,
    "sgd": SGDClassifier,
    "passive_aggressive": PassiveAggressiveClassifier,
    "svm": SVC,
    "decision_tree": DecisionTreeClassifier,
    "mlp": MLPClassifier,
}


def find_stop(df: pd.DataFrame, uc: jp.UserConfig) -> int:
    '''
    Finds Stop index for features
    '''
    return df.columns.get_loc((uc.get_labels()[0]['name']))

# DEPRECATED
def model_training(df: pd.DataFrame, to_col: int,
                   *args) -> RandomForestClassifier:
    '''Training for Scikit learn models'''
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

def scikit_train(ud: ud.UserMLConfig, df: pd.DataFrame) -> None:
    '''Training for Scikit learn models'''

    model = MODEL_DISPATCH[ud.config.get_model_type()](**ud.config.get_hyperparameters())

    X_bal, y_bal = dp.rebalance_df(ud, df)

    X_train, X_test, y_train, y_test = train_test_split(X_bal, y_bal,
                                                        test_size=0.2,
                                                        random_state=42)
    
    model.fit(X_train, y_train)

    y_pred = model.predict(X_test)
    print("Test Accuracy:", accuracy_score(y_test, y_pred))

    with open("models/" + ud.config.get_model_name(), 'wb') as f:
        pickle.dump(model, f)

def lightgbm_train(ud: ud.UserMLConfig, df: pd.DataFrame) -> None:
    '''Training for LightGBM models'''
    
    model = lgb.LGBMClassifier(**ud.config.get_hyperparameters())

    X_bal, y_bal = dp.rebalance_df(ud, df)

    X_train, X_test, y_train, y_test = train_test_split(X_bal, y_bal,
                                                        test_size=0.2,
                                                        random_state=42)
    
    model.fit(X_train, y_train)

    y_pred = model.predict(X_test)
    print("Test Accuracy", accuracy_score(y_test, y_pred))

    model.booster_.save_model("models" + ud.config.get_model_name())

def xgboost_train(ud: ud.UserMLConfig, df: pd.DataFrame) -> None:
    '''Training for XGBoost models'''

    X_bal, y_bal = dp.rebalance_df(ud, df)

    y_bal = y_bal + 1 # adj for XGBoost


    X_train, X_test, y_train, y_test = train_test_split(X_bal, y_bal, test_size
                                                        =0.2, 
                                                        random_state=42)
    # Strip col names for all 
    X_train_np = X_train.values
    X_test_np = X_test.values
    y_train_np = y_train.values
    y_test_np = y_test.values

    model = xgb.XGBClassifier()
    model.fit(X_train_np, y_train_np)

    y_pred = model.predict(X_test_np)

    print("Test Accuracy:", accuracy_score(y_test_np, y_pred))

    model.save_model("models/" + ud.config.get_model_name())