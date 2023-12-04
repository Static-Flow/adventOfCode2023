# Day 4 Part 2:

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

    For part 2 we have to find total count of winning cards and the number of matched numbers on each card
    adds that many copies of future cards.

    To build the deck of scratcher we use the same procedure as part 1 but instead of adding up matches we record how many.
    Once you have the deck of scratchers loop over them and do the following:
        add the count of times we've won this card to the total
        if card FOO has 1 or more matches:
            starting from the next card BAR, step len(matches) number of times:
                add 1 * FOO.won_count to BAR.won_count
</details>

## Performance (AVG. Per Line):
```
Run Time [ using now:=time.Now() //do_work print(time.Since(now) ]: 519.1µs

Benchmark:  
Name                        CPU    ns/op      # bytes alloc'd per op    # of allocs per op
BenchmarkProcessCards-8     1263   ns/op      0 B/op	                0 allocs/op
```