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
	var ratios [2]int
	var lastEngineLocation int
	var exists bool
	engineGears := map[int][2]int{}
	for {
		if character = day3.WalkableReader.ReadOrMoveDirection(true, day3.EAST); character != 0 {
			if internal.IsRuneANumber(rune(character)) {
				if storedNumber == 0 {
					storedNumber = int(character - '0')
					if day3.SearchNeighboursForGears(day3.CARDINAL_DIRECTIONS) {
						engineNumber = true
						lastEngineLocation = day3.WalkableReader.LastSearchLocation()
					}
				} else {
					storedNumber = (storedNumber * 10) + int(character-'0')
					if engineNumber == false {
						if day3.SearchNeighboursForGears(day3.QUICK_CARDINAL_DIRECTIONS) {
							// number touches a gear
							engineNumber = true
							lastEngineLocation = day3.WalkableReader.LastSearchLocation()
						}
					}
				}
			} else {
				if storedNumber != 0 {
					if engineNumber {
						if ratios, exists = engineGears[lastEngineLocation]; !exists {
							engineGears[lastEngineLocation] = [2]int{storedNumber}
						} else if ratios[1] != 0 {
							delete(engineGears, lastEngineLocation)
						} else {
							ratios[1] = storedNumber
							engineGears[lastEngineLocation] = ratios
						}
					}
					engineNumber = false
					storedNumber = 0
				}
			}
		} else {
			break
		}
	}
	for _, ratios := range engineGears {
		if ratios[1] != 0 {
			sum += ratios[0] * ratios[1]
		}
	}
	return sum
}

func main() {
	fmt.Println(processEngine())
}
