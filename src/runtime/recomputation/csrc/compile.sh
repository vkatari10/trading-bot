#! /usr/bin/bash
gcc -c src/*.c
ar rcs liblive_data.a *.o
