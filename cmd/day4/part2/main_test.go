package main

import (
	"github.com/Static-Flow/adventOfCode2023/cmd/day4"
	"testing"
)

var globalRes int

func BenchmarkProcessCards(b *testing.B) {
	sc = day4.NewScratchCards("../input")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum = processCards()

	}
	globalRes = sum

}
