# Architecture 

This file explains
- Implementation 
- Technology choices & justification 

## Config Files

Language:
- `JSON`

Implementation
- Contains 
    - Features
    - Labelling logic
    - ML model settings
    - Execution settings
    - Training tickers
    - Live trading tickers
- Acts as a Single Source of Truth (SSOT)
    - Preventing feature misalignment
    - Keeping every service uniform 
- All services when called by the user should be given the same file

Justification:
- Universally Recognized and is supported well in `Python` and `Go`
- Why not `YAML`?
    - `YAML` is loosely type compared to `JSON`
    - `JSON`s are easier to parse in `Python` and `Go`

Drawbacks
- Cannot comment in `JSON` files to take notes
    - Could create a section not found in `config/demo.json`
    - The program will never read additional sections added by the user outside of required sections


## ML Training Pipeline (V4)

Language:
- `Python`

Technologies:
- `Pandas`
- `NumPy`
- `Scikit-learn`
- `TA-lib`
- `Rich`

External API(s):
- `YFinance`

Implementation 
- Reads features, labelling logic, and ML settings
- Constructs a `Pandas DataFrame` using config file settings
- Trains and dumps the model 

Justification 
- `Python` 
    - Mature ecosystem for Data and ML libraries
- `Pandas`
    - More intuitive to use than 2D `NumPy` arrays
    - Although slower than `NumPy`, runtime speed at training time is not critical
- `Rich`
    - Used for improved UX on the CLI 
- `TA-lib` 
    - Mature library to compute technical indicators
    - Why write these methods again from scratch if this supports >150 technicals?
- `Scikit-learn` 
    - Offers a wide selection of models to users
    - Builds upon the modularity and flexibility of the platform
- `YFinance` 
    - Returns back as `DataFrames` for historical stock data
    - Reduces extra code for `DataFrame` conversions

Drawbacks
- `YFinance`
    - The lowest period of data we can get back is 1 minute
    - Meaning all ML models at the minimum can only be trained on 1 minute data
    - This puts an implied floor of 60 second cycles for the execution engine since 
    - More frequent predictions made below 60 seconds leads to potentially degraded performance

## Backtesting (V1)

Language:
- `Python`

Technologies
- `Pandas`
- `NumPy`
- `Scikit-learn`
- `TA-lib`
- `Rich`

External API(s):
- `YFinance`

Implementation 
- Borrows `DataFrame` fitting method from the `ML Pipeline`
- Walks a ML model down features, making predictions, and marking buy and sell signals
- Returns final P/L

Justification 
- See ML Pipeline
- Same `DataFrame` fitting method
    - By using the same method to fit `DataFrames` at train time and validation time ensures that a `DataFrame` is not made different from the ones it was trained on

Drawbacks
- Current implementation only shows 
    - Final P/L
    - This is fine for the MVP but for true performance analysis more data is needed

## Runtime Engine (V4)
Languages: 
- `Go`
- `Python`

Technologies:
- `FastAPI`
- `Gorilla WebSocket`
- `TA-Lib`

External APIs:
- `Alpaca`

Implementation 
- There two pieces used during live execution 
    - Runtime Engine
    - ML API 
- Runtime Engine
    - Using the config file, manages features for declared ticker(s)
    - Uses `TA-lib` for technical indicator initialization
- ML API 
    - Hosts the ML model and is parallelized to handle n amount of inferences at once 
    - 1:1 ratio of workers spawned per ticker on server startup
- Combined
    - Both communicate via websockets asynchronously 
    - `Go` send computed features to the server
    - ML inference is ran and returns an inference value
    - `Go` communicates with `Alpaca` to place orders 

Justification 
- Language Split
    - Concurrency is harder in `Python` with GIL, `Go` makes it performant compilation 
    - Splitting between `Python` and `Go` establishes clear separation of concerns 
    - `Scikit-learn` models cannot be easily used in `Go`
- `TA-lib`
    - Use established library to initialize technical indicator values
- `FastAPI` 
    - Allows for parallelized ML inference by dynamically adjust workers on the server to each ticker being tracker (1:1)
- `Gorilla Websocket`
    - Rather than use `HTTP REST` to communicate with the `FastAPI` server we use websockets
    - Since we constantly need to send and get updates from the server websockets reduce latency and connection overhead
`Alpaca`
    - Chosen as the exclusive broker due to ease of integration and extensive API support for trading

Drawbacks
- Language Split
    - Introduces more complexity as the project must maintain two different languages 
- `FastAPI` and `Gorilla WebSocket` 
    - Introduces latency to getting ML inferences
    - Can be overcome by non-`Scikit-learn` models by using `ONXX`
