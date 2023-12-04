package main

import (
	"adventOfCode2023/cmd/day3"
	"adventOfCode2023/internal"
	"fmt"
	"math"
)

var storedNum int
var index int
var value byte
var results = [2]int{}

func processNorthOfGear() int {
	storedNum = 0
	index = 1
	results[0], results[1] = 0, 0
	if value = day3.WalkableReader.ReadOrMoveDirection(false, day3.NORTH); internal.IsRuneANumber(rune(value)) {
		// north is a number
		storedNum = int(value - '0')
		for {
			if value = day3.WalkableReader.ReadAtLocation(day3.WalkableReader.LastSearchLocation() - index); internal.IsRuneANumber(rune(value)) {
				storedNum = (int(value-'0') * int(math.Pow10(index))) + storedNum
				index++
			} else {
				break
			}
		}
		index = 1
		for {
			if value = day3.WalkableReader.ReadAtLocation(day3.WalkableReader.LastSearchLocation() + index); internal.IsRuneANumber(rune(value)) {
				storedNum = (storedNum * 10) + int(value-'0')
				index++
			} else {
				break
			}
		}
		results[0] = storedNum
	} else {
		if value = day3.WalkableReader.ReadOrMoveDirection(false, day3.NORTH_WEST); internal.IsRuneANumber(rune(value)) {
			// north west is a number
			index = 1
			storedNum = int(value - '0')
			for {
				if value = day3.WalkableReader.ReadAtLocation(day3.WalkableReader.LastSearchLocation() - index); internal.IsRuneANumber(rune(value)) {
					storedNum = (int(value-'0') * int(math.Pow10(index))) + storedNum
					index++
				} else {
					break
				}
			}
			results[0] = storedNum
		}
		if value = day3.WalkableReader.ReadOrMoveDirection(false, day3.NORTH_EAST); internal.IsRuneANumber(rune(value)) {
			// north east is a number
			index = 1
			storedNum = int(value - '0')
			for {
				if value = day3.WalkableReader.ReadAtLocation(day3.WalkableReader.LastSearchLocation() + index); internal.IsRuneANumber(rune(value)) {
					storedNum = (storedNum * 10) + int(value-'0')
					index++
				} else {
					break
				}
			}
			results[1] = storedNum
		}
	}
	// fmt.Println("north", storedNum)
	return storedNum
}

func processSouthOfGear() int {
	storedNum = 0
	index = 1
	results[0], results[1] = 0, 0
	if value = day3.WalkableReader.ReadOrMoveDirection(false, day3.SOUTH); internal.IsRuneANumber(rune(value)) {
		// south is a number
		storedNum = int(value - '0')
		for {
			if value = day3.WalkableReader.ReadAtLocation(day3.WalkableReader.LastSearchLocation() - index); internal.IsRuneANumber(rune(value)) {
				storedNum = (int(value-'0') * int(math.Pow10(index))) + storedNum
				index++
			} else {
				break
			}
		}
		index = 1
		for {
			if value = day3.WalkableReader.ReadAtLocation(day3.WalkableReader.LastSearchLocation() + index); internal.IsRuneANumber(rune(value)) {
				storedNum = (storedNum * 10) + int(value-'0')
				index++
			} else {
				break
			}
		}
		results[0] = storedNum
	} else {
		if value = day3.WalkableReader.ReadOrMoveDirection(false, day3.SOUTH_WEST); internal.IsRuneANumber(rune(value)) {
			// south west is a number
			index = 1
			storedNum = int(value - '0')
			for {
				if value = day3.WalkableReader.ReadAtLocation(day3.WalkableReader.LastSearchLocation() - index); internal.IsRuneANumber(rune(value)) {
					storedNum = (int(value-'0') * int(math.Pow10(index))) + storedNum
					index++
				} else {
					break
				}
			}
			results[0] = storedNum
		}
		if value = day3.WalkableReader.ReadOrMoveDirection(false, day3.SOUTH_EAST); internal.IsRuneANumber(rune(value)) {
			// south east is a number
			index = 1
			storedNum = int(value - '0')
			for {
				if value = day3.WalkableReader.ReadAtLocation(day3.WalkableReader.LastSearchLocation() + index); internal.IsRuneANumber(rune(value)) {
					storedNum = (storedNum * 10) + int(value-'0')
					index++
				} else {
					break
				}
			}
			results[1] = storedNum
		}
	}
	// fmt.Println("south", storedNum)
	return storedNum
}

func processEastOfGear() int {
	storedNum = 0
	if value = day3.WalkableReader.ReadOrMoveDirection(false, day3.EAST); internal.IsRuneANumber(rune(value)) {
		storedNum = int(value - '0')
		if value = day3.WalkableReader.ReadAtLocation(day3.WalkableReader.LastSearchLocation() + 1); internal.IsRuneANumber(rune(value)) {
			storedNum = (storedNum * 10) + int(value-'0')
			if value = day3.WalkableReader.ReadAtLocation(day3.WalkableReader.LastSearchLocation() + 2); internal.IsRuneANumber(rune(value)) {
				storedNum = (storedNum * 10) + int(value-'0')
			}
		}
	}
	// fmt.Println("east", storedNum)
	return storedNum
}

func processWestOfGear() int {
	storedNum = 0
	if value = day3.WalkableReader.ReadOrMoveDirection(false, day3.WEST); internal.IsRuneANumber(rune(value)) {
		storedNum = int(value - '0')
		if value = day3.WalkableReader.ReadAtLocation(day3.WalkableReader.LastSearchLocation() - 1); internal.IsRuneANumber(rune(value)) {
			storedNum = (int(value-'0') * 10) + storedNum
			if value = day3.WalkableReader.ReadAtLocation(day3.WalkableReader.LastSearchLocation() - 2); internal.IsRuneANumber(rune(value)) {
				storedNum = (int(value-'0') * 100) + storedNum
			}
		}
	}
	// fmt.Println("west", storedNum)
	return storedNum
}

func processGear() int {
	gearSum := 1
	gearNums := 0
	if northValue := processNorthOfGear(); northValue > 0 {
		for _, res := range results {
			if res != 0 {
				gearSum *= res
				gearNums++
			}
		}
	}
	if northValue := processSouthOfGear(); northValue > 0 {
		for _, res := range results {
			if res != 0 {
				gearSum *= res
				gearNums++
			}
		}
	}
	if northValue := processEastOfGear(); northValue > 0 {
		gearSum *= northValue
		gearNums++
	}
	if northValue := processWestOfGear(); northValue > 0 {
		gearSum *= northValue
		gearNums++
	}
	// fmt.Println(gearSum)
	if gearNums != 2 {
		return 0
	}
	return gearSum
}

func processEngine() int {
	var character byte
	var sum int
	// fmt.Println(day3.WalkableReader)
	for {
		if character = day3.WalkableReader.ReadOrMoveDirection(true, day3.EAST); character != 0 {
			if day3.IsEngineGear(character) {
				// check gear
				sum += processGear()
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
