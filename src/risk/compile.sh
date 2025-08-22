#! /usr/bin/bash
gcc -c -O2 -fPIC c-src/* 
gcc -shared -o librisk.so *.o -lm
rm -f *.o