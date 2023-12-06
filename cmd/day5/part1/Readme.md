# Day 5 Part 1:

## Sample Input:
```
seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
```

## Algorithm Overview
<details>
  <summary>Spoiler warning</summary>

    For day 5 it's about finding the lowest location through the map of X-to-Y
    
    To do this:
        for every seed:
            loop through the map ranges and find a matching range and shift to that value
            once we reach the end of the map check if the location is less than the lowest and replace if so
        return lowest
</details>

## Performance (AVG. Per Line):
```
Run Time [ using now:=time.Now() //do_work print(time.Since(now) ]: 74.3µs
```