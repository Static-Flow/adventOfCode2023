package main

import (
	"adventOfCode2023/internal"
	"log"
)

var digits = map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}

var digitsMap = internal.NewTree()

func processLine(line string) int {
	var ints []int
	var currentTree *internal.Tree
	var exists bool
	for index, lineRune := range line {
		if internal.IsRuneANumber(lineRune) {
			// rune is a number
			ints = append(ints, int(lineRune-'0'))
		} else {
			// rune is letter
			if currentTree, exists = digitsMap[lineRune]; exists {
				// rune matches the start of a "digit"
				for idx, lRune := range line[index+1:] {
					if currentTree, exists = currentTree.Next[lRune]; exists {
						// next character matches, continue
						if currentTree == nil {
							// end of a match so a digit has been found
							ints = append(ints, digits[line[index:index+idx+2]])
							break
						}
					} else {
						// next character didn't match, so we skip ahead in the main index
						index += idx + 1
						break
					}
				}
			}
		}
	}
	return (ints[0] * 10) + ints[len(ints)-1]
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
