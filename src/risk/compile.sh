#! /usr/bin/bash
gcc -c -O2 -fPIC c_src/*
gcc -shared -o librisk.so *.o 
rm -f *.o