# Go Source Code

This folder contains the `Go` source code representing the main runtime evironment<br>
- `api/` contains both external and internal API handling
- `engine/` contains runtime logic for technical indicators 
- `eventloop/` combines both `api` and `engine` to produce the runtime environemnt 
- `tests/` contains testing methods for critical parts in `api/` and `engine/`

## Requirements 
- Current implementation depends on `Alpaca`, create an account to get your own API keys
- Included is a `.env.example` of what the `.env` file should look like, this module depends on it 
- **Rename `.env.example` to `.env`** after configuration, the module looks for this name

## Documentation 
- Documentation can be found [here](https://pkg.go.dev/github.com/vkatari10/trading-bot/src/runtime/go-src)