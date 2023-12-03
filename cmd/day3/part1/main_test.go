package main

import (
	"testing"
)

var globalRes int

func BenchmarkMain(b *testing.B) {
	var res int
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		walkableFileReader.ResetFile()
		b.StartTimer()
		res = processEngine()

	}
	globalRes = res

}
