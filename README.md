# Trading Bot

Python based ML-based trading bot that can buy, sell, execute orders, with runtime optimizations in C++.

## Requirements

In order to make and view trades you will need an Alpaca API account. The API keys will need to be stored in a `.env` file in the top level directory. You will also need to register for a finnhub account as well to your those associated keys as well. These should also be stored in the `.env` file.<br>

Required packages are listed in the requirements.txt, those can be stored in a venv directory by calling `python -m venv venv` from the top level directory and then installed by calling `pip install -r requirememts.txt`. There is a provided script to activate the virtual environment by calling `source env.sh`. 

## Motivation

Trading in real time is hard to predict and make split second decisions. By handing over this duty to trained machine learning models this can quickly resolve trades while also stripping the emotional aspects of trading. This tool is pratical as a way to mangage trading in the background as it is very hard to activley manage trades in real life all the time.<br> 

## Tech Stack

The majority of this project is done in Python thanks to its massive ML library ecosystem, however some optimziations were done in other lanaguages to increase the performance of critical runtime elements such as:<br>
- C++ - used to handle the recomputation of technical indicators at runtime
- Go - to be implemented to help speed up API calls by using websockets<br>

Doing so helps the trading bot run far faster in real time which removes some of the overhead with Python, especially since we can take advantage of the fact that these two additional languages are compiled.

# Limitations

There are several limitaions with this trading bot, however the most signficant are<br>
- Ingestion of delayed data
- Burn-in period when running in real time (~30 minutes)

# Version History 

## v0.1.0-alpha
- MVP proof of concept with no regard to performance.
- Prototyped fully in Python.
- Used Pandas directly to set up the data pipeline.
- No server implementation.

## v0.2.0-alpha (in progress)
- Improve performance and remove reliance on the `finta` library
- Replaces Pandas to use raw NumPy arrays.
- Integrated `pybind11` and C++ code to improve performance of technical computations.
- Still no server implmentation.



