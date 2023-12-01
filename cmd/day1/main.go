package main

import (
	"adventOfCode2023/internal"
	"log"
	"regexp"
)

var characterRegex = regexp.MustCompile("\\D")

func processLine(line string) int {
	filteredStr := characterRegex.ReplaceAllString(line, "")
	if len(filteredStr) == 1 {
		return int(filteredStr[0]-'0') * 11
	} else {
		return (int(filteredStr[0]-'0') * 10) + int(filteredStr[len(filteredStr)-1]-'0')
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
