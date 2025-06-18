#! /usr/bin/bash
# compilation script for C dependencies
gcc -c -O2 src/*.c
ar rcs liblive_data.a *.o