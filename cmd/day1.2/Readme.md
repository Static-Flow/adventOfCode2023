# Day 1 Part 2:

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
<details>
  <summary>Spoiler warning</summary>

Part 2 adds a wrinkle where textual digits i.e. ("one" -> 1) are valid as well.

The way I went about it was:

Looping over every character in the line:

    A. If the character is a number, handle it like Part 1, set it as the first or second number
    B. If the character is a letter, flow into a switch statement where each case is the first letter of the 9 digits
       then within each switch case look ahead the number of letter in each digit to see if there is a match. If so 
       set it as the first or second number.



</details>


## Performance (AVG. Per Line):
```
Name                         CPU   ns/op          # bytes alloc'd per op      # of allocs per op
BenchmarkProcessLine-8       122.6 ns/op          0 B/op	              0 allocs/op
```