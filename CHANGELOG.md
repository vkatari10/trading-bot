# Version History

This gives an overview of all tagged version of the repo and associated changes.

## [v0.1.0-alpha]
- MVP version with zero regard to performance
- Prototyped fully in Python.
- Used Pandas directly to set up the data pipeline, sci-kit learn for ML
- No server implementation.

## [v0.1.1-alpha] 
- Introduced `pybind11` to speed up runtime technical computations using C++
- Still no server implmentation.

## [v0.2.0]
- Restructure entire back end architecture
- Remove `pybind11` integration, use `Go` exclusively
- Modularize ML training, runtime engine on JSONs in `src/logic`
- Migrate runtime environment to `Go`
- Develop `Flask` API to expose ML model to `Go` Runtime
- Introduce `src/runtime/c-src` folder for future C-based improvements that can be called in `Go`

## [v0.2.1]
- Improved ML features including new `Delta` and `Diff` objects
- Developed internal runtime engine API in `Go`
- Implemented `Next` based frontend MVP implementation using `TypeScript`
- Refactored `src/runtime/go-src` to improve modularity

## [WIP]
In Progress<br>
- Clean frontend styling, utility
- Continue refactoring `src/runtime/go-src` (Introduce generic methods)
- Add Custom backtesting engine