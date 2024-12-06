package main

import (
	"strings"
	"testing"
)

func BenchmarkPart1(b *testing.B) {
	inp := strings.NewReader(input)
	for i := 0; i < b.N; i++ {
		part1(inp)
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part2(strings.Split(input, "\n"))
	}
}
