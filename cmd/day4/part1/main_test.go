package main

import (
	"adventOfCode2023/cmd/day3"
	"testing"
)

var globalRes int

func BenchmarkProcessCards(b *testing.B) {
	day3.WalkableReader = day3.NewWalkableByteStream("../input")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		day3.WalkableReader.Reset()
		day3.WalkableReader.SetLocation(7)
		sum = 0
		b.StartTimer()
		processCards()
	}
	globalRes = sum

}
