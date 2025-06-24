# `c-src/` - C Source Code

This folder intended to contain `C` code that could be called from `Go` to enhance runtime performance, but due to extra complexity of managing `C` arrays in `Go` there is no current implementation.  

If you would like to use `C` in `Go`, only create methods that return non-heap values (doubles, ints, etc.) and use `import C` in `Go`. You can use the provided compilation script to compile your `C` code so `Go` can link and execute it.  