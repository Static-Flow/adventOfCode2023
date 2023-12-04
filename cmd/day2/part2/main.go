package main

import (
	"adventOfCode2023/internal"
	"log"
	"strings"
)

func processGame(line string) int {
	var maxGreen, maxRed, maxBlue = 1, 1, 1
	var lineRune rune
	var lastNumber int
	for _, lineRune = range line {

		if internal.IsRuneANumber(lineRune) {
			if lastNumber == 0 {
				lastNumber = int(lineRune - '0')
			} else {
				lastNumber = (lastNumber * 10) + int(lineRune-'0')
			}
		} else {
			if lastNumber != 0 {
				switch lineRune {
				case 'r':
					if lastNumber > maxRed {
						maxRed = lastNumber
					}
					lastNumber = 0
				case 'g':
					if lastNumber > maxGreen {
						maxGreen = lastNumber
					}
					lastNumber = 0
				case 'b':
					if lastNumber > maxBlue {
						maxBlue = lastNumber
					}
					lastNumber = 0
				}
			}
		}
	}
	// //fmt.Println("Result:", maxRed*maxBlue*maxGreen)
	return maxRed * maxBlue * maxGreen
}

func main() {

	if iterator, err := internal.NewLineIterator("../input"); err == nil {
		var sum int
		for iterator.Next() {
			line := iterator.Line()
			pieces := strings.SplitN(line, ":", 2)
			// Process the line
			sum += processGame(pieces[1])
		}

		if err = iterator.Err(); err != nil {
			log.Fatalln(err)
		}
		// fmt.Println(sum)

	} else {
		log.Fatalln(err)
	}
}
