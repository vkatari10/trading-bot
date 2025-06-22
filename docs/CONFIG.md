# JSON Configurations

All features of BotBuilder are dependent on JSON configs, knowing how to structure them properly and ensuring that they are being used will smoothen the experience with this suite.  

On this page you will learn 
- Why the JSON config files are so important
- How to set up the JSON config properly
- Avaiable technical inidcators and definitions
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
    - Beyond that it does not matter


### Available Technical Indicators

### Adjusting for new config files in `.env`
