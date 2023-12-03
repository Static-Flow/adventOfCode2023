# Day 2 Part 2:

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

  For Day 2 Part 2 is much like part 1 except instead of checking if the last number seen is too high, the check is done to see if it's the highest for 
    the related color. 

    For each character in line
        A. If the character is a digit store it
        B. If the character is a digit and the previous character *was* too, build a 2 digit number from the two digits
        C. If the character is *not* a number:
            C.a if the character is 'r':
                C.a.a if the highest number associated with 'red' is lower than the stored number:
                    replace the previously higher number with it
                C.a.b reset the stored digit to 0
            C.b if the character is 'g':
                C.b.a if the highest number associated with 'green' is lower than the stored number:
                    replace the previously higher number with it
                C.b.b reset the stored digit to 0
            C.c if the character is 'b':
                C.c.a if the highest number associated with 'blue' is lower than the stored number:
                    replace the previously higher number with it
                C.c.b reset the stored digit to 0
    return highest_red * highest_green * highest_blue 
    

</details>

## Performance (AVG. Per Line):
```
Name                      CPU    ns/op      # bytes alloc'd per op      # of allocs per op
BenchmarkProcessGame-8    138.33 ns/op	    0 B/op	                0 allocs/op
```