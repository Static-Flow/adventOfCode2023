package main

import (
	"adventOfCode2023/internal"
	"strings"
	"testing"
)

var globalRes int

func BenchmarkProcessGame(b *testing.B) {
	var lines []string
	var innerRes int
	if iterator, err := internal.NewLineIterator("../benchmarkData"); err == nil {
		for iterator.Next() {
			line := iterator.Line()
			// Process the line
			lines = append(lines, line)
		}
		b.ResetTimer()
		for _, line := range lines {
			b.Run(line, func(b *testing.B) {
				pieces := strings.SplitN(line, ":", 2)
				for i := 0; i < b.N; i++ {
					innerRes = processGame(pieces[1])
				}
				globalRes = innerRes
			})
		}
	}

}
