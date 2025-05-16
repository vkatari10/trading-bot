# Trading Bot

Python based ML-based trading bot that can buy, sell, execute orders.

## Requirements

In order to make and view trades you will need an Alpaca API account. The API keys will need to be stored in a `.env` file in the top level directory. You will also need to register for a finnhub account as well to your those associated keys as well. These should also be stored in the `.env` file.<br>

Required packages are listed in the requirements.txt, those can be stored in a venv directory by calling `python -m venv venv` from the top level directory. There is a provided script to activate the virtual environment by calling `source env.sh`. You can also optionally use pypy3 to run this code instead of the default python3 intreperter.<br>

## Motivation

Trading in real time is hard to predict and make split second decisions. By handing over this duty to trained machine learning models this can quickly resolve trades while also stripping the emotional aspects of trading. This tool is pratical as a way to mangage trading in the background as it is very hard to activley manage trades in real life all the time.<br> 

## Tech Stack

All of the current code right now to load, process financial data and train machine learning models and run them in a live environment is all done in python. Future plans include to use websockets using `Go` for API calls and computing technical indicators on historical data using `C++` as a way to reduce runtime. Beyond this repo includes a future project to add a dashboard to moinitor the trading bot in its runtime environment that utliizes `Next.js` and `Typescript`.

