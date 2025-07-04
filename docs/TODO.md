# TODO


## In Progress
### ML Pipeline
  - Conversion to `XGBoost` 
  - Hyperparameter tuning via config files
  - Configurable label rebalancing for training
  - FUTURE: allow users to us e `Pytorch` as well 
### Backtesting Module
  - Graphical representation of buy/sell history using `matplotlib`
### Live Execution 
- Reduce ML inference latency by using gRPC
- Reduce market data polling latency by using websockets
- Enhanced Support for multiple live tickers
    - We need to use channels to support API calls not being repeated 
    - Broker API call for each cycle gets repeated N times where N is the number of stocks being managed at a single time
    - ML API needs to be able to support multiple calls at once (try `FastAPI` first)
    - Configurable retry logic for API failure 
### UI
- TUI Client
  - Using `textual`
  - Planned live dashboard monitoring logging, account, and data information
  -Easier interface to modify `.env` or `JSON` configs
### Risk Module 
- Prevent negative shares to go through (why does Alpaca allow you to sell more than you have? Maybe its shorts?)
- Increase buy/sell signal strictness; implement sensitivity control for order sizing 

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