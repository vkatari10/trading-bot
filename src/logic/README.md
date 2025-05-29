# Logic

This folder cotains the user defined logic to determine the technical indicators to train the ML models on and the real time indicators needed to compute in real time.<br>

If you would like to customize the technical indicators and relationships between to train the ML models please read below.<br>

## Training Features

In `features.json` a single array should be loaded with the technical indicators that the ML model should use (note that a lot of technical indicators have not been implemented yet) and also allows the runtime engine to determine which indicators to compute in real time.<br>

A quick example:<br>
```
[
    {
        "tech": "SMA",
        "window": 10
    },
    {
        "tech": "SMA",
        "window": 30
    }
]
```

### Supported Technical Indicators

Currently Supported Technical Indicators and associated arguments are found below.<br>

- SMA
  - "tech": "SMA"
  - "window": int
- EMA
  - "tech": "EMA"
  - "window": int

## Determine Buy and Sell signals

In `signals.json` a single array should be loadead with the relationships between the technical indicators delcared in `features.json` to determine what constitutes a buy/sell signal.<br>

A quick example:<br>
``` 
[
    {
        "sig": "crossover",
        "name": "SMA_CROSS",
        "col1": "SMA_10",
        "col2": "SMA_30"
    }
]
```
This object means a buy signal is indicated when `SMA_10` crosses above `SMA_30` or a sell signal is indicated when `SMA_10` crosses below `SMA_30`.

### Supported Signals

All objects declared in `signals.json` must contain a `"sig"` and `"name"` field.<br>

Currently supported signals are found below.<br>

- Crossover (col1 crosses above col2: buy, else sell)
  - `"col1"`: If this crosses above `col2` it signifies a buy signal, else sell
  - `"col2"`: Other column
- Cross Up Only (col1 crosses above col2: buy, else nothing)
  - `"col1"`: If this crosses above `col2` it signifies a buy signal, else 
  - `"col2"`: 
- Cross Down Only (col1 crosses below col2, buy, else nothing)
  - `"col1"`: If this crosses below  `col2` it signifies a buy signal, else sell
  - `"col2"`: Other column
  

