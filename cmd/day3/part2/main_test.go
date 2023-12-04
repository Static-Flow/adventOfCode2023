package main

import (
	"github.com/Static-Flow/adventOfCode2023/cmd/day3"
	"testing"
)

var globalRes int

func BenchmarkProcessEngine(b *testing.B) {
	day3.WalkableReader = day3.NewWalkableByteStream("../input")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		day3.WalkableReader.Reset()
		res = processEngine()
	}
	globalRes = res

}
