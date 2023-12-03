package main

import (
	"adventOfCode2023/cmd/day2"
	"adventOfCode2023/internal"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

func processPull(pullLine string) *day2.Cube {
	var cube *day2.Cube
	var count int
	var color string
	var max int
	if pullLine[1] == ' ' {
		count = int(pullLine[0] - '0')
		color = pullLine[2:]
		if max, _ = day2.ColorMap[color]; count <= max {
			cube = &day2.Cube{
				CubeColor: color,
				Count:     count,
			}
		}
	} else {
		count = (int(pullLine[0]-'0') * 10) + int(pullLine[1]-'0')
		color = pullLine[3:]
		if max, _ = day2.ColorMap[color]; count <= max {
			cube = &day2.Cube{
				CubeColor: color,
				Count:     count,
			}
		}
	}
	return cube
}

var game = &day2.Game{Rounds: []*day2.Round{}}
var round = &day2.Round{Cubes: []*day2.Cube{}}

func processGame(line string) *day2.Game {
	game.Rounds = game.Rounds[:0]
	if line[6] == ':' {
		game.Id = int(line[5] - '0')
		line = line[8:]
	} else if line[7] == ':' {
		game.Id = (int(line[5]-'0') * 10) + int(line[6]-'0')
		line = line[9:]
	} else {
		game.Id = 100
		line = line[10:]
	}
	var lastMatchIndex int
	round.Cubes = round.Cubes[:0]
	var cubePull *day2.Cube
	var index int
	var lineRune rune
	for index, lineRune = range line {
		switch lineRune {
		case ',':
			if cubePull = processPull(line[lastMatchIndex:index]); cubePull != nil {
				// cube pull was within allowed range
				round.Cubes = append(round.Cubes, cubePull)
				lastMatchIndex = index + 2
			} else {
				// cube pull was outside allowed range so return a nil game
				game.Rounds = nil
				return game
			}
		case ';':
			if cubePull = processPull(line[lastMatchIndex:index]); cubePull != nil {
				round.Cubes = append(round.Cubes, cubePull)
				game.Rounds = append(game.Rounds, round)
				round.Cubes = round.Cubes[:0]
				lastMatchIndex = index + 2
			} else {
				game.Rounds = nil
				return game
			}
		}
	}
	if cubePull = processPull(line[lastMatchIndex:]); cubePull != nil {
		round.Cubes = append(round.Cubes, processPull(line[lastMatchIndex:]))
		game.Rounds = append(game.Rounds, round)
		return game
	} else {
		game.Rounds = nil
		return game
	}
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
			processGame(line)
			if game.Rounds != nil {
				sum += game.Id
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
