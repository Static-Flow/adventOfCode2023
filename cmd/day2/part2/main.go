package main

import (
	"adventOfCode2023/cmd/day2"
	"adventOfCode2023/internal"
	"fmt"
	"log"
)

var game = &day2.GamePart2{MaxColorsMap: map[string]int{day2.GREEN: 0, day2.BLUE: 0, day2.RED: 0}}

func processGame(line string) *day2.GamePart2 {
	game.MaxColorsMap[day2.RED] = 0
	game.MaxColorsMap[day2.GREEN] = 0
	game.MaxColorsMap[day2.BLUE] = 0

	skip := 5
	if line[skip+1] == ':' {
		skip += 3
	} else if line[skip+2] == ':' {
		skip += 4
	} else {
		skip += 5
	}
	var lastMatchIndex int
	var index int
	var lineRune rune
	for index, lineRune = range line[skip:] {
		switch lineRune {
		case ',':
			game.ProcessPull(line[skip+lastMatchIndex : skip+index])
			lastMatchIndex = index + 2
		case ';':
			game.ProcessPull(line[skip+lastMatchIndex : skip+index])
			lastMatchIndex = index + 2
		}
	}
	game.ProcessPull(line[skip+lastMatchIndex:])
	return game
}

func main() {

	if iterator, err := internal.NewLineIterator("../input"); err == nil {
		var sum int
		var gameSum int
		var value int
		for iterator.Next() {
			line := iterator.Line()
			// Process the line
			// processGame(line)
			// sum += 1 * game.MaxColorsMap[day2.BLUE] * game.MaxColorsMap[day2.GREEN] * game.MaxColorsMap[day2.RED]
			gameSum = 1
			for _, value = range processGame(line).MaxColorsMap {
				gameSum *= value
			}
			sum += gameSum
		}

		if err = iterator.Err(); err != nil {
			log.Fatalln(err)
		}
		fmt.Println(sum)

	} else {
		log.Fatalln(err)
	}
}
