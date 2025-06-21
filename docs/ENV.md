# .env Settings

The [`.env`](../.env.example) file contains many settings that can be changed according to your needs.  

This documentation serves to help get new users accustomed to the settings that are available and what they do

## Alpaca API Settings

### ALPACA_API_KEY

Alpaca API key associated with your own account

### ALPACA_SECRET_KEY

Alpaca secret key associated with your account

### ALPACA_XXXX_LINK

All other Alpaca API links DO NOT need to be modified if you only plan to use paper money.   

IF you choose to use real money then just update the base URL to your needs

## Runtime Engine Settings

### REFRESH_RATE

Time in seconds to pull a new quote. In other words how quickly do you want your bot to move at.   

**NOTE:** if you do not have premium data it is NOT suggested to go below 60 as `Alpaca` free tier updates only every 60 seconds. Going below 60 with a free tier acconunt will lead to degraded ML inference performance. 

### TICKER

Ticker to trade on during live execution. Example: AAPL

### BURN_IN_WINDOW_TIME

Time in minutes of how long to burn in data, should be the value of the longest window of a technical indicator.  

**If you use a `delta` or `diff` object**: Add 1 to this value.

### ENGINE__LOG_API_FLUSH_TIME

Time in milliseconds of how fast to flush the API buffer for each item in the buffer to the destination of the `DATA_API_DEST_LINK`.

### LOG_LINK

Logging information of the bot as a JSON, should be set to an API routing link of your choice, will update at a rate specified by the `ENGINE_API_FLUSH_TIME` to get each new log

### DATA_LINK

Live Data information of technical indicators of the bot as an JSON, should be set to an API routing link of your choice

### ENV_LINK

One time JSON payload send to send environment variables contained inside the `.env` file, should be set to an API routing link of your choice

### BROKER_LINK

Live brokerage account information of the account associated with your `Alpaca` API keys in real time as a JSON, shoud be set to an API routing link of your choice

## ML API Settings

### ML_API_LINK

UNUSED

## ML Pipeline Settings

### FEATURE_CONFIG_FILE

Contains the `JSON` config file to use for features in the ML pipeline

### LABEL_CONFIG_FILE

Contains the `JSON` config file to use for labelling logic in the ML pipeline

### TRAIN_TICKER

The ticker as a stock to train on. Example: AAPL

### TRAIN_MODEL_TYPE
UNUSED

### DUMP_MODEL_NAME
UNUSED

### RUNTIME_MODEL_NAME
UNUSED