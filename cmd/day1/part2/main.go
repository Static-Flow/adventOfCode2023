package main

import (
	"adventOfCode2023/internal"
	"log"
)

var digits = map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}

func process(line string) int {
	var firstInt int
	var secondInt int
	lineLength := len(line)
	for index, lineRune := range line {
		if internal.IsRuneANumber(lineRune) {
			// rune is a number
			if firstInt == 0 {
				firstInt = int(lineRune - '0')
			} else {
				secondInt = int(lineRune - '0')
			}
		} else {
			switch lineRune {
			case 'e':
				if index+5 > lineLength {
					continue
				}
				if value, exists := digits[line[index:index+5]]; exists {
					if firstInt == 0 {
						firstInt = value
					} else {
						secondInt = value
					}
				}
			case 'f':
				fallthrough
			case 'n':
				if index+4 > lineLength {
					continue
				}
				if value, exists := digits[line[index:index+4]]; exists {
					if firstInt == 0 {
						firstInt = value
					} else {
						secondInt = value
					}
				}
			case 'o':
				if index+3 > lineLength {
					continue
				}
				if value, exists := digits[line[index:index+3]]; exists {
					if firstInt == 0 {
						firstInt = value
					} else {
						secondInt = value
					}
				}
			case 's':
				if index+3 > lineLength {
					continue
				}
				if value, exists := digits[line[index:index+3]]; exists {
					if firstInt == 0 {
						firstInt = value
					} else {
						secondInt = value
					}
				}
				if index+5 > lineLength {
					continue
				}
				if value, exists := digits[line[index:index+5]]; exists {
					if firstInt == 0 {
						firstInt = value
					} else {
						secondInt = value
					}
				}
			case 't':
				if index+3 > lineLength {
					continue
				}
				if value, exists := digits[line[index:index+3]]; exists {
					if firstInt == 0 {
						firstInt = value
					} else {
						secondInt = value
					}
				}
				if index+5 > lineLength {
					continue
				}
				if value, exists := digits[line[index:index+5]]; exists {
					if firstInt == 0 {
						firstInt = value
					} else {
						secondInt = value
					}
				}
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
	if iterator, err := internal.NewLineIterator("../input"); err == nil {
		for iterator.Next() {
			line := iterator.Line()
			// Process the line
			sum += process(line)
		}

		if err = iterator.Err(); err != nil {
			log.Fatalln(err)
		}
	} else {
		log.Fatalln(err)
	}
	log.Println(sum)
}
