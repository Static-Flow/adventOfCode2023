package main

import (
	"adventOfCode2023/internal"
	"fmt"
)

/*
	fmt.Printf("%+v\n", walkableFileReader)
	for i := 0; i < 16; i++ {
		fmt.Println(string(walkableFileReader.MoveEast()))
	}
	res, err := walkableFileReader.ReadNorth()
	fmt.Println(string(res), err)
	res, err = walkableFileReader.ReadWest()
	fmt.Println(string(res), err)
	res, err = walkableFileReader.ReadEast()
	fmt.Println(string(res), err)
	res, err = walkableFileReader.ReadSouth()
	fmt.Println(string(res), err)
	res, err = walkableFileReader.ReadNorthWest()
	fmt.Println(string(res), err)
	res, err = walkableFileReader.ReadNorthEast()
	fmt.Println(string(res), err)
	res, err = walkableFileReader.ReadSouthWest()
	fmt.Println(string(res), err)
	res, err = walkableFileReader.ReadSouthEast()
	fmt.Println(string(res), err)

	fmt.Println("moving north")
	fmt.Println(string(walkableFileReader.MoveNorth()))
	fmt.Println("moving west")
	fmt.Println(string(walkableFileReader.MoveWest()))
	fmt.Println("moving south")
	fmt.Println(string(walkableFileReader.MoveSouth()))
	fmt.Println("moving south")
	fmt.Println(string(walkableFileReader.MoveSouth()))
	fmt.Println("moving east")
	fmt.Println(string(walkableFileReader.MoveEast()))
*/

var walkableFileReader = internal.NewWalkableFileReader("../input")

func isEnginePart(character byte) bool {
	if character != 46 && (character >= 58 || (character >= 33 && character <= 47)) {
		// fmt.Println(string(character), " is an engine part")
		return true
	}
	return false
}

// searchNumberNeighbours searches each character in a cardinal direction counter-clockwise from the current character and if any are engine parts returns true
func searchNumberNeighbours(directions []internal.Direction) bool {
	var searchedCharacter byte
	var err error

	for _, DIRECTION := range directions {
		if searchedCharacter, err = walkableFileReader.ReadDirection(DIRECTION); err == nil && isEnginePart(searchedCharacter) {
			return true
		}
	}
	return false
}

func processEngine() int {
	var character byte
	var storedNumber int
	var sum int
	var engineNumber bool
	for {
		if character = walkableFileReader.MoveEast(); character != 0 {
			// fmt.Println(string(character))
			if internal.IsRuneANumber(rune(character)) {
				if storedNumber == 0 {
					storedNumber = int(character - '0')
					// fmt.Println("new Number: ", storedNumber, " Checking neighbours")
					if searchNumberNeighbours(internal.CARDINAL_DIRECTIONS) {
						engineNumber = true
					}
				} else {
					storedNumber = (storedNumber * 10) + int(character-'0')
					// fmt.Println("more of the current number: ", storedNumber)
					if engineNumber == false {
						// fmt.Println("checking only numbers to the east")
						if searchNumberNeighbours(internal.QUICK_CARDINAL_DIRECTIONS) {
							engineNumber = true
						}
					} else {
						// fmt.Println("not checking neighbours because it's already touching")
					}
				}
			} else {
				if storedNumber != 0 {
					if engineNumber {
						// fmt.Println("engine part touching: ", storedNumber)
						sum += storedNumber
					} else {
						// fmt.Println("engine part not touching: ", storedNumber)
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

	fmt.Println(processEngine())
}
