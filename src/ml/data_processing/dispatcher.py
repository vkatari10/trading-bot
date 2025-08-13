'''
Dispatch Tables used to condense long if-else 
statements when parsing JSONs DataFrame Contstruction

Author: Vikas Katari 
Date: 06/29/2025
'''

import src.ml.data_processing.signals as sig
import src.ml.data_processing.technicals as te
import pandas as pd
from typing import Dict, Any


FEATURE_DISPATCH = {
    "delta": te.delta,
    "diff": te.diff
} 

# DEPRECATED
LABEL_DISPATCH = {
    "crossover": sig.crossover,
    "above": sig.above,
    "below": sig.below
}


RELATIONSHIP_DISPATCH = {
    "crossover": sig.crossover,
    "above": sig.above_sig,
    "below": sig.below_sig
}


def dispatch_feature(dispatch: str, df: pd.DataFrame, 
                     json: Dict[str, Any]) -> pd.Series:
    '''Dispatcher for technical indicators (features)'''
    return FEATURE_DISPATCH[dispatch](df, json, col="close")

# DEPRECATED
def dispatch_label(dispatch: str, df: pd.DataFrame, col1: str, 
                   col2: str) -> pd.Series:
    '''Dispatcher for labeling logic from technical indicators (features)'''
    return LABEL_DISPATCH[dispatch](df, col1, col2)

def dispatch_signal(dispatch: str, df: pd.DataFrame, col1: str, 
                    col2: str, persist: int) -> pd.Series:
    '''v2 dispatcher for label logic/signal from user configs'''
    return RELATIONSHIP_DISPATCH[dispatch](df, col1, col2, persist)

