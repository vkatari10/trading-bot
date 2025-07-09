# TODO

## In Progress

### ML Pipeline

#### CURRENT
- Conversion to `XGBoost` 
- Hyperparameter tuning via config files\

#### FUTURE
- Configurable label rebalancing for training
- Allow users to use `Pytorch` NN as well 

#### MAYBE?
- None

### Backtesting 

#### CURRENT
- Graphical representation of buy/sell history using `matplotlib`

#### FUTURE
- More technicals computed rather than just P/L

#### MAYBE?

### Live Execution 

#### CURRENT
  - None

#### FUTURE
  - Websockets on Alpaca API market data calls 
  - Use a singular call with all tickers (comma separated) instead of calls on each eventloop for each ticker

#### MAYBE?
  - gRPC, Redis use between ML API and Runtime Engine

### UI
#### CURRENT
- Improve CLI styling
#### FUTURE
- TUI Client
  - Using `Rich`
  - Live dashboard monitoring logging, account and data information
  - Filter between stocks, logs, and more 
- JSON builder
  - some sort of graphical interactive way to build JSONs rather than having to type them out manually
- DSL Validator 
  - Validate user JSONs before accepting/parsing them
#### MAYBE?
- None
### Risk Module 
#### CURRENT
- Implement MVP
#### FUTURE
- Prevent negative shares to go through (why does Alpaca allow you to sell more than you have? Maybe its shorts?)
- Increase buy/sell signal strictness; implement sensitivity control for order sizing
#### MAYBE?
- None

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



## Long Term Improvements
- YFinance DataFrames only go down to the minute, so if someone wants to train an ML model on 1-sec data for example we need to make a logger in go that just takes the bars at some interval and then write as a csv to another file (with buffering) so then we can just make it a DataFrame back in python
- `Pytorch` is a NN how do we convert training DFs to be used in NN objects?
- Maybe move to a conventional front end using `React` instead of a TUI (like a SaaS wrapper)