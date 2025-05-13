'''
File containing methods to train ML models on DataFrames with
historical stock data with all features and labels present

Modules Used
- pandas
- sklearn

Author: Vikas Katari
Date: 05/13/2025
'''
import pandas as pd
from pandas import DataFrame
from sklearn.model_selection import train_test_split
from sklearn.ensemble import RandomForestClassifier
from sklearn.metrics import accuracy_score, classification_report
import warnings
