package main

import (
	"adventOfCode2023/cmd/day2"
	"adventOfCode2023/internal"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

func processPull(pullLine string) bool {
	var max int
	if pullLine[1] == ' ' {
		if max, _ = day2.ColorMap[pullLine[2:]]; int(pullLine[0]-'0') <= max {
			return true
		}
	} else {
		if max, _ = day2.ColorMap[pullLine[3:]]; (int(pullLine[0]-'0')*10)+int(pullLine[1]-'0') <= max {
			return true
		}
	}
	return false
}

var gameNumber int

func processGame(line string) bool {

	if line[6] == ':' {
		gameNumber = int(line[5] - '0')
		line = line[8:]
	} else if line[7] == ':' {
		gameNumber = (int(line[5]-'0') * 10) + int(line[6]-'0')
		line = line[9:]
	} else {
		gameNumber = 100
		line = line[10:]
	}
	var lastMatchIndex int
	var index int
	var lineRune rune
	for index, lineRune = range line {
		switch lineRune {
		case ',':
			fallthrough
		case ';':
			if processPull(line[lastMatchIndex:index]) == false {
				// cube pull was not within allowed range
				return false
			}
			lastMatchIndex = index + 2
		}
	}
	if processPull(line[lastMatchIndex:]) == false {
		// cube pull was not within allowed range
		return false
	}
	return true
}

func main() {
	f, err := os.Create("memprofile.pprof")
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
	defer f.Close() // Make sure to close the file when done

	if iterator, err := internal.NewLineIterator("../input"); err == nil {
		var sum int
		var line string
		for iterator.Next() {
			line = iterator.Line()
			// Process the line
			if processGame(line) {
				sum += gameNumber
			}
		}

		if err = iterator.Err(); err != nil {
			log.Fatalln(err)
		}
		fmt.Println(sum)

	} else {
		log.Fatalln(err)
	}
	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}
}
