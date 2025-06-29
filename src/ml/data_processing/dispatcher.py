'''
Dispatch Tables used to condense long if-else 
statements when parsing JSONs DataFrame Contstruction

Author: Vikas Katari 
Date: 06/29/2025
'''

import src.ml.data_processing.signals as sig
import src.ml.data_processing.technicals as te
from pandas import DataFrame, Series

FEATURE_DISPATCH = {
    "SMA": te.sma,
    "EMA": te.ema,
    "delta": te.delta,
    "diff": te.diff
} 

LABEL_DISPATCH = {
    "crossover": sig.crossover,
    "above": sig.above,
    "below": sig.below
}

def dispatch_feature(dispatch: str, df: DataFrame) -> Series:
    pass

def dispatch_label(dispatch: str, df: DataFrame, col1: str, 
                   col2: str) -> Series:
    return LABEL_DISPATCH[dispatch](df, col1, col2)
