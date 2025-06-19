# `c-src/` - C Source Code

This folder was originally intended to contain `C` code that could offload heavy operations from `Go`. Due to the extra complexity of handling `C` arrays in `Go`, there are no current implementations.<br>

However in the future if we need to add `C` source code to move heavy numerical operations away from `Go` this folder makes it easy to implement and integrate into the `go-src` folder.<br>