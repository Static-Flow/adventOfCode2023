package main

import (
	"fmt"
	"github.com/Static-Flow/adventOfCode2023/cmd/day4"
	"math"
	"time"
)

var sc *day4.ScratchCards
var sum float64

func processCards() float64 {
	sum = 0
	for i := 0; i < len(sc.Cards); i++ {
		if sc.Cards[i].Matches > 0 {
			sum += math.Pow(2, float64(sc.Cards[i].Matches-1))
		}
	}
	return sum
}

func main() {
	sc = day4.NewScratchCards("../input")
	now := time.Now()
	processCards()
	fmt.Println(sum)
	fmt.Println(time.Since(now))
}
