# Trading Bot

A platform to develop machine Learning based trading bots that execute orders in real time on predefined technical indicator patterns.

## Goals

- Complete Machine Learning Pipeline
- Real-time trading engine
- Modular Structure to easily replace trading logic

## Motivation

The goal of this project is automate trading in real time by using a pre-trained Machine Learning (ML) model that can predict buy and sell signals. By including various financial technical indicators at training time we can test different models and their effectivness. 

## Architecture

The following diagram highlights the modular back end architecture.

![Architecture](docs/images/Architecture_diagram2.svg)

### Legend

#### Box Colors
- Golang (Cyan)
- Python (Blue)
- C (Gray)

#### Box Borders 
- ML Pipeline (Pink)
- APIs (Red)
- Runtime logic (Orange)

## Tech Stack

This project includes a mix of Python, Go, and C in the back end to take advantage of the unique benefits each language offers. In a high level sense, this is what each language is used for<br>
- Python: ML training and execution
- Go: Runtime environemnt
- C: Runtime computations

Doing so allows for maximal runtime using Go and C while also taking advantage of the vast ecosystem of ML and Finance related libraries Python offers.
    
## Future Additions
- General 
  - More built-in technical indicators at training and run time
  - Make extnernal broker/data streaming API structure to be modular
- Machine Learning Pipeline
  - Additions to train on multiple stocks
  - Allow users to tune hyperparameters
  - Allow uesrs to choose ML model (KNN, Linear Regression, etc.)
- APIs
  - Find a higher frequency market data streamer
  - Convert ML API to use `FastAPI` instead of `Flask`

## Limitations

There are several limitaions with this trading bot, however the most signficant are listed below.<br>
- Ingestion of delayed data, uo to 1-15 minutes
- Burn-in period during live execution (~30 minutes) (Planned)

## Requirements

In order to make and view trades you will need an Alpaca API account. The API keys will need to be stored in a `.env` file in the top level directory. You will also need to register for a finnhub account to stream real time data as well to your those associated keys as well. These should also be stored in the `.env` file.<br>

Required packages are listed in the requirements.txt, those can be stored in a venv directory by calling `python -m venv venv` from the top level directory and then installed by calling `pip install -r requirememts.txt`. There is a provided script to activate the virtual environment by calling `source env.sh`.

## Notes

To see a current list of the WIP efforts check the `TODO.md`, and to see the version histories of this repo chceck the `CHANGELOG.md`.
