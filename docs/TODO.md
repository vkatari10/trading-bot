# TODO

## High Priority
- Improve frontend CSS styling to reduce rough UI  
- Prevent negative share orders on Alpaca (check holdings before trade)  

## Long Term Improvements
- Add custom backtesting engine to evaluate models on historical data  
- Support trading multiple stocks concurrently  
- YFinance Dataframes only go down to the minute, so if someone wants to train an ML model on 1-sec data for example we need to make a logger in go that just takes the bars at some interval and then write as a csv to another file (with buffering) so then we can just make it a dataframe back in python

### ML
- [ ] When multiple stock tracking at execution time is complete, create a JSON config to declare which stocks to use
- [ ] same thing as above but for train time as well 
- [ ] Add support for additional ML models (e.g., KNN, Linear Regression), configured via `.env`  

### Trading Logic
- [ ] Increase buy/sell signal strictness; implement sensitivity control for order sizing 
- [ ] Potentially create a dedicated package in Go to determine how to to order size based on predictions

### API
- [ ] Integrate websockets and gRPC for improved ML <-> Go communication (FastAPI)
- [ ] Use websockets to optimize Alpaca broker interaction  
