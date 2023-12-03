package main

import (
	"testing"
)

var globalRes int

func BenchmarkMain(b *testing.B) {
	var res int
	for i := 0; i < b.N; i++ {
		walkableReader.Reset()
		res = processEngine()
	}
	globalRes = res

}
