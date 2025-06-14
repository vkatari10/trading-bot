# Version History
Highlights of all tagged versions of the repo and associated changes.

## [v0.1.0-alpha]

The goal of this version was to get a working prototype of the project

- Initial full stack version prototyped entirely in Python

- Hardcoded ML Pipeline using Scikit-learn

- Basic eventloop for runtime


## [v0.1.1-alpha]

The goal of this version was to improve some of the runtime performance by moving some runtime methods to lower level languages

- Introduced `pybind11` to speed up runtime computations using `C++`

  

## [v0.2.0]

The goal of this version was to completely change the makeup of this project to include a higher distinction between the ML training pipeline and focus heavily on modularity. Since a lot of the ML pipeline was hardcoded I decided to use JSON config files to give a cleaner interface to users. Also I wanted to used another language outside of `Python` with lower latency. By decoupling both the ML and runtime languages I could achieve a better separation of concerns

- Restructured entire back end architecture

- Remove `pybind11` integration

- Used JSON config files to replace hardcoded ML pipeline

- Migrate runtime environment to `Go`

- Develop a `Flask` based API server to expose trained ML models

- Introduce `src/runtime/c-src` folder to implement future C methods inside the `Go` runtime

  

## [v0.2.1]

The goal of this version was to improve the features given to the ML models including deltas and differences rather than just raw technical indicator values. I also wanted to develop an internal dashboard to monitor the runtime engine using its own API.

- Improved ML features to include `Delta` and  `Diff` of technical indicators, supported by JSON configs

- Exposed runtime engine via with an API

- Implemented a `Next` based frontend MVP using runtime engine API and `Typescript` API routes

- Refactored `src/runtime/go-src` to improve modularity of `Go` packages  

  

## [WIP]

In Progress<br>

- Refine frontend styling

- Continue refactoring `src/runtime/go-src`, especially using generic methods to avoid repetitive code

- Add Custom back testing engine in `Python`