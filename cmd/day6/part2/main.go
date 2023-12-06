package main

import (
	"fmt"
	"github.com/Static-Flow/adventOfCode2023/cmd/day3"
	"github.com/Static-Flow/adventOfCode2023/internal"
	"time"
)

var race [2]int

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
					race[0] = storedNum
					storedNum = 0
				}
			}
		} else {
			break
		}
	}
	race[1] = storedNum
	var winCount = 0
	now := time.Now()
	var endIndex int
	var index = 1
	endIndex = race[0] - 1
	for {
		// fmt.Println("Checking:", index, endIndex)
		if index*(race[0]-index) > race[1] {
			// fmt.Println("Winner: ", index, index*(race[0]-index))
			if winCount == 0 {
				winCount = index
			} else {
				winCount = winCount - index
				break
			}
		} else {
			index++
		}
		if endIndex*(race[0]-endIndex) > race[1] {
			// fmt.Println("Winner: ", endIndex, endIndex*(race[0]-endIndex))

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
	fmt.Println("final win count:", winCount+1)
	fmt.Println(time.Since(now))

}
