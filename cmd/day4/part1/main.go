package main

import (
	"fmt"
	"github.com/Static-Flow/adventOfCode2023/cmd/day3"
	"github.com/Static-Flow/adventOfCode2023/internal"
	"time"
)

var character byte
var winningNumbers = [10]int{}
var candidatesMap = [100]bool{}
var multiplier = 0
var storedDigit int
var index int
var checkingWinners bool
var sum int

func checkWinner() {
	if candidatesMap[storedDigit] {
		if multiplier == 0 {
			multiplier = 1
		} else {
			multiplier *= 2
		}
	}
}

func processCharacter() {
	if internal.IsRuneANumber(rune(character)) {
		// found a number
		if storedDigit == 0 {
			storedDigit = int(character - '0')
		} else {
			storedDigit = storedDigit*10 + int(character-'0')
		}
	} else {
		if storedDigit != 0 {
			if !checkingWinners {
				winningNumbers[index] = storedDigit
				candidatesMap[storedDigit] = true
				index++
			} else {
				checkWinner()
			}
			storedDigit = 0
		} else if character == 124 {
			// vertical bar
			checkingWinners = true
		} else if character == '\r' || character == '\n' {
			for i := 0; i < index; i++ {
				candidatesMap[winningNumbers[i]] = false
			}

			day3.WalkableReader.SetLocation(day3.WalkableReader.GetLocation() + day3.OFFSET + 7)
			checkingWinners = false
			sum += multiplier
			multiplier = 0
			index = 0
		}
	}
}

func processCards() {
	for {
		if character = day3.WalkableReader.ReadOrMoveDirection(true, day3.EAST); character != 0 {
			processCharacter()
		} else {
			break
		}
	}
	checkWinner()
	sum += multiplier
	multiplier = 0
	index = 0
}

func main() {
	day3.WalkableReader = day3.NewWalkableByteStream("../input")
	day3.WalkableReader.SetLocation(7)
	now := time.Now()
	processCards()
	fmt.Println(sum)
	fmt.Println(time.Since(now))
}
