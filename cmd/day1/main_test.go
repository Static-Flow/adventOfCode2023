package main

import (
	"adventOfCode2023/internal"
	"testing"
)

func BenchmarkProcessLine(b *testing.B) {
	var lines []string
	if iterator, err := internal.NewLineIterator("benchmarkData"); err == nil {
		for iterator.Next() {
			line := iterator.Line()
			// Process the line
			lines = append(lines, line)
		}
		b.ResetTimer()
		for _, line := range lines {
			b.Run(line, func(b *testing.B) {
				var sum int
				for i := 0; i < b.N; i++ {
					sum += processLine(line)
				}
			})
		}
	}

}
