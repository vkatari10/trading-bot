'''
Dispatch Tables used to condense long if-else 
statements when parsing JSONs DataFrame Contstruction

Author: Vikas Katari 
Date: 06/29/2025
'''

import src.ml.data_processing.signals as sig
import src.ml.data_processing.technicals as te
from pandas import DataFrame, Series
from typing import Dict, Any


# delta, diff objects are handled seperately since they 
# are special cases

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

def dispatch_feature(dispatch: str, df: DataFrame, json: Dict[str, Any]) -> Series:
    return FEATURE_DISPATCH[dispatch](df, json, col="Close")

def dispatch_label(dispatch: str, df: DataFrame, col1: str, 
                   col2: str) -> Series:
    return LABEL_DISPATCH[dispatch](df, col1, col2)
