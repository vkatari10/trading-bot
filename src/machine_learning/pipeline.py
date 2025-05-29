'''
machine_learning folder script that allows user the define feature
relationships and call all methods to train and export the ML models.

Author: Vikas katari
Date: 05/28/2025
'''

import pandas as pd
import pickle


# to declare signals to buy/sell
import src.machine_learning.data_processing.signals as sig

# to process dataframes
import src.machine_learning.data_processing.data_processing as dp

# to train models
import src.machine_learning.training.training as train


# ======== USER DEFINED FUNCTIONS ========

training_ticker = "AAPL"


df = dp.get_df(training_ticker) # DO NOT MODIFY

stop = []


def relationships(df: pd.DataFrame) -> pd.DataFrame:
    '''
    Insert the relationships you would like to define
    based on the technicals indicators defined in
    src/logic/features.json

    You can use the signals modules import as .sig
    '''

    # ==== Insert relationships here here ====
    df["SMA_CROSS"] = sig.crossover(df, "SMA_10", "SMA_30")


    # ==== replace index tuple with first relationship defined ====
    index = df.columns.get_loc(('SMA_CROSS', ''))
    stop.append(index)

    # ==== do not modify ====
    df['final_signal'] = sig.sum_to_sigs(df, index)
    return df


# DO NOT MODIFY THE FILE BELOW THIS LINE

# ======= SCRIPT ========

# add user defined relationships
df = relationships(df)

# train model on this dataframe
# we exclude volume for now
model = train.model_training(df, stop[0], 4)

# export model to runtime destination
with open('src/machine_learning/models/decider/model.pkl', 'wb') as f:
    # replace the df with the actual model when its compelte
    pickle.dump(model, f)
