# .env Settings

The [`.env`](../.env.example) file contains many settings that can be changed according to your needs.  

This documentation serves to help get new users accustomed to the settings that are available and what they do  

**Note:**  If changes are made to the `.env` file you must `source .env` or run `. .env` at the root to update these values. 

## Main Config Files

Learn more about creating your own JSON configurations [here](CONFIG.md)

| Setting | Description | Default |
|:-------:|-------------|---------|
| FEATURE_CONFIG_FILE   |  Contains the `JSON` config file inside `/config` to use for feature, live execution technicals, backtesting, and more | [features.json](../config/features.json) |
| LABEL_CONFIG_FILE     |  Contains the `JSON` config file inside `/config` to use for labelling logic at ML training time | [signals.json](../config/signals.json)| 

## Alpaca API 
| Setting | Description | Default |
|:-------:|-------------|---------|
| ALPACA_API | Alpaca API associated with your own Alpaca account |- |
| ALPACA_SECRET | Alpaca secret key associated with your own Alpaca account |- | 
| ALPACA_XYZ_LINK | All other Alpaca API links DO **NOT** need to be modified if you plan to use paper money. Adjust base URLs if you want to use real money and understand the risk of using non-simulated funds | Various | 

## Runtime Engine Settings
| Setting | Description | Default |
|:-------:|-------------|---------|
| REFRESH_RATE_TIME | Time in **seconds** to pull a new quote. In other words the frequency of the bot's actions. **NOTE**: If you do not have premium data it is **NOT** suggested to go below 60 as Alpaca free tier updates quotes every 60 seconds. Going below 60 with a free tier account will lead to degraded ML inference (The same is true for ML models not train on <60 historical data) | 60 |
| LIVE_TRADE_TICKER* | Ticker to trade on during live execution | AAPL |
| BURN_IN_WINDOW_TIME | Time in **minutes** of how long to burn in data before starting to executing trades. Should be the highest value of the longest window of a technical indicator (**IF USING delta or diff objects** add 1 to this value) | 31 (If SMA 30 and a Diff object were used, else without Diff could just be 30) |
| ENGINE_LOG_API_FLUSH_TIME | Time in **milliseconds** to flush each item in the buffer to the destination of the LOG_API_LINK | 2000 |
| LOG_API_LINK | Destination of log updates as a JSON, should be an API routing link of your choice. The update rate is specified by the `ENGINE_LOG_API_FLUSH_TIME` | `http://localhost:3000/api/log` |
| DATA_API_LINK | Destination of live technical data as a JSON should be an API routing link of your choice. The update rate is specified by `REFRESH_RATE` | `http://localhost:3000/api/data` |
| ENV_API_LINK | Destination of the environment variables inside .env as a JSON, only sent once. | `http://localhost:3000/api/env` |
| BROKER_API_LINK | Destination of the brokerage account information as a JSON. The update rate is specified by the `REFRESH_RATE` | `http://localhost:3000/api/broker` |
| LOG_TO_STDIO | Dev tool to print log statements to standard output for easier debugging | TRUE or FALSE (FALSE if in production)
| ALWAYS_RUN | Dev tool to execute the runtime engine even when the market is closed | TRUE or FALSE (FALSE if in production) | 
| OVERRIDE_BURN_IN | Dev tool to override burn data to start live execution immediately by using dummy data | TRUE or FALSE (FALSE if in production) |

## ML API Settings
| Setting | Description | Default |
|:-------:|-------------|---------|
| ML_API_LINK | Destination to host the ML API Server | `http://127.0.0.1:5000` | |

<!-- ## ML Pipeline Settings
| Setting | Description | Default |
|:-------:|-------------|---------|
| TRAIN_TICKER* | The ticker stock to train on | AAPL | 
| TRAIN_MODEL_TYPE* | **UNUSED** Type of model to use at train time | -|
| TRAIN_DF_TIMEFRAME | How far back to download historical data from `YFinance` | Click [here](https://ranaroussi.github.io/yfinance/reference/api/yfinance.download.html#yfinance.download) to view options| 
| TRAIN_DF_INTERVAL | Interval from historical data to get prices from `Yfinance`  | Click [here](https://ranaroussi.github.io/yfinance/reference/api/yfinance.download.html#yfinance.download) to view options | 
| DUMP_MODEL_NAME | Model name that should be written, overwrites an existing model if it shares the same name |-| 
| RUNTIME_MODEL_NAME | **UNUSED** Model to use at runtime if using live execution | -|  -->

*Stars next to a setting indicate settings that will later be incorporated into JSON config files.*