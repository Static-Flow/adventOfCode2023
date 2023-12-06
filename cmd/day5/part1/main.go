package main

import (
	"fmt"
	"github.com/Static-Flow/adventOfCode2023/internal"
	"log"
	"strconv"
	"time"
)

var seeds []int
var numIndex = -1
var num int
var locationMap [][][3]int

func processSeeds(line string) {
	for index, character := range line {
		if internal.IsRuneANumber(character) {
			if numIndex == -1 {
				numIndex = index
			}
		} else {
			if numIndex != -1 {
				num, _ = strconv.Atoi(line[numIndex:index])
				seeds = append(seeds, num)
				numIndex = -1
			}
		}
	}
	num, _ = strconv.Atoi(line[numIndex:])
	seeds = append(seeds, num)
	numIndex = -1
}

func processLine(line string) [3]int {
	var destRange, srcRange, offset = -1, -1, -1
	for index, character := range line {
		if internal.IsRuneANumber(character) {
			if numIndex == -1 {
				numIndex = index
			}
		} else {
			if numIndex != -1 {
				if destRange == -1 {
					destRange, _ = strconv.Atoi(line[numIndex:index])
				} else if srcRange == -1 {
					srcRange, _ = strconv.Atoi(line[numIndex:index])
				}
				numIndex = -1
			}
		}
	}
	offset, _ = strconv.Atoi(line[numIndex:])
	numIndex = -1
	return [3]int{destRange, srcRange, offset}
}

var sectionMap [][3]int

func calculateArbitraryValue(index int) int {
	var res = [2]int{index, index}
	var value, dest, src, max int
	for _, mapSection := range locationMap {
		for _, section := range mapSection {
			dest = section[0]
			src = section[1]
			max = section[2] + dest
			if value = (dest - src) + res[1]; value >= dest && value < max {
				res[1] = value
				break
			}
		}
	}
	return res[1]
}

func main() {
	if iterator, err := internal.NewLineIterator("../input"); err == nil {
		iterator.Next()
		processSeeds(iterator.Line())
		iterator.Next() // skip blank line
		iterator.Next() // skip first section header
		for iterator.Next() {
			line := iterator.Line()
			if len(line) > 0 {
				if internal.IsRuneANumber(rune(line[0])) {
					section := processLine(line)
					sectionMap = append(sectionMap, section)
				}
			} else {
				iterator.Next()
				locationMap = append(locationMap, sectionMap)
				sectionMap = nil
			}
			// Process the line
		}

		if err = iterator.Err(); err != nil {
			log.Fatalln(err)
		}
	} else {
		log.Fatalln(err)
	}
	var storeSeeds = make([]int, len(seeds))
	copy(storeSeeds, seeds)
	locationMap = append(locationMap, sectionMap)
	now := time.Now()
	var lowestIndex, lowest, src, dest, max, newValue, index int
	var mapSection [][3]int
	var section [3]int
	for index = range seeds {
		for _, mapSection = range locationMap {
			for _, section = range mapSection {
				dest = section[0]
				src = section[1]
				max = section[2] + dest
				if newValue = (dest - src) + seeds[index]; newValue >= dest && newValue < max {
					fmt.Println(newValue)
					seeds[index] = newValue
					break
				}
			}
		}
		if seeds[index] < lowest || lowest == 0 {
			lowest = seeds[index]
			lowestIndex = index
		}
	}
	fmt.Println(lowest, lowestIndex, storeSeeds[lowestIndex], calculateArbitraryValue(storeSeeds[lowestIndex]))
	fmt.Println(time.Since(now))
}
