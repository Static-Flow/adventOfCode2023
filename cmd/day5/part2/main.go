package main

import (
	"fmt"
	"github.com/Static-Flow/adventOfCode2023/internal"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"log"
	"strconv"
	"strings"
	"time"
)

var seeds [][2]int
var numIndex = -1
var num int
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

func binarySearch(low, high int) int {
	lowest := calculateArbitraryValue(high)
	val := calculateArbitraryValue(low)
	isAsc := val < lowest
	if isAsc {
		lowest = val
	}
	for low <= high {
		mid := low + (high-low)/2
		if val = calculateArbitraryValue(mid); val < lowest {
			lowest = val
		}

		if isAsc {
			if val < lowest {
				low = mid + 1
			} else {
				high = mid - 1
			}
		} else {
			if val > lowest {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	fmt.Println("Lowest at: ", low)
	return lowest
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
	// var lowest, lowestIndex, src, dest, max, newValue, index int
	// var mapSection [][3]int
	// var section [3]int
	var low, lowest int

	p := message.NewPrinter(language.English)
	// fmt.Printf("test:%s\n", p.Sprintf("%d", calculateArbitraryValue(682758545)))
	// fmt.Printf("test:%s\n", p.Sprintf("%d", brute(630335678, 701491196)))
	// fmt.Printf("test:%s\n", p.Sprintf("%d", brute(260178142, 385183562)))
	for index := 0; index < len(seeds); index += 2 {
		max := seeds[index][0] + seeds[index+1][0] - 1
		// binaryWalk(seeds[index][0], max)
		mid := seeds[index][0] + ((seeds[index+1][0] - 1) / 2)
		half := seeds[index][0] + ((seeds[index+1][0] - 1) / 4)
		secondHalf := seeds[index][0] + ((seeds[index+1][0] - 1) - ((seeds[index+1][0] - 1) / 4))
		fmt.Println("Start:", seeds[index][0], "Half:", half, "Mid:", mid, "SecondHalf:", secondHalf, "Max:", max)
		fmt.Println("Comparing between: ", seeds[index][0], " and ", max)
		fmt.Println("Start Value: ", seeds[index][0], p.Sprintf("%d", calculateArbitraryValue(seeds[index][0])))
		fmt.Println("Second Value: ", seeds[index][0]+1, p.Sprintf("%d", calculateArbitraryValue(seeds[index][0]+1)))
		fmt.Println("Third Value: ", seeds[index][0]+2, p.Sprintf("%d", calculateArbitraryValue(seeds[index][0]+2)))
		fmt.Println("Fourth Value: ", seeds[index][0]+3, p.Sprintf("%d", calculateArbitraryValue(seeds[index][0]+3)))
		fmt.Println("First Half Value: ", half, p.Sprintf("%d", calculateArbitraryValue(half)))
		fmt.Println("Middle Value: ", mid, p.Sprintf("%d", calculateArbitraryValue(mid)))
		fmt.Println("Second Half Value: ", secondHalf, p.Sprintf("%d", calculateArbitraryValue(secondHalf)))
		fmt.Println("Last Fourth Value: ", max-3, p.Sprintf("%d", calculateArbitraryValue(max-3)))
		fmt.Println("Last Third Value: ", max-2, p.Sprintf("%d", calculateArbitraryValue(max-2)))
		fmt.Println("Last Second Value: ", max-1, p.Sprintf("%d", calculateArbitraryValue(max-1)))
		fmt.Println("Last Value: ", max, p.Sprintf("%d", calculateArbitraryValue(max)))
		fmt.Println()
		fmt.Printf("test:%s\n", p.Sprintf("%d", brute(seeds[index][0], max)))
		if low = binarySearch(seeds[index][0], max); low < lowest || lowest == 0 {
			lowest = low
		}
		fmt.Println("Low", low)
		fmt.Println("Current lowest:", lowest)
		// low = binarySearch(seeds[index][0], seeds[index][0]+seeds[index+1][0]-1)
		// for _, mapSection = range locationMap {
		// 	for _, section = range mapSection {
		// 		dest = section[0]
		// 		src = section[1]
		// 		max = section[2] + dest
		// 		if newValue = (dest - src) + seeds[index][1]; newValue >= dest && newValue < max {
		// 			seeds[index][1] = newValue
		// 			break
		// 		}
		// 	}
		// }
		// if low < lowest || lowest == 0 {
		//	lowest = low
		// }
	}
	// fmt.Println(seeds)
	fmt.Println(lowest)
	// fmt.Println(lowest, "lowest Index:", lowestIndex, "lowest start seed:", seeds[lowestIndex][0])
	// fmt.Println(binarySearch(seeds[lowestIndex][0], seeds[lowestIndex][0]+seeds[lowestIndex+1][0]))
	fmt.Println(time.Since(now))

}
