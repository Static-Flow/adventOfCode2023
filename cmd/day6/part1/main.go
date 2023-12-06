package main

import (
	"fmt"
	"github.com/Static-Flow/adventOfCode2023/cmd/day3"
	"github.com/Static-Flow/adventOfCode2023/internal"
	"time"
)

var races [][2]int

func main() {
	walker := day3.NewWalkableByteStream("../input")
	var storedNum int
	var storingDistances bool
	var index int
	for {
		if character := walker.ReadOrMoveDirection(true, day3.EAST); character != 0 {
			if internal.IsRuneANumber(rune(character)) {
				if storedNum == 0 {
					storedNum = int(character - '0')
				} else {
					storedNum = storedNum*10 + int(character-'0')
				}
			} else {
				if storedNum != 0 {
					if !storingDistances {
						races = append(races, [2]int{storedNum})
					} else {
						races[index][1] = storedNum
						index++
					}
					storedNum = 0

				}
				if character == '\r' || character == '\n' {
					storingDistances = true
				}

			}
		} else {
			break
		}
	}
	races[index][1] = storedNum

	now := time.Now()
	var winCount int
	var sum = 1
	var endIndex int
	for _, race := range races {
		winCount = 0
		index = 1
		endIndex = race[0] - 1
		for {
			if index*(race[0]-index) > race[1] {
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
		sum *= winCount + 1
	}
	fmt.Println(sum)
	fmt.Println(time.Since(now))
}
