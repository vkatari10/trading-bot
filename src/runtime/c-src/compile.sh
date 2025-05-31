#! /usr/bin/bash
gcc -c -O2 src/*.c
ar rcs liblive_data.a *.o