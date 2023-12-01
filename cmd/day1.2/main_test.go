package main

import (
	"adventOfCode2023/internal"
	"testing"
)

func BenchmarkProcessLine(b *testing.B) {
	var lines []string
	if iterator, err := internal.NewLineIterator("input"); err == nil {
		for iterator.Next() {
			line := iterator.Line()
			// Process the line
			lines = append(lines, line)
		}
		b.ResetTimer()
		var sum int
		for i := 0; i < b.N; i++ {
			sum = 0
			for _, line := range lines {
				sum += processLine(line)
			}
		}
	}

}
