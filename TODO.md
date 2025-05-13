# TODO

## alpaca_api/

-Potentially Combine all three files into a proper class rather than just mehods to reduce import calling later.

## backtesting/

-Refine strategy to not be as simple and dependent on a singular technical indicator in the future.

## data/

-Add more methods into `technicals.py` in order to fit more complex technicals onto the dataframes to trade the ML models on.

## engine/

-Complete model training first

## exceptions/

-if refining the `alpaca_api/` folder then add more exceptions as necessary

## finnhub_api/

-Comlpete engine first to decide if more live data needs to be pulled than just the live OHLCV data.

## live_data/

-Complete model training first, may just be able to use the same methods from data/technicals.py just slightly modified

## models/

- Decide what specific model to use (Knearestneighbors, Gradient Boosting, Random Forest)
- Somehow come up with a method to train the models and then return the model using `pickle`.

## server/

- Finish engine first will use flask to create an API probably with just log information of what the model on live data is doing

## yfinance_api/

- If needed add a new method that can specify the start and end dates as necessary