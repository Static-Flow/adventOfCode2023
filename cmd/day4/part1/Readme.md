# Day 4 Part 1:

## Sample Input:
```
Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
```

## Algorithm Overview
<details>
  <summary>Spoiler warning</summary>

    For day 4 it's about finding the count of winning numbers in a card (one per line)
    
    To do this:
        First build a map of all possible winning numbers
        Then, while parsing winning numbers, set them to true in the map
        When a '|' is encountered, switch to matching and for every number found, check if its map value is true
            If it is, add the current multiplier to the running sum and double the multiplier
        When you reach a newline character reset the card state and continue
</details>

## Performance (AVG. Per Line):
```
Run Time [ using now:=time.Now() //do_work print(time.Since(now) ]: 585.6µs

Name                        CPU    ns/op      # bytes alloc'd per op    # of allocs per op
BenchmarkProcessCards-8     259682 ns/op      0 B/op                    0 allocs/op

```