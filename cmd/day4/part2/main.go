package main

import (
	"adventOfCode2023/cmd/day4"
	"fmt"
	"time"
)

var sc *day4.ScratchCards
var sum int

func processCards() int {
	sum = 0
	for i := 0; i < len(sc.Cards); i++ {
		sum += sc.Cards[i].Count
		if sc.Cards[i].Matches > 0 {
			for j := 1; j <= sc.Cards[i].Matches; j++ {
				sc.Cards[i+j].Count += 1 * sc.Cards[i].Count
			}
		}
	}
	return sum
}

func main() {
	day4.InitCandidateMap()
	sc = day4.NewScratchCards("../input")
	now := time.Now()
	fmt.Println(processCards())
	fmt.Println(time.Since(now))
}
