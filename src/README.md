# Trading Bot Source Files

This folder contains the backend logic of the trading bot from the testing of technical indicators on historical data to the actual runtime environment.<br>

Future Additions to this folder will contain dedicated README's to explain which each individual component does, however since this repo is a work in progress there can be sudden changes and I will avoid writing entire README's until there is a full stack MVP that implements the current structure.<br>

A brief overview of what each folder contains follows
- `api/`
  - Contains both external APIs for brokers, market data, and historical data as well as internal APIs for local machine learning API servers and to expose the runtime environment.
- `backtesting/`
  - Standalone folder that allows users to test their own strategies utilizing the `Backtrader` library in python to test the effectiveness of certain technical indicators on historical data. This folder has no impact on the machine learning model traning pipeline or the runtime environment.
- `exceptions/`
  - Standalone folder that allows for the creation of custom exceptions that external APIs may throw and other custom execptions that may need to be declared.
- `logic/`
  - Standalone folde that allows users to declare in a JSON format what features they would like to train the ML models on and also informs the runtime environment which technicals to recompute in real time to advise the machien learning model.
- `ml/`
  - Contains the entire process from constructing the Pandas DataFrame to trianing and exporting the machine learning model that can then be exposed to an API server. Trained and predicts buy/sell signals from technicals indicators declared in the `logic/` folder. 
- `runtime/`
  - Contains the runtime environment that allows the main Go runtime to talk to the machine learning API and compute technical in real time and also provide info the expose the the `bot_api` which is the basis of the monitoring dashboard.
