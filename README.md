# Trading Bot

Python based trading bot that utilizes Alpaca API to place, order, and manage trades.

## Structure

Back end is done entirely in Python.

## Requirements

In order to make and view trades you will need an Alpaca API account. The API keys will need to be stored in a `.env` file in the top level directory. You will also need to register for a finnhub account as well to your those associated keys as well. These should also be stored in the `.env` file.<br>

Required packages are listed in the requirements.txt, those can be stored in a venv directory by calling `python -m venv venv` from the top level directory. There is a provided script to activate the virtual environment by calling `source env.sh`. You can also optionally use pypy3 to run this code instead of the default python3 intreperter.
