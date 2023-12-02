# Day 1 Part 1:

## Sample Input:
```
nqninenmvnpsz874
8twofpmpxkvvdnpdnlpkhseven4ncgkb
six8shdkdcdgseven8xczqrnnmthreecckfive
qlcnz54dd75nine7jfnlfgz
7vrdhggdkqbnltlgpkkvsdxn2mfpghkntzrhtjgtxr
cdhmktwo6kjqbprvfour8
```

## Algorithm Overview
As all "Advents of Code" this starts off fairly simple. Walk each rune of the string and do the following:

A. If it's a number, and we haven't seen a number before, store it as the first number.

B. If it's a number, and we _have_ seen a number before, store it as the second number, replacing the old second number.

Once every character is looped over, return the first number multiplied by 11 if there is no second number or the first number multiplied by 10 plus the second number.

## Performance:
```
Name                               iterations        CPU op/ns      MEM use per op      MEM allocs per op
BenchmarkProcessLine-8   	   34329	     35213 ns/op    0 B/op	        0 allocs/op
```