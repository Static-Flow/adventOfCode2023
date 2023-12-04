package main

import (
	"adventOfCode2023/cmd/day3"
	"adventOfCode2023/internal"
	"fmt"
)

func processEngine() int {
	var character byte
	var storedNumber int
	var sum int
	var engineNumber bool
	for {
		if character = day3.WalkableReader.ReadOrMoveDirection(true, day3.EAST); character != 0 {
			if internal.IsRuneANumber(rune(character)) {
				if storedNumber == 0 {
					storedNumber = int(character - '0')
					if day3.SearchNumberNeighbours(day3.CARDINAL_DIRECTIONS) {
						engineNumber = true
					}
				} else {
					storedNumber = (storedNumber * 10) + int(character-'0')
					if engineNumber == false {
						if day3.SearchNumberNeighbours(day3.QUICK_CARDINAL_DIRECTIONS) {
							engineNumber = true
						}
					}
				}
			} else {
				if storedNumber != 0 {
					if engineNumber {
						sum += storedNumber
					}
					engineNumber = false
					storedNumber = 0
				}
			}
		} else {
			break
		}
	}
	return sum
}

func main() {
	day3.WalkableReader = day3.NewWalkableByteStream("../input")
	fmt.Println(processEngine())
}
