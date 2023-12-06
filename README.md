# adventOfCode2023
Advent of Code 2023 Solutions in Golang

Each day is it's own folder under cmd so they're all indivdually buildable. Solution is spoilered to avoid ruining it while wanting to look at just the performance results.

### Day 1
<details>

#### Problem Text: [Link](cmd/day1)

#### Part 1: [Link](cmd/day1/part1)
Performance (AVG. Per Line):
```
Name                      CPU   ns/op      # bytes alloc'd per op      # of allocs per op
BenchmarkProcessLine-8    14.27 ns/op      0 B/op	               0 allocs/op
```

#### Part 2: [Link](cmd/day1/part2)
Performance (AVG. Per Line):
```
Name                      CPU   ns/op      # bytes alloc'd per op      # of allocs per op
BenchmarkProcessLine-8    122.6 ns/op      0 B/op	               0 allocs/op
```
</details>

### Day 2
<details>

#### Problem Text: [Link](cmd/day2)

#### Part 1: [Link](cmd/day2/part1)
Performance (AVG. Per Line):
```
Name                      CPU   ns/op      # bytes alloc'd per op      # of allocs per op
BenchmarkProcessGame-8    60.53	ns/op	   0 B/op	               0 allocs/op
```

#### Part 2: [Link](cmd/day2/part2)
Performance (AVG. Per Line):
```
Name                      CPU     ns/op      # bytes alloc'd per op      # of allocs per op
BenchmarkProcessGame-8    138.33  ns/op      0 B/op	                 0 allocs/op
```
</details>

### Day 3
<details>

#### Problem Text: [Link](cmd/day3)

#### Part 1: [Link](cmd/day3/part1)
Performance:
```
Name                        CPU    ns/op      # bytes alloc'd per op    # of allocs per op
BenchmarkProcessEngine-8    103517 ns/op      0 B/op	                0 allocs/op
```

#### Part 2: [Link](cmd/day3/part2)
Performance:
```
Name                        CPU     ns/op      # bytes alloc'd per op   # of allocs per op
BenchmarkProcessEngine-8    68783   ns/op      0 B/op	                0 allocs/opp
```
</details>

### Day 4
<details>

#### Problem Text: [Link](cmd/day4)

#### Part 1: [Link](cmd/day4/part1)
Performance (AVG. Per Line):
```
Run Time [ using now:=time.Now() //do_work print(time.Since(now) ]: 25µs

Name                        CPU    ns/op      # bytes alloc'd per op    # of allocs per op
BenchmarkProcessCards-8     2295   ns/op      0 B/op                    0 allocs/op
```

#### Part 2: [Link](cmd/day4/part2)
Performance (AVG. Per Line):
```
Run Time [ using now:=time.Now() //do_work print(time.Since(now) ]: 20µs

Benchmark:  
Name                        CPU    ns/op      # bytes alloc'd per op    # of allocs per op
BenchmarkProcessCards-8     1263   ns/op      0 B/op	                0 allocs/op
```
</details>

### Day 5
<details>

#### Problem Text: [Link](cmd/day5)

#### Part 1: [Link](cmd/day5/part1)
Performance (AVG. Per Line):
```
Run Time [ using now:=time.Now() //do_work print(time.Since(now) ]: 74µs
```

#### Part 2: [Link](cmd/day5/part2)
Performance (AVG. Per Line):
```
Run Time [ using now:=time.Now() //do_work print(time.Since(now) ]: 7 minutes....
```
</details>

### Day 6
<details>

#### Problem Text: [Link](cmd/day6)

#### Part 1: [Link](cmd/day6/part1)
Performance (AVG. Per Line):
```
Run Time [ using now:=time.Now() //do_work print(time.Since(now) ]: 20µs
```

#### Part 2: [Link](cmd/day6/part2)
Performance (AVG. Per Line):
```
Run Time [ using now:=time.Now() //do_work print(time.Since(now) ]: 4 ms
```
</details>