![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) ![Python](https://img.shields.io/badge/python-3670A0?style=for-the-badge&logo=python&logoColor=ffdd54) ![scikit-learn](https://img.shields.io/badge/scikit--learn-%23F7931E.svg?style=for-the-badge&logo=scikit-learn&logoColor=white)  [![Licence](https://img.shields.io/github/license/Ileriayo/markdown-badges?style=for-the-badge)](LICENSE)

# StratForge 

StratForge is a platform for building and running custom trading bots, where users define strategies through configuration, no code edits required. It handles the entire ML pipeline, backtesting, and live trading infrastructure, making it easy to prototype and deploy new ideas quickly.

Note: This project is actively under improvement. Here is a list of current and upcoming features. 

## Completed Features

### Config Driven Architecture (JSON)
- Choose a strategy from supported technical indicators (EMA, Delta, etc.) 
- Declare labelling logic as the relationships between technical indicators
- Every service uses the same config, eliminating feature misalignment 
### Full ML Pipeline (Python)
- Scikit-learn Random Forest Classifier Model 
- Multi-asset training (configurable)
- Dynamically adjusts features, labels, and training tickers based on config files
### Backtesting Module (Python)
- Quickly test trained ML models on historical data to analyze P/L
### Live Execution w/ Broker Integration (Go)
- Runtime engine
  - Polls Market Data
  - Manages delcared config value features
  - Communicates via HTTP REST with a ML API Server for real time inference
  - Places trades (paper or real) using individual Alpaca accountz
  - Exposes own API for live monitoring
  - Error handling from bad configs, failed API calls, etc
  - Supports trading intervals from minutes to seconds

## In Progress
### ML Pipeline
  - Conversion to `XGBoost` models exclusivley 
  - Hyperparameter tuning via config files
  - Configurable label rebalancing for training
### Backtesting Module
  - Graphical representation of buy/sell history using `matplotlib`
### Live Execution 
- Reduce ML inference latency by using gRPC
- Reduce market data polling latency by using websockets
-  Multi asset trading configurable by JSON
### UI
- CLI
  - Python based CLI for quick jobs (training, backtesting)
- TUI Client
  - Using `textual`
  - Planned live dashboard monitoring logging, account, and data information 
  - Easier interface to modify `.env` or `JSON` configs
### Other 
- Support for more technical indicators / features
  - Current
    - EMA
    - SMA
    - Deltas (and Difference of Deltas)
  - Future
    - RSI
    - MACD
    - Bollinger Bands

Project size is around 2000 LOC divided between ~60 files. 

## Architecture
Philosophy
- Config Driven 
  - Prevent feature misalignment by ensuring every service uses the same config setup
  - Think of the config like a contract that every other service must agree to
- Decoupled Systems
  - All services are seperated and independent
  - At Runtime
    - The ML Model runs on its own server
    - The Go runtime computes features and recieves inferences via HTTP REST calls

This design allows for independent parts to be scaled and improved without affecting other systems

![Architecture](docs/images/TradingPlatformDiagram4.svg)

## Tech Stack
- `Go`
  - Used for the runtime engine during live execution 
- `Python`
  - Used for ML Pipeline
    - `Pandas`
    - `NumPy`
    - `Scikit-learn`

### Why Use Go?
You might be wondering why `Go` was used for this project and not `Python` or `C++` as the runtime engine.<br>

Well, since this project is intended to be used for both low and mid frequency trading:
- `Python` is interpreted and has GIL overhead
- `C++` is too complex (for this project), and complicates API calls and multithreading among other things

`Go` served as a fair trade off between performance and complexity, allowing for fast iteration while still providing low latency code.<br>
    
## Setup

To use this program you will need to follow these steps
- Clone this repository
- Set up the market data and broker API
  - [Register](https://app.alpaca.markets/account/login) for an Alpaca Account (Free)
  - Once signed in there will be both API keys and an API secret 
  - At the root is a [`.env.example`](.env.example) file, put the API key and secret in their proper places, and rename the file to `.env`
  - If you wish to test only, make sure you are using **paper money**
- Create the virtual environment for `Python` (Python3)
  - If on the bash shell run these commands
    - `python3 -m venv venv` (Creates the virtual environment folder)
    - `pip install -r requirements.txt` (Install Python dependencies)
    - `source ./scripts/env.sh` (Shorthand to start the virtual environment)
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
The most significant limitations of this platform are
- Ingestion of delayed data
  - More of an `Alpaca` issue since the free tier for market data is delayed by default
- Burn-in period during live execution 
  - This is caused by a current lack of market data available after hours, so using the same technical indicator values from market close at market open could lead to large jumps
  - Pair this with an ML model, it could lead to unwanted predictions or trades, so as of now burn-in is required at live execution to initialize technical indicator values for the sake of ML inference
- Lack of support for higher frequency trading (`<60` seconds refresh rates)
  - This again is another issue with `Alpaca` free tier since data only updates once every minute, unless you have access to higher frequency data
  - But also caused by the `YFinance` API, used for ML training, only goes down to bars of 1 minute for historical stock data
  - Therefore it is a good idea not to go below `60` seconds for the refresh rate for the runtime engine unless 
    - You have another pickled ML model that was trained on higher frequency data
    - You have access to a higher tier `Alpaca` account 
  - It's important to note that the runtime engine can handle refresh rates for small intervals
    - But if you lack the requirements above
      - Technical data becomes diluted from the same stock price being used
      - Which in turn degrades ML inference
      - Or your ML model was not trained on such high frequencies
      - Causing poor performance
      - So be careful!

## Testing

All `Python` based tests are contained in `./tests/` and all `Go` based tests are located in `./src/runtime/go-src/tests`.<br>

You can call `./scripts/tests/` from the root to run both `Python` and `Go` tests.

## Documentation 
- [ENV](docs/ENV.md) - Understand how the `.env` file is used
- [CONFIG](docs/CONFIG.md) - Learn how to construct your own strategies using our custom JSON schema

## Notes
- [CHANGELOG](docs/CHANGELOG.md) - Version history 
- [TODO](docs/TODO.md) - List of current efforts and future 
updates
- [LICENSE](LICENSE) - License info 

*This platform is intended for research and development purposes only, please use paper trading or simulated funds to prevent real financial losses.*