'''
This file contains methods to add signals to Pandas DataFrames that
contain technical indicators

Modules Used
- pandas
'''
import pandas as pd
from pandas import DataFrame

def crossover(df: DataFrame, col1: str, col2: str) -> DataFrame:
    '''
    Finds the crossover between two indicators contained in two
    columns of a DataFrame.

    Where col1 crosses col2 -> 1
    Where col2 crosses col1 -> -1
    Where no cross -> 0
    '''
    pass #  TODO implement

def above(df: DataFrame, col1: str, col2: str) -> DataFrame:
    '''
    Finds if col1 is above col2, where true the value will be
    1, else 0
    '''
    pass #  TODO implement

def below(df: DataFrame, col1: str, col2: str) -> DataFrame:
    '''
    Finds if col1 is below col2, where true the value will be
    1, else 0
    '''
