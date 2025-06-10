# Version History

This gives an overview of all tagged version of the repo and associated changes.

## [v0.1.0-alpha]
- MVP proof of concept with no regard to performance.
- Prototyped fully in Python.
- Used Pandas directly to set up the data pipeline.
- No server implementation.

## [v0.1.1-alpha] 
- Introduce `pybind11` to speed up runtime technical computations using C++
- Still no server implmentation.

## [v0.2.0]
- Restructure entire back end architecture
- Remove `pybind11` integration, use `Go` exclusively
- Modularize ML training, runtime engine on JSONs in `src/logic`
- Migrate runtime environment to `Go`
- Develop `Flask` API to expose ML model to `Go` Runtime ("Server" implementation)
- Introduce `src/runtime/c-src` folder for future C-based improvements that can be called in `Go`

## [WIP]
In Progress<br>
- Improve features to include technical deltas and differences
- Develop internal API to expose the runtime environment 
- Construct monitoring dashboard for the entire engine
