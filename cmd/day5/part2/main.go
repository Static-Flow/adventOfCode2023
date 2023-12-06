package main

import (
	"fmt"
	"github.com/Static-Flow/adventOfCode2023/internal"
	"log"
	"strconv"
	"strings"
	"time"
)

var seeds [][2]int
var numIndex = -1
var seed, seedRange int
var sectionMap [][3]int
var locationMap [][][3]int

func processSeeds(line string) {
	seedLine := strings.Split(line, " ")
	for i := 0; i < len(seedLine); i += 2 {
		seed, _ = strconv.Atoi(seedLine[i])
		seedRange, _ = strconv.Atoi(seedLine[i+1])
		seeds = append(seeds, [2]int{seed, seed})
		seeds = append(seeds, [2]int{seedRange, seedRange})
		// seeds = append(seeds, [2]int{seed + seedRange, seed + seedRange})
	}
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

func brute(start, finish int) (lowest int) {
	fmt.Println("Finding lowest the hard way between: ", start, finish)
	lowestIndex := 0
	for {
		if start != finish {
			if value := calculateArbitraryValue(start); lowest == 0 || value < lowest {
				lowest = value
				lowestIndex = start
				fmt.Println("NEW LOW:", value, " at ", start, "before: ", calculateArbitraryValue(start-1), " after: ", calculateArbitraryValue(start+1))
			}
			start++
		} else {
			break
		}
	}
	fmt.Println("Found lowest", lowest, " at ", lowestIndex)
	return
}

func main() {
	if iterator, err := internal.NewLineIterator("../input"); err == nil {
		iterator.Next()
		processSeeds(iterator.Line()[7:])
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
	locationMap = append(locationMap, sectionMap)
	now := time.Now()
	var low, lowest int

	for index := 0; index < len(seeds); index += 2 {
		max := seeds[index][0] + seeds[index+1][0] - 1
		low = brute(seeds[index][0], max)
		if low < lowest {
			lowest = low
		}
	}
	fmt.Println(lowest)
	fmt.Println(time.Since(now))

}
