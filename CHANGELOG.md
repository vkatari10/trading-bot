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

## [WIP]
- Restructures the entire back end architecture
- Remove `pybind11` integration
- Recreate `pybind11` functions in C
- Migrate runtime environment to Go
- Develop a simple API to expose the ML model
