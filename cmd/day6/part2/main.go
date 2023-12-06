package main

import (
	"fmt"
	"github.com/Static-Flow/adventOfCode2023/cmd/day3"
	"github.com/Static-Flow/adventOfCode2023/internal"
	"math"
	"time"
)

var raceTime, distance float64

/*
The reddit way which _is_ much faster but for the sake of honesty my version is below
*/
func quadratic() float64 {
	halfTime := raceTime / 2
	root := math.Pow(halfTime, 2)
	sqrRoot := math.Sqrt(root - distance)
	return (math.Floor(halfTime + sqrRoot)) - (math.Ceil(halfTime - sqrRoot)) + 1
}

/*
*
Start from the ends and work towards the middle till you find both the min and max win
*/
func search() (winCount float64) {
	var endIndex float64
	var index = float64(1)
	endIndex = raceTime - 1
	for {
		if index*(raceTime-index) > distance {
			if winCount == 0 {
				winCount = index
			} else {
				winCount = winCount - index
				break
			}
		} else {
			index++
		}
		if endIndex*(raceTime-endIndex) > distance {
			if winCount == 0 {
				winCount = endIndex
			} else {
				winCount = endIndex - winCount
				break
			}
		} else {
			endIndex--
		}
	}
	winCount++
	return
}

func main() {
	walker := day3.NewWalkableByteStream("../input")
	var storedNum int
	for {
		if character := walker.ReadOrMoveDirection(true, day3.EAST); character != 0 {
			if internal.IsRuneANumber(rune(character)) {
				if storedNum == 0 {
					storedNum = int(character - '0')
				} else {
					storedNum = storedNum*10 + int(character-'0')
				}
			} else {
				if storedNum != 0 && (character == '\r' || character == '\n') {
					raceTime = float64(storedNum)
					storedNum = 0
				}
			}
		} else {
			break
		}
	}
	distance = float64(storedNum)
	now := time.Now()

	fmt.Println("final win count:", search())
	fmt.Println(time.Since(now))

}
