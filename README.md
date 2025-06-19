# AI/ML Based Trading Platform

This project is a modular, end-to-end trading platform that combines machine learning with real-time execution to empower users to develop, backtest, and deploy custom trading strategies seamlessly. Featuring a Python-based ML pipeline, a high-performance Go runtime engine, and a configurable monitoring dashboard, it supports hot-swappable models and live broker integration — making it ideal for low- to mid-frequency trading. Designed with flexibility and extensibility in mind, it enables both researchers and developers to experiment with technical indicators and ML-driven signals in a production-ready environment.

## Features
- Complete Machine Learning Pipeline
- Real-time trading engine
- User Defined Trading Logic
- Monitoring Dashboard
- Modular structure
- Ability to Support Low to Mid Frequency Trading

## Tech Stack
This project uses `Go` and `Python` to split the runtime engine and ML training repsonsibilities
- `Python` is used for all things ML
- `Go` is used for the runtime 
= `Next` is used for the front end dashboard

### Why Use Go?
`Python` or `C++` could have been used for runtime but:
- `Python` is intreperted, and also has GIL overhead2
- `C++` has higher complexity than `Go`, especially for multithreading and API calls

`Go` serves as a fair trade off between performance and complexity, allowing for fast iteration of low latency code.<br>

<!-- ## Architecture

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

-->
    
## Future Additions
- General 
  - More built-in technical indicators at training and run time
- Machine Learning Pipeline
  - Additions to train on multiple stocks
  - Allow users to tune hyperparameters
  - Allow uesrs to choose ML model (KNN, Linear Regression, etc.)
- APIs
  - Find a higher frequency market data streamer
  - Convert ML API to use `FastAPI` instead of `Flask`
- Runtime
  - Be able to track multiple stocks at once

## Limitations

There are several limitaions with this trading bot, however the most significant are listed below.<br>
- Ingestion of delayed data, up to 1-15 minutes
- Data refresh rates of 1 minute, limiting predictions to every minute
- Burn-in period during live execution (~30 minutes)

## Requirements

In order to make and view trades you will need an Alpaca API account. The API keys will need to be stored in a `.env` file in the top level directory. You will also need to register for a finnhub account to stream real time data as well to your those associated keys as well. These should also be stored in the `.env` file.<br>

Required packages are listed in the requirements.txt, those can be stored in a venv directory by calling `python -m venv venv` from the top level directory and then installed by calling `pip install -r requirememts.txt`. There is a provided script to activate the virtual environment by calling `source env.sh`.

<!-- 
To use this program you will need to follow these steps
- Set up market data and broker API
  - [Register](https://app.alpaca.markets/account/login) for an Alpaca Account (Free)
  - Once signed in there will be both API keys and an API secret 
  - Copy those into `src/runtime/go-src/.env.example`
  - Rename `.env.example` to `.env` 
- Create the virtual environment for `Python` (Python3)
  - If on bash shell run these commands
    - `python3 -m venv venv` (Creates the virtual environment folder)
    - `pip install -r requirements.txt` (Install Python dependenciesz)
    - `source ./scripts/env.sh` (Shorthand to start the virtual environment)
- Config
- Build the docker image  


-->

## Notes
- [CHANGELOG](docs/CHANGELOG.md)
- [TODO](docs/TODO.md)
