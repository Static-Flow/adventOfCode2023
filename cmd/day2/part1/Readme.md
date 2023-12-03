# Day 2 Part 1:

## Sample Input:
```
Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
```

## Algorithm Overview
<details>
  <summary>Spoiler warning</summary>

  For Day 2 Part 1 a fairly tight O(n) loop can be built by performing the following for each character:
        
        A. If the character is a digit store it
        B. If the character is a digit and the previous character *was* too:
            B.a Build a 2 digit number from the two digits
            B.b If the 2 digit number is > 12:
                B.b.a if the character 2 steps ahead is 'r':
                    This is an invalid game
                B.b.b if the character 2 steps ahead is 'g' and the 2 digit number is > 13:
                    This is an invalid game
                B.b.c if the character 2 steps ahead is 'g' and the 2 digit number is > 14:
                    This is an invalid game
        C. If the character is *not* a number, reset the stored digit to 0
</details>

## Performance (AVG. Per Line):
```
Name                      CPU   ns/op      # bytes alloc'd per op      # of allocs per op
BenchmarkProcessGame-8    60.53	op/ns	   0 B/op	               0 allocs/op
```