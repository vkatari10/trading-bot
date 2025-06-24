# Version History
Highlights of all tagged versions of the repo and associated changes.

## [v0.1.0-alpha]
Goal
- Create basic MVP  

Changes
- Initial full stack version prototyped entirely in Python
- Hardcoded ML Pipeline using Scikit-learn
- Basic eventloop for runtime

## [v0.1.1-alpha]
Goal
- Improve runtime performance using a lower level language  

Changes
- Introduced `pybind11` to speed up runtime computations using `C++`

## [v0.2.0]
Goal
- Focus on modularity and remove most hardcoded methods  

Changes
- Restructured entire beack end architecture
- Remove `pybind11` integration
- Used JSON config files to replace hardcoded ML pipeline
- Migrate runtime environment to `Go`
- Develop a `Flask` based API server to expose trained ML models
- Introduce `src/runtime/c-src` folder to implement future C methods inside the `Go` runtime

## [v0.2.1]
Goal 
- Improve available features and create a demo frontend   

Changes
- Improved ML features to include `Delta` and  `Diff` of technical indicators, supported by JSON configs
- Exposed runtime engine via with an API
- Implemented a `Next` based frontend MVP using runtime engine API and `Typescript` API routes
- Refactored `src/runtime/go-src` to improve modularity of `Go` packages  

## [v0.2.2]
Goal
- Increase modularity via `.env` file settings, improve documentation, and add tests

Changes
- Condensed all setings into a single `env` file, increased number of settings available
- Improved frontend stying
- Improved documentation overall in `docs/` and main `README`
- Added unit tests for `Python` and `Go`
- Cleaned repo root with new folders: `scripts/`, `configs/`

In Progress
- Continue refactoring `src/runtime/go-src`, especially using generic methods to avoid repetitive code
- Add Custom backtesting engine in `Python`
- Finish `Docker` contanierization 
    