package main

import (
	"adventOfCode2023/internal"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

var gameNumber int

func processGame(line string) bool {

	if line[6] == ':' {
		gameNumber = int(line[5] - '0')
		line = line[7:]
	} else if line[7] == ':' {
		gameNumber = (int(line[5]-'0') * 10) + int(line[6]-'0')
		line = line[8:]
	} else {
		gameNumber = 100
		line = line[9:]
	}
	var index int
	var lineRune rune
	for index, lineRune = range line {

		switch lineRune {
		case 'r':
			if line[index-1] == ' ' {
				if line[index-2] > '2' && line[index-3] != ' ' {
					return false
				}
			}
		case 'g':
			if line[index-2] > '3' && line[index-3] != ' ' {
				return false
			}
		case 'b':
			if line[index-2] > '4' && line[index-3] != ' ' {
				return false
			}
		}
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
