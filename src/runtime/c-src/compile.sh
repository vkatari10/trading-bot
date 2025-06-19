#! /usr/bin/bash

# This will compile C source code into obj files and 
# compress it into a single library since Go handles
# linking
gcc -c -O2 src/*.c
ar rcs liblive_data.a *.o   