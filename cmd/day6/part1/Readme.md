# Day 6 Part 1:

## Sample Input:
```
Time:      7  15   30
Distance:  9  40  200
```

## Algorithm Overview
<details>
  <summary>Spoiler warning</summary>

    For day 6 it's about finding the number of possible wins between a race time and distance
    
    Part 1 is simple to find by starting at the ends and working inwards, stopping the search once both ends are found
</details>

## Performance (AVG. Per Line):
```
Run Time [ using now:=time.Now() //do_work print(time.Since(now) ]: 20.5µs
```