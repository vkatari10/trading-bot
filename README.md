![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge) ![CI](https://img.shields.io/github/actions/workflow/status/vkatari10/trading-bot/go-tests.yml?branch=main&label=Go%20Tests&style=for-the-badge&color=2ECC71) ![CI](https://img.shields.io/github/actions/workflow/status/vkatari10/trading-bot/python-tests.yml?branch=main&label=Python%20Tests&style=for-the-badge&color=2ECC71)
# ConTrade

**ConTrade** is a modular, config-driven trading platform designed for rapid prototyping and deployment of **machine learning based trading strategies**. It allows users to define everything from a single config file and easily run:
- ML training pipeline
- Historical validation backtests
- Live trading via Alpaca 

ConTrade is language-decoupled, real-time capable, and architected for both experimentation and deployment **without ever** touching core logic, only config files.

For more detailed information about ConTrade, visit the [docs](https://vkatari10.github.io/trading-bot/)

**Note**: This project is under active development, expect new features weekly. Below is a list of current features

## Core Features

### Unified Config Driven Architecture

- One config file defines **everything**: features, signals, ML settings, Live Settings, and more 
- Used for all services: Training, Backtesting, and Live Execution
- This eliminates config drift and ensures alignment for all services

### ML Pipeline (Python)
- Supports 
  - 12 Scikit-learn models
  - XGBoost
  - LightGBM
- Train Models with 
  - Multiple assets
  - Intervals from minutes to days
  - TA-Lib derived features supporting most technicals found [here](https://ta-lib.github.io/ta-lib-python/doc_index.html)
  - Custom labelling logic (crossover, thresholds, weights, persistence windows)
  - Hyperparameter tuning 
- Easily extensible to support more features, labels, and ML frameworks

### Backtesting (Python)
- CLI powered historical testing on asset strategies
- Simulate commission, slippage, latency 
- Quickly obtain profit/loss metrics 

### Live Execution Engine (Go)
- Serve ML models via FastAPI 
- Real time trading with Alpaca (paper or live accounts)
- High performance via
  - TA-Lib integration w/ custom C wrapper via `cgo`
  - Async ML prediction microservice via FastAPI + WebSockets
- ~190 concurrent asset tracking @ 60s intervals (on Alpaca free tier)
- 500ms floor for cycle times, supporting mid frequency trading
- Exposed API for trade monitoring and internal logging

### CLI Tools & TUI Dashboard
- Easily call commands and interact with ConTrade's features via the CLI
  - Training
  - Backtesting
  - Live Deployment
  - Tests
- Live TUI dashboard
  - View live:
     - Logging Information
     - Technical Data for all tickers

### Testing and CI 
- Unit tested across Python and Go
- Github Actions CI for automated checks

Project size is ~11000 total lines of code across ~50 files.

### Tech Stack Overview 
| Feature | Language | Libraries/Frameworks | External APIs |
|:-------:|----------|--------------|------|
| Config Files | JSON | - | - | 
| ML Training Pipeline | Python | Pandas, NumPy, Scikit-learn, XGBoost, LightGBM, TA-Lib | YFinance |
| Backtesting Engine | Python | -| YFinance |
| Runtime Engine | Go, Python | FastAPI, Gorilla WebSocket, TA-Lib, `cgo` | Alpaca Market |
| Risk Engine (WIP) | Go | - | Alpaca Account |
| CLI | Python | - | - |  
| TUI | Python | Textual | Internal Runtime Engine API | 
| Testing & CI | Go, Python, YAML | Pytest, `go test`, Github Actions | - | 

## Architecture

### Modular & Decoupled 
- Python handles ML and backtesting
- Go handles real-time data and broker execution
- Python and Go communicate **only through WebSockets** (no shared code)

### Config as a Contract 
- The JSON config acts as the **single source of truth** for all services
- No repeated logic and no mismatch between what's trained and what's deployed

![Architecture](docs/images/TradingPlatformDiagram7.svg)

### Why Use Go?
To support both low (minutes) and mid (seconds) frequency trading strategies `Go` provided a fair trade between speed and complexity

Here are alternatives and why they were not chosen:

- `Python`: GIL overhead and interpreted, not ideal for concurrency
- `Java`: Harder to iterate with compared to `Go`'s module system and single compiled binary
- `C++`: too complex (for this project), complicating concurrency, API calls, and other things for ultra low latency which is not the goal of this project

Go is fast enough to support both types of frequency strategies while allowing for faster iteration of code with simple concurrency and low latency

## Setup

To use this program you will need to follow these steps
- Clone this repository
- Set up the market data and broker API
  - [Register](https://app.alpaca.markets/account/login) for an Alpaca Account (Free)
  - Once signed in there will be both API keys and an API secret 
  - At the root is a [`.env.example`](.env.example) file, put the API key and secret in their proper places, and rename the file to `.env`
  - If you wish to test only, make sure you are using **paper money**
- Download `TA-Lib` dependencies
  - This project depends on the `TA-Lib` library 
  - Instructions [here](https://ta-lib.org/install/) and [here](https://github.com/TA-Lib/ta-lib-python?tab=readme-ov-file)
- Create the virtual environment for `Python` (Python3)
  - If on the bash shell run these commands
    - `python3 -m venv venv` (Creates the virtual environment folder)
    - `pip install -r requirements.txt` (Install Python dependencies)
    - `source venv/bin/activate` (Shorthand to start the virtual environment)
- Configuration
  - The JSON files are configurable, more info [here](config/README.md)
  - Make sure that if you wish to use a different configuration file, change the respective values in [`.env`](.env.example)
  - There are other settings available in the [`.env`](.env.example) file that impact things from refresh rates to api links. 

<!-- 
- Docker Container
  - Docker files are contained in docker/, as well as a `docker-compos,e`
  - 
- Build the docker image  
-->

## Limitations
- Ingestion of delayed data 
  - `Alpaca` free tier market API uses delayed data by default
  - Can be overcome with a higher tier `Alpaca` account
- Forced burn-in Period
  - During live execution we need to burn-in data 
  - `Alpaca` does not provide after hours market data 
  - To prevent major jumps from market close to market open, burn in is forced (or override for debug purposes)
- 60 second floor for live execution cycle rate
  - `Alpaca` market API + `YFinance` API induced limitation
  - `Alpaca` data only updates every minute, so polling market data within the same minute dilutes feature values
  - `YFinance` only has historical data down to a minute, so ML models cannot be trained on faster intervals
  - Therefore unless you can provide a higher tier `Alpaca` account AND higher frequency trained model, keep the `cycle_rate` in the JSON config `>=60`

## Documentation 
- [CONFIG](docs/CONFIG.md) - Learn the ConTrade custom JSON schema
- [ENV](docs/ENV.md) - Environment settings details
- [CLI](docs/CLI.md) - How to use the ConTrade CLI

## Notes
- [CHANGELOG](docs/CHANGELOG.md) - Version history details
- [TODO](docs/TODO.md) - Other long-term improvements
- [LICENSE](LICENSE) - License info 

**Disclaimer**: *This project is provided for educational and research purposes only. It is **not** financial advice. Use it at your own risk.* *The author is **not** responsible for any financial losses or damages resulting from the use of this software, including but not limited to trading with real money.*