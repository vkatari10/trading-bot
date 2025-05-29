# Logic

This folder cotains the user defined logic to determine the technical indicators to train the ML models on and the real time indicators needed to compute in real time.<br>

## Usage

In `features.json` a single array should be loaded with the technical indicators that the ML model should use (note that a lot of technical indicators have not been implemented yet) and also allows the runtime engine to determine which indicators to compute in real time.<br>

At this point in time the user will need to manually enter in the relationships into a method to determine buy/sell signals until another JSON format can be adapted to fit this structure. However the current iteration makes the training and live computations far more modular.<br> 
## Supported Technical Indicators

The current list of supported technicals arguments are provided below as well as the attributes that come along with the technical indicators.<br>

Note: You will need to declare your own relationships in the `src/machine_learning/pipeline.py` to develop the actual buy/sell signals.<br>

- SMA
  - "tech": "SMA"
  - "window": int
- EMA
  - "tech": "EMA"
  - "window": int
