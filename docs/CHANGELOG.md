# Version History
Highlights of all tagged versions of the repo and associated changes.

## [WIP]
Goal 

- Clean up Runtime Engine, Backtesting Services

Changes

- Releasing As: "ConTrade"
- Improves Backtesting CLI usage and UI/UX
- Replaces `Flask` ML API with `FastAPI` to support concurrent inference
- Utilizes `TA-lib` for all technical computations outside of `Delta` of `Diff`
- ML Pipeline overhauled by modularizing around `TA-lib`
- Implemented CI scripts for automated testing via `GitHub Actions`
- Added configurable model selection including `XGBoost` and `LightGBM` and 12 `Scikit-learn` models


In Progress

- Continue refactoring `src/runtime/go-src` (Go source module)
- Continue `Docker` containerization
    
## [v0.3.0] (07/02/2025)
Goal 

- Improve JSON configurability, overhaul current services, improve UI/UX

Changes

- Released As: "StratForge"
- Condense `JSON` configs into a single file, add more configuration features
- Refactored `Go`, ML Pipeline Modules. 
- Added MVP Backtesting Service 
- Updated ML Pipeline, Runtime Engine to handle `JSON` configs more elegantly
- Removed long `if-else` changes with new dispatch tables in ML Pipeline
- Configurable multi-asset training and trading
- Added CLI tool, TUI skeleton to improve UX/UI for services
- Removed `Next` demo frontend

## [v0.2.2] (06/24/2025)
Goal

- Increase modularity via `.env` file settings, improve documentation, and add tests

Changes

- Released As: "BotBuilder"
- Condensed all settings into a single `env` file, increased number of settings available
- Improved frontend styling
- Improved documentation overall in `docs/` and main `README`
- Added unit tests for `Python` and `Go`
- Cleaned repo root with new folders: `scripts/`, `configs/`

## [v0.2.1] (06/12/2025)
Goal 

- Improve available features and create a demo frontend   

Changes

- Released As: "Mid-Frequency Trading Bot"
- Improved ML features to include `Delta` and  `Diff` of technical indicators, supported by JSON configs
- Exposed runtime engine via with an API
- Implemented a `Next` based frontend MVP using runtime engine API and `Typescript` API routes
- Refactored `src/runtime/go-src` to improve modularity of `Go` packages  

## [v0.2.0] (06/04/2025)
Goal

- Focus on modularity and remove most hardcoded methods  

Changes

- Released as: "Trading Bot"
- Restructured entire back end architecture
- Remove `pybind11` integration
- Used JSON config files to replace hardcoded ML pipeline
- Migrate runtime environment to `Go`
- Develop a `Flask` based API server to expose trained ML models
- Introduce `src/runtime/c-src` folder to implement future C methods inside the `Go` runtime

## [v0.1.1-alpha] (05/24/2025)
Goal

- Improve runtime performance using a lower level language  

Changes

- Released as: "Trading Bot"
- Introduced `pybind11` to speed up runtime computations using `C++`

## [v0.1.0-alpha] (05/14/2025)
Goal

- Create basic MVP  

Changes

- Released As: "Trading Bot"
- Initial full stack version prototyped entirely in Python
- Hardcoded ML Pipeline using Scikit-learn
- Basic eventloop for runtime