# Trading Bot

A platform to develop machine Learning based trading bots that execute orders in real time on predefined technical indicator patterns.

## Requirements

In order to make and view trades you will need an Alpaca API account. The API keys will need to be stored in a `.env` file in the top level directory. You will also need to register for a finnhub account to stream real time data as well to your those associated keys as well. These should also be stored in the `.env` file.<br>

Required packages are listed in the requirements.txt, those can be stored in a venv directory by calling `python -m venv venv` from the top level directory and then installed by calling `pip install -r requirememts.txt`. There is a provided script to activate the virtual environment by calling `source env.sh`. 

## Motivation

The goal of this project is automate trading in real time by using a pre-trained Machine Learning (ML) model that can predict buy and sell signals. By including various financial technical indicators we can test different models and their effectivness.

## Tech Stack

This project includes a mix of Python, Go, and C in the back end to take advantage of the unique benefits each language offers. In a high level sense, this is what each language is used for<br>
- Python: ML training and execution
- Go: Runtime environemnt
- C: Runtime computations

Doing so allows for maximal runtime using Go and C while also taking advantage of the vast ecosystem of ML and Finance related libraries Python offers.

# Limitations

There are several limitaions with this trading bot, however the most signficant are listed below.<br>
- Ingestion of delayed data, up to 15 minutes
- Burn-in period during live execution (~30 minutes)
