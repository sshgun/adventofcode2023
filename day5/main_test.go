package main

import "testing"

var result int

func BenchmarkPart2(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = part2()
	}
	result = r
}
