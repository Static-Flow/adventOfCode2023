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
Performance (AVG. Per Line):
```
Name                        CPU    ns/op      # bytes alloc'd per op    # of allocs per op
BenchmarkProcessEngine-8    103517 ns/op      0 B/op	                0 allocs/op
```

#### Part 2: [Link](cmd/day3/part2)
Performance (AVG. Per Line):
```
Name                        CPU     ns/op      # bytes alloc'd per op   # of allocs per op
BenchmarkProcessEngine-8    68783   ns/op      0 B/op	                0 allocs/opp
```
</details>