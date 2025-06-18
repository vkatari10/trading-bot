# TODO

## High Priority (Next 2 weeks)
- Improve frontend CSS styling to reduce rough UI  
- Prevent negative share orders on Alpaca (check holdings before trade)  
- Implement tests folder and add basic unit/integration tests  

## Long Term Improvements (Q3–Q4)
- Add custom backtesting engine to evaluate models on historical data  
- Support trading multiple stocks concurrently  

### ML
- [ ] Allow user to choose ML model via JSON config  
- [ ] Add support for additional ML models (e.g., KNN)  

### Trading Logic
- [ ] Increase buy/sell signal strictness; implement sensitivity control for order sizing  

### API
- [ ] Integrate websockets and gRPC for improved ML <-> Go communication  
- [ ] Use websockets to optimize Alpaca broker interaction  
