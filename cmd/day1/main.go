package main

import (
	"adventOfCode2023/internal"
	"log"
)

func processLine(line string) int {
	var firstInt int
	var secondInt int
	for _, lineRune := range line {
		if internal.IsRuneANumber(lineRune) {
			if firstInt == 0 {
				firstInt = int(lineRune - '0')
			} else {
				secondInt = int(lineRune - '0')
			}
		}
	}
	if secondInt == 0 {
		return firstInt * 11
	} else {
		return (firstInt * 10) + secondInt
	}
}

func main() {
	sum := 0
	if iterator, err := internal.NewLineIterator("input"); err == nil {
		for iterator.Next() {
			line := iterator.Line()
			// Process the line
			sum += processLine(line)
		}

		if err = iterator.Err(); err != nil {
			log.Fatalln(err)
		}
	} else {
		log.Fatalln(err)
	}
	log.Println(sum)
}
