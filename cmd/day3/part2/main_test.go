package main

import (
	"adventOfCode2023/cmd/day3"
	"testing"
)

var globalRes int

func BenchmarkMain(b *testing.B) {
	var res int
	for i := 0; i < b.N; i++ {
		day3.WalkableReader.Reset()
		res = processEngine()
	}
	globalRes = res

}
