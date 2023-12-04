package main

import (
	"adventOfCode2023/cmd/day3"
	"adventOfCode2023/internal"
	"fmt"
)

var value byte
var results = [2]int{}
var gearSum int
var gearNums int
var res int

func processNorthOrSouthOfGear(direction day3.Direction) {
	results[0], results[1] = 0, 0
	if value = day3.WalkableReader.ReadOrMoveDirection(false, direction); internal.IsRuneANumber(rune(value)) {
		results[0] = int(value - '0')
		if value = day3.WalkableReader.ReadAtLocation(day3.WalkableReader.LastSearchLocation() - 1); internal.IsRuneANumber(rune(value)) {
			results[0] = (int(value-'0') * 10) + results[0]
			if value = day3.WalkableReader.ReadAtLocation(day3.WalkableReader.LastSearchLocation() - 2); internal.IsRuneANumber(rune(value)) {
				results[0] = (int(value-'0') * 100) + results[0]
			}
		}
		if value = day3.WalkableReader.ReadAtLocation(day3.WalkableReader.LastSearchLocation() + 1); internal.IsRuneANumber(rune(value)) {
			results[0] = (results[0] * 10) + int(value-'0')
			if value = day3.WalkableReader.ReadAtLocation(day3.WalkableReader.LastSearchLocation() + 2); internal.IsRuneANumber(rune(value)) {
				results[0] = (results[0] * 10) + int(value-'0')
			}
		}
	} else {
		if value = day3.WalkableReader.ReadOrMoveDirection(false, direction+1); internal.IsRuneANumber(rune(value)) {
			results[0] = int(value - '0')
			if value = day3.WalkableReader.ReadAtLocation(day3.WalkableReader.LastSearchLocation() - 1); internal.IsRuneANumber(rune(value)) {
				results[0] = (int(value-'0') * 10) + results[0]
				if value = day3.WalkableReader.ReadAtLocation(day3.WalkableReader.LastSearchLocation() - 2); internal.IsRuneANumber(rune(value)) {
					results[0] = (int(value-'0') * 100) + results[0]
				}
			}
		}
		if value = day3.WalkableReader.ReadOrMoveDirection(false, direction+2); internal.IsRuneANumber(rune(value)) {
			results[1] = int(value - '0')
			if value = day3.WalkableReader.ReadAtLocation(day3.WalkableReader.LastSearchLocation() + 1); internal.IsRuneANumber(rune(value)) {
				results[1] = (results[1] * 10) + int(value-'0')
				if value = day3.WalkableReader.ReadAtLocation(day3.WalkableReader.LastSearchLocation() + 2); internal.IsRuneANumber(rune(value)) {
					results[1] = (results[1] * 10) + int(value-'0')
				}
			}
		}
	}
}

func processEastOfGear() {
	results[0] = 0
	if value = day3.WalkableReader.ReadOrMoveDirection(false, day3.EAST); internal.IsRuneANumber(rune(value)) {
		results[0] = int(value - '0')
		if value = day3.WalkableReader.ReadAtLocation(day3.WalkableReader.LastSearchLocation() + 1); internal.IsRuneANumber(rune(value)) {
			results[0] = (results[0] * 10) + int(value-'0')
			if value = day3.WalkableReader.ReadAtLocation(day3.WalkableReader.LastSearchLocation() + 2); internal.IsRuneANumber(rune(value)) {
				results[0] = (results[0] * 10) + int(value-'0')
			}
		}
	}
}

func processWestOfGear() {
	results[0] = 0
	if value = day3.WalkableReader.ReadOrMoveDirection(false, day3.WEST); internal.IsRuneANumber(rune(value)) {
		results[0] = int(value - '0')
		if value = day3.WalkableReader.ReadAtLocation(day3.WalkableReader.LastSearchLocation() - 1); internal.IsRuneANumber(rune(value)) {
			results[0] = (int(value-'0') * 10) + results[0]
			if value = day3.WalkableReader.ReadAtLocation(day3.WalkableReader.LastSearchLocation() - 2); internal.IsRuneANumber(rune(value)) {
				results[0] = (int(value-'0') * 100) + results[0]
			}
		}
	}
}

func processGear() {
	gearSum = 1
	gearNums = 0
	processNorthOrSouthOfGear(day3.NORTH)
	for _, res = range results {
		if res != 0 {
			gearSum *= res
			gearNums++
		}
	}
	processNorthOrSouthOfGear(day3.SOUTH)
	for _, res = range results {
		if res != 0 {
			gearSum *= res
			gearNums++
		}
	}
	processEastOfGear()
	if results[0] > 0 {
		gearSum *= results[0]
		gearNums++
	}
	processWestOfGear()
	if results[0] > 0 {
		gearSum *= results[0]
		gearNums++
	}
	if gearNums != 2 {
		gearSum = 0
	}
}

func processEngine() int {
	var character byte
	var sum int
	// fmt.Println(day3.WalkableReader)
	for {
		if character = day3.WalkableReader.ReadOrMoveDirection(true, day3.EAST); character != 0 {
			if day3.IsEngineGear(character) {
				// check gear
				processGear()
				sum += gearSum
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
