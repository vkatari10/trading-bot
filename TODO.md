# TODO

## alpaca_api/ <!--============================================== -->

# Medium Level
- [ ] Potentially Combine all three files into a proper class rather than just mehods to reduce import calling later.

## backtesting/ <!--============================================== -->

### Low Level
- [ ] Refine strategy to be more strict and use more than basic indicators

## data/ <!--============================================== -->

### High Level
- [ ] Create more methods in `signal.py` to allow for signal calculations to be easilyt done on realtionhip based columns on the `DataFrame`

### Low Level
- [ ] Add more methods into `technicals.py` in order to fit more complex technicals onto the dataframes to trade the ML models on.

## engine/ <!--============================================== -->

- [ ] Complete model training first

## exceptions/ <!--============================================== -->

### Low Level
- [ ] if refining the `alpaca_api/` folder then add more exceptions as necessary

## finnhub_api/ <!--============================================== -->

### Low Level
- [ ] Comlpete engine first to decide if more live data needs to be pulled than just the live OHLCV data.

## live_data/ <!--============================================== -->

### Medium Level
- [ ] Complete model training first, may just be able to use the same methods from data/technicals.py just slightly modified

## models/ <!--============================================== -->

### High Level
- [ ] Fix the model to have higher overall F1 score instead of the current 32% accurate one.
- [ ] Decide what specific model to use (Knearestneighbors, Gradient Boosting, Random Forest)

## server/ <!--============================================== -->

- [ ] Finish engine first will use `FAST API` to create an API probably with just log information of what the model on live data is doing

## yfinance_api/ <!--============================================== -->

- [ ] If needed add a new method that can specify the start and end dates as necessary

# Long Term Improvements <!--===========================================-->
- [ ] Move all API calls to use `Go` instead of python
- [ ] Potentially Adjust ML models to use a different learning technique
- [ ] Move DataFrame Processing to CPP code instead of python and external libraries
- [ ] Determine a more complex set of technicals
- [ ] Heighten buy, sell strictness
- [ ] Find an alternative to `finnhub` that updates faster