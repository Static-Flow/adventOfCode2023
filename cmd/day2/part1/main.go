package main

import (
	"adventOfCode2023/internal"
	"fmt"
	"log"
	"strings"
)

func processGame(line string) bool {
	var lineRune rune
	var lastNumber int
	var index int
	for index, lineRune = range line {

		if internal.IsRuneANumber(lineRune) {
			if lastNumber == 0 {
				lastNumber = int(lineRune - '0')
			} else {
				lastNumber = (lastNumber * 10) + int(lineRune-'0')
				if lastNumber > 12 {
					switch line[index+2] {
					case 'r':
						return false
					case 'g':
						if lastNumber > 13 {
							return false
						}
					case 'b':
						if lastNumber > 14 {
							return false
						}
					}
				}
			}
		} else {
			lastNumber = 0
		}
	}
	return true
}

func main() {

	if iterator, err := internal.NewLineIterator("../input"); err == nil {
		var sum int
		var line string
		for iterator.Next() {
			line = iterator.Line()
			pieces := strings.SplitN(line, ":", 2)
			// Process the line
			if processGame(pieces[1]) {
				switch len(pieces[0]) {
				case 6:
					sum += int(line[5] - '0')
				case 7:
					sum += (int(line[5]-'0') * 10) + int(line[6]-'0')
				default:
					sum += 100
				}
			}
		}

		if err = iterator.Err(); err != nil {
			log.Fatalln(err)
		}
		fmt.Println(sum)

	} else {
		log.Fatalln(err)
	}
}
