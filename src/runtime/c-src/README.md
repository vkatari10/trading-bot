# `csrc/`

This file contains C source code used at runtime in the `go-src/` folder to optimize heavy array math computations.

## Purpose

The reason why we offload some operations to C is simply because of lower overhead, especially when compared to Go and Python since we can get closer to the metal and avoid potential disruptions from garbage collectors.<br>

## Structure

- Source files are contained in `src/`
- Header files are contianed in `include/`
- Compilation script to precompile C code into a library for Go to call
