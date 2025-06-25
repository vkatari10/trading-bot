# TODO

## Major Problems (In order)
- Prevent negative share orders on Alpaca (check holdings before trade)  
- Finish backtesting engine MVP (use CLI wrapper for now)
- Support multiple stock trading / training
    - combine JSONs
        - instead of seperate faeture and label config use something like [this](../config/future.json)
    - training
        - Run a loop for training, just adapt the training module to accept a list of stocks and adjust accordingly
    - runtime
        - leverage `Go` for concurrency where computations are done in paralell 
        - Need to use channels
        - place technical indicator declerations in their own package
        - modularize other parts of `Go`
    - ML API
        - needs to be able to predict multiple perhaps at a time? 
        - might need to convert to websockets using `FastAPI` (scrap `Flask`)

## Long Term Improvements
- YFinance Dataframes only go down to the minute, so if someone wants to train an ML model on 1-sec data for example we need to make a logger in go that just takes the bars at some interval and then write as a csv to another file (with buffering) so then we can just make it a dataframe back in python

## Configs

Ability to put everything into a singular JSON file rather than having seperate features and labels.

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
