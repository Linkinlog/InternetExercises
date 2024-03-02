package main

import (
	"strings"
	"testing"
)

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1(strings.Split(input, "\n"), subset{
			red:   69,
			blue:  69,
			green: 420,
		})
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part2(strings.Split(input, "\n"), subset{
			red:   69,
			blue:  69,
			green: 420,
		})
	}
}
