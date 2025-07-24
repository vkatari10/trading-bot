# Creating your own config files

All services of **ConTrade** depend on the config file, which is specified for most CLI commands

This page explains 

1. Why the config file is important and what it manages

3. A breakdown of a config file example

4. What strategies are supported

5. Additional documentation to help write strategies

## Importance

Most major services **depend** on your config file as it tells each service important information

The config file covers these main parts
- Feature selection
- Label logic between features (How are signals generated between features?)
- ML settings
- Runtime settings
- Train/Trade Tickers
- (Optional) Backtesting defaults

It should be mentioned that the name of the config file does not matter, but **make sure** from the CLI you are passing the same file for each service. If you do not the system **will not** work as intended.

## Example

The best way to learn how to write a config file is to see one.

```JSON
{
    "features": [
        {
            "tech": "SMA",
            "args": {"timeperiod": 10},
            "name": "T_SMA_10"
        },
        {
            "tech": "SMA",
            "args": {"timeperiod": 30},
            "name": "T_SMA_30"
        },
        {
            "tech": "delta",
            "col1": "T_SMA_10",
            "col2": null,
            "name": "D_SMA_10_delta"
        },
        {
            "tech": "diff",
            "col1": "T_SMA_30",
            "col2": "T_SMA_10",
            "name": "D_SMA_30_10_diff"
        },
        {
            "tech": "delta",
            "col1": "T_SMA_30",
            "col2": "T_SMA_10",
            "name": "D_SMA_10_30_delta_diff"
        }
    ],
    "label_logic": [
        {
            "sig": "crossover",
            "name": "SMA_CROSS",
            "col1": "T_SMA_10",
            "col2": "T_SMA_30",
            "weight": 1.0
        },
        {
            "sig": "above",
            "name": "SHORT_SMA_ABOVE",
            "col1": "T_SMA_10",
            "col2": "T_SMA_30",
            "weight": 1.0,
            "persist": 2
        }
    ],
    "train_stocks": [
        "BA",
        "MSFT"
    ],
    "live_trade_stocks": [
        "AAPL",
        "AMZN", 
        "BA"
    ],
    "ml_settings": {
        "model_framework": "xgboost",
        "scikit_model_type": "RandomForrestClassifier",
        "model_name": "model_xgb.json",
        "model_training_timeframe": "1m",   
        "model_training_interval": "5d",
        "use_OHLCV_diffs": true,
        "hyperparameters": {
            "device": "cpu"
        }
    },  
    "runtime_settings": {
        "cycle_time": 15,
        "burn_window_time": 31,
        "log_api_flush_time": 2000, 
        "log_to_stdout": true,
        "run_after_close": false,
        "override_burn_in": true
    },
    "backtest_settings": {
        "starting_cash": 10000,
        "commission": 0.00,
        "position_size": 1
    } 
}
```

This is a basic SMA crossover strategy, but it sure does look like a lot, because **it is** doing a lot. 

Lets walk over the sections

### Features

This is where we define what we want our model to treat as features

A feature object in the JSON file array of `features` is **required** to have fields of:

- `tech`
    - The actual technical value being used
- `name`
    - The given name to refer to this object
    - A name in `features` should **always** start with either a `T_` or `D_`

Common Questions:

- How do I know which `tech` name to give?
    - `TA-Lib` [docs](https://ta-lib.github.io/ta-lib-python/funcs.html)
        - This link contains all supported features and their abbreviations
        - **ConTrade does not** support MAVP, MAMA, or **any** Math Transform/Operations
- What does `T_` or `D_` mean?
    - This is a naming convention to differentiate between **ConTrade** developed features and `TA-Lib` features
    - A quick rule of thumb: If you are using `diff` or `delta` use `D_`, and then  `T_` for all `TA-Lib` features
    - Beyond the prefix, it **does not** matter what you put
- How do I know which args to use?
    - `TA-Lib` [docs](https://ta-lib.github.io/ta-lib-python/funcs.html), click on one of the subsections say Overlap studies
    - Here's an example `real = SMA(real, timeperiod=30)`
    - For every supported technical listed on this website **only supply \*\*kwargs** to the JSON, do not include the positional arguments
    - i.e. if the arg has an equal sign **you must provide a value** into the JSON, if you are unsure about what value to use, plug the listed default arg(s) on the website 

The easiest part about naming your feature objects if that you can refer to them after they have been declared, for example notice how the `delta` and `diff` objects in the `features` array refer to the two `SMA` objects. This is also done in the `label_logic` as well. **You must** declare an object before you reference it. 


### Label Logic
```JSON
"label_logic": [
        {
            "sig": "crossover",
            "name": "SMA_CROSS",
            "col1": "T_SMA_10",
            "col2": "T_SMA_30",
            "weight": 1.0
        },
        {
            "sig": "above",
            "name": "SHORT_SMA_ABOVE",
            "col1": "T_SMA_10",
            "col2": "T_SMA_30",
            "weight": 1.0,
            "persist": 2
        }
    ],
```

Here you will now define what constitutes a buy/sell signal by creating relationships between declared objects in the `features` section. 

Let's walk through the two parts

1. What should be contained inside a `label_logic` object
2. Supported `sig` options

Quick note: Every object in the `label_logic` section **must** contain a `sig`, `name`, and `weight` field

| Field | Value | Explanation | 
|:-----:|-------|-------------|
| `sig` | string | Synonymous with `tech` from the features section, declares what type of signal generation between both objects must be followed |
| `name` | string | The name of this object, you **do not** need to follow the `T_` or `D_` naming conventions | 
| `col1` & `col2` | string | These must refer to existing objects in the `features` section. These are **order specific** which will be explained in the second table | 
| `weight` | float | Supports assigning weights to certain signals if you believe a relationship is more significant than another, for equal weights use `1.0` for every argument | 
| `persist` | int | Depending on the type of `sig` option, you can define how long a signal must be sustained for before classifying it as a buy/sell |


| `sig` | Args | Meaning | 
|:-----:|------|---------|
| `crossover` | `col1`: str, `col2`: str | if `col1` crosses above `col2` (buy or 1), else if `col1` crosses below `col2`  (sell or -1), else 0  | 
| `above` | `col1`: str, `col2`: str, `persist`: int | if `col1` is above `col2` for `persist` number of times in a row (buy or 1), else if `col1` is below `col2` -1, else 0 |
| `below` |  `col1`: str, `col2`: str, `persist`: int | if `col1` is below `col2` for `persist` number of times in a row (buy or 1), else if `col1` is above `col2` -1, else 0 |



### Train/Trade Stocks (Asset Tickers)
```JSON
   "train_stocks": [
        "BA",
        "MSFT"
    ],
    "live_trade_stocks": [
        "AAPL",
        "AMZN", 
        "BA"
    ]
```

| Field | Value | Explanation | 
|:-----:|-------|-------------|
| "train_stocks" | List[str] | Tickers that you want the ML model to train on |
| "live_trade_stocks" | List[str] | Number of tickers to manage concurrently, **note:** On a free Alpaca tier you can go up to 190 tickers @ 60s `cycle_rate`, assuming no hardware limitations | 


### Runtime Settings
```JSON
 "runtime_settings": {
        "cycle_time": 15,
        "burn_window_time": 31,
        "log_api_flush_time": 2000, 
        "log_to_stdout": true,
        "run_after_close": false,
        "override_burn_in": true
    },
```

| Field | Value | Explanation | 
|:-----:|-------|-------------|
| `cycle_time` | float | **Seconds** to recompute features, should align with granularity of ML model | 
| `burn_window_time` | int | How many `cycle_time`s to burn in data, should **always** be the value of the largest period or window in features |
| `log_api_flush_time` | float | **Milliseconds** to flush single log information payloads , configurable to prevent overwhelming API endpoints | 
| `log_to_stdout` | bool | Dev Tool to print log statements to `stdout` | 
| `run_after_close` | bool | For users with Free Alpaca tiers this **should** be set to `false`, else with higher tiers `true` if you would like after hours trading | 
| `override_burn_in` | bool | Dev Tool to skip the burn in process and start trading live without wait, uses randomly generated data in place of real stock data, **do not** use if you have configured real money trading API endpoints in `.env` | 

### Backtest Settings

```JSON
    "backtest_settings": {
        "starting_cash": 10000,
        "commission": 0.00,
        "position_size": 1
    } 
```

This section is actually completely optional, but it does make using the `backtest` command on the CLI a little bit easier if you want default values, but you can always specify on the CLI other arguments and these will be overridden. 

| Field | Value | Explanation | 
|:-----:|-------|-------------|
| `starting_cash` | float | How much cash to start the backtest with | 
| `commission` | float | How much the commission cost should be for each buy/sell trade made | 
| `position_size` | int | How many shares to buy for each trade |

## Supported Strategies

Currently **ConTrade** is implemented in such a way that ML models are **Classifiers** not **Regressors**
- This means ML models are trained to **predict buy/sell** signals, **not** predict actual asset prices
- Features define what the ML model should be aware of, but the `label_logic` **actually translates** the `features` to determine what a buy/sell signal actually is

## Helpful Materials

- [CLI Instructions](CLI.md) -- How to use the CLI to operate `ConTrade`
- [Environment Settings](ENV.md) -- configurable endpoints for Alpaca and Internal APIs
- [`TA-Lib` docs](https://ta-lib.github.io/ta-lib-python/funcs.html)