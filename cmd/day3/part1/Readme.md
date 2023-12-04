# Day 3 Part 1:

## Sample Input:
```
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
```

## Algorithm Overview
<details>
  <summary>Spoiler warning</summary>

    For day 3 it's about finding rows of numbers where 1 (or more) of the digits touches a "gear" in one of the 8 cardinal directions.

    To do this, loop through each character and if it's a "gear": do the following:
        A. If the character "north" of the current character is a digit:
            A.a check both characters for digits to the left and right in case "north" is the middle of a number
        B. If the character "north-west" is a digit:
            B.a check two characters "west" for more of the number
        C. If the character "north-east" is a digit:
            C.a check two characters "east" for more of the number
        D. If the character "south" of the current character is a digit:
            D.a check both characters for digits to the left and right in case "south" is the middle of a number
        E. add any numbers found from A-D to the total sum
</details>

## Performance (AVG. Per Line):
```
Name                        CPU    ns/op      # bytes alloc'd per op    # of allocs per op
BenchmarkProcessEngine-8    103517 ns/op      0 B/op	                0 allocs/op
```