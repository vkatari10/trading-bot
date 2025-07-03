# Source

The `src/` folder contains the backend logic of the trading bot and its runtime environment.<br>  

A brief overview:
## Overview of `src/` Structure

This platform is a modular, ML-powered trading system built for strategy experimentation, live inference, and end-to-end automation. Key components:  
- `api/`
  - Python service that:
    - Pulls historical data via `YFinance`
    - Serves trained ML models over REST for real-time inference
  - Interfaces directly with the live `runtime/` engine
- `backtesting/`
  - Strategy sandbox using `Backtrader` for testing indicator-based strategies on historical data
  - Future: will integrate ML models for simulated validation
- `ml/`
  - Full training pipeline that:
    - Loads config files from `./config`
    - Uses `/api` to source historical price data
    - Engineers features, trains models, and serializes them for deployment
  - Model-agnostic and designed for plug-and-play extensibility
- `runtime/`
  - Production trading engine written in `Go` that:
    - Ingests live market data (`Alpaca`)
    - Recomputes indicators per config
    - Calls `api/` for predictions
    - Places trades
    - Exposes its own API (repsented in `/frontend` and live dashboard)
