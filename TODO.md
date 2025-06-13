# TODO

# High Priority 
- Make the frontend not look as rough, improve the CSS styling
- Create way to check if we have the stock are not stop negative shares on Alpaca


## Long Term Improvements <!--===========================================-->

- Add custom backtesting engine on trained models to test performance on historical data

### ML
- [ ] Potentially allow user to decide which ML model to use (JSONs)
- [ ] Add support for different ML model types (KNN)

### Trading Logic
- [ ] Heighten buy, sell strictness (one indicator should not dictate a buy/sell action). Also include a way to determine the sensitivity of a signal to increase and decrease the ordering amount of shares accordingly.

### API

- [ ] Figure out how use websockets and gRPC to improve ML <-> Go communication
- [ ] Also use websockets in Go to work with Alpaca to improve Broker commumincation





