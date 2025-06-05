# Logic

This folder contains JSONs that the ML training pipeline relies on along with the runtime engine.

## Training Features

In `features.json` a single array of objects should be loaded with the technical indicators that the ML should treat as features that will be exposed at runtime.<br>

A quick example:<br>
```
[
    {
        "tech": "SMA",
        "window": 10,
        "name": "SMA_10",
        "delta": true
    },
    {
        "tech": "SMA",
        "window": 30,
        "name": "SMA_30"
        "delta": true
    }
]
```

### Supported Technical Indicators

For each JSON object you will need a `name` field in addition to the values described below for each technical indicator<br>

- SMA
  - "tech": "SMA"
  - "window": int
- EMA
  - "tech": "EMA"
  - "window": int

If we want to expose differences or deltas of techincal indicators create new objects with the specified values<br>

"col1" and "col2" must already declared in previous objects

- Deltas (Rate of Change)
  - "tech": "delta"
  - "col1": string
  - "col2": string or null
  - Note: If "col2" is left as null then only the delta of the first column will be taken
- Difference 
  - "tech": "diff"
  - "col1": string 
  - "col2": string

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

- Crossover (col1 crosses above col2: buy, else sell)
  - `"col1"`: If this crosses above `col2` it signifies a buy signal, else sell
  - `"col2"`: Other column
- Cross Up Only (col1 crosses above col2: buy, else nothing)
  - `"col1"`: If this crosses above `col2` it signifies a buy signal, else 
  - `"col2"`: Other column to compare to  
- Cross Down Only (col1 crosses below col2, buy, else nothing)
  - `"col1"`: If this crosses below  `col2` it signifies a buy signal, else sell
  - `"col2"`: Other column to compare to 
  

