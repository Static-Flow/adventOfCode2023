package main

import (
	"strconv"
	"testing"
)

var globalRes = 0

func BenchmarkMyConv(b *testing.B) {
	test := "2977255263"
	for i := 0; i < b.N; i++ {
		for _, character := range test {
			if globalRes == 0 {
				globalRes = int(character - '0')
			} else {
				globalRes = (globalRes * 10) + int(character-'0')
			}
		}
		globalRes = 0
	}

}

func BenchmarkStdLibConv(b *testing.B) {
	test := "2977255263"
	for i := 0; i < b.N; i++ {
		globalRes, _ = strconv.Atoi(test)
	}
}
