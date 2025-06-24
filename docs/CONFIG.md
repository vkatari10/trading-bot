# JSON Configurations

All features of BotBuilder are dependent on JSON configs, knowing how to structure them properly and ensuring that they are being used will smoothen the experience with this suite.  

On this page you will learn 
- Why the JSON config files are so important
- How to set up the JSON config properly
- Avaiable technical indicators and definitions
- How to adjust the `.env` to use specific files

## Importance

The reason why these JSON config files are so important is because every part of this suite depends on the strategy that you choose.  

Dependencies
- ML Pipeline (What features do we train on? What labelling logic do we use?)
- Backtesting (What user defined strategy do we need to base our trades on?)
- Live execution (What do we need to compute for ML model in real time?)  

Being able to write these JSON config files properly ensures that your strategies are interpreted properly and prevent any unwanted behavior/inferences.

## General Setup

In general you will need two files  
1. Declares the features (technical indicators) that should be used
2. Declares the labels (relationships between technical indicators) that represent buy/sell signals

- File naming
    - Should end in `.json`
    - Beyond that it does not matter since they can be configured in the `.env` flie 
- How to write
    - Should be a single array of objects (`[{}, {}, {}]`)
    - Each object is required to have two fields in order to work 
        - "name" = name of the specifc object
        - "tech" = type of technical indicator
    - If a object does not have these fields then the configuration file **Will not work**  

Beyond this, there is not much more you need to know but most technical inidcators reqiure certain fields.  

Let's look at an example 

```JSON
[
    {
        "tech": "SMA",
        "window": 20,
        "name": "SMA_20"
    },
    {
        "tech": "EMA",
        "window": 20,
        "name": "EMA_20",
        "smoothing": 2
    },
    {
        "tech": "Delta",
        "col1": "SMA_20",
        "col2": null,
        "name": "SMA_20_Delta"
    }
]
```  

What does this show?
- Some technical indicators require certain fields
    - For `EMA` we need a `smoothing` field whereas `SMA` does not require this
- We can use previously declared objects down the list
    - Look at the third object we use the `SMA_20` column to find deltas of it
    - The objects are interpeted in order, imagine the program sweeping from top the bottom of the list of objects
    - **Therefore**: Declare objects in order, do not use something that does not exist yet (like if you were writing in an intpreted language)  

We can see that making config files is straightfoward which makes this platform so powerful for testing new srategies. The hardest part is learning how to properly declare each technical indicator, which is explained below 

## Available Technical Indicators

Remember: each object must contain a `tech` field, which **must** match the example in the table, and a `name` field which can be one of your choice.  

### SMA - Simple Moving Average  
| Field | Description | Data Type | Example |
|:-----:|-------------|-----------|---------|
| tech | identifier | string | sma |
| window| Window of the SMA | int | 20      |  

### EMA - Exponential Moving Average 
| Field | Description | Data Type | Example |
|:-----:|-------------|-----------|---------|
| tech | identifier | string | ema |
| window | Window of the EMA | int | 20
| smoothing | Smoothing value to apply for calculations, usually 2 for most cases | int | 2 |

### Delta - Rate of Change Values

| Field | Description | Data Type | Example |
|:-----:|-------------|-----------|---------|
| tech | identifier | string | delta |
| col1 | Object to find the delta values of, should be the `name` field of another object already declared | string | SMA_10 |
| col2 | Not required, can be left as `null` but if you want to find the delta of differences then include another object that was also preivously declared | string/null | SMA_30 |  

### Diff - Difference between two Objects
| Field | Description | Data Type | Example |
|:-----:|-------------|-----------|---------|
| tech | identifier | string | diff |
| col1 | computed as col1 - col2 | string | SMA_10 | 
| col2 | **Reqiured** cannot be left as `null` | string | SMA_30 |



### Adjusting for new config files in `.env`

At the top of the `.env` file are two options called `FEATURE_CONFIG_FILE` and `LABEL_CONFIG_FILE`.  

If
- You modify an object 
    - just call `./scripts/env.sh` to update the `.env`
- You create a new JSON file
    - redirect the name starting from `/config`, i.e do not include the `config` path in the value  

To see other options available in `.env` read [here](ENV.md)