# .env Settings

The [`.env`](../.env.example) file contains many settings that can be changed according to your needs.  

This documentation serves to help get new users accustomed to the settings that are available and what they do  

**Note:**  If changes are made to the `.env` file you must `source .env` or run `. .env` at the root to update these values. 

## Alpaca API 
| Setting | Description | Default |
|:-------:|-------------|---------|
| ALPACA_API | Alpaca API associated with your own Alpaca account |- |
| ALPACA_SECRET | Alpaca secret key associated with your own Alpaca account |- | 
| ALPACA_XYZ_LINK | All other Alpaca API links DO **NOT** need to be modified if you plan to use paper money. Adjust base URLs if you want to use real money and understand the risk of using non-simulated funds | All default paper links | 

## Runtime Engine Settings
| Setting | Description | Default |
|:-------:|-------------|---------|
| LOG_API_LINK | Destination of log updates as a JSON, should be an API routing link of your choice. The update rate is specified by the `ENGINE_LOG_API_FLUSH_TIME` | `http://localhost:3000/api/log` |
| DATA_API_LINK | Destination of live technical data as a JSON should be an API routing link of your choice. The update rate is specified by `REFRESH_RATE` | `http://localhost:3000/api/data` |
| ENV_API_LINK | Destination of the environment variables inside .env as a JSON, only sent once. | `http://localhost:3000/api/env` |
| BROKER_API_LINK | Destination of the brokerage account information as a JSON. The update rate is specified by the `REFRESH_RATE` | `http://localhost:3000/api/broker` |

## ML API Settings
| Setting | Description | Default |
|:-------:|-------------|---------|
| ML_API_LINK | Link to host the ML API Server | `http://localhost:8000` | 

<!-- ## ML Pipeline Settings
| Setting | Description | Default |
|:-------:|-------------|---------|
| TRAIN_TICKER* | The ticker stock to train on | AAPL | 
| TRAIN_MODEL_TYPE* | **UNUSED** Type of model to use at train time | -|
| TRAIN_DF_TIMEFRAME | How far back to download historical data from `YFinance` | Click [here](https://ranaroussi.github.io/yfinance/reference/api/yfinance.download.html#yfinance.download) to view options| 
| TRAIN_DF_INTERVAL | Interval from historical data to get prices from `Yfinance`  | Click [here](https://ranaroussi.github.io/yfinance/reference/api/yfinance.download.html#yfinance.download) to view options | 
| DUMP_MODEL_NAME | Model name that should be written, overwrites an existing model if it shares the same name |-| 
| RUNTIME_MODEL_NAME | **UNUSED** Model to use at runtime if using live execution | -|  -->