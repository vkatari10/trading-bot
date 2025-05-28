# Data Processing

This folder contains methods to modify the DataFrames containing historical data from the historical data API to include columns with technical indicators and signals to "BUY" and "SELL".  These modified dataframes will then be used to train the ML models.

# Files

## `data_processing.py`
This file contains methods to get a DataFrame with technical indicators, you can specificy which indicators and relationships between indicators that you would like to define in the `process_data()` method.

## `technicals.py`
This file contains methods to calculate technical indicators which can be called in `data_processing.py` when defining which strategy you would like the ML bot to train and execute on.

## `signals.py`
This file contains relationships that can be defined various indicators such as finding crossovers, or if one is above or below one another.
