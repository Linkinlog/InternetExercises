package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input          int
	ExpectedOutput int
}{
	{
		Input:          0,
		ExpectedOutput: 0,
	},
	{
		Input:          -15,
		ExpectedOutput: 0,
	},
	{
		Input:          100,
		ExpectedOutput: 2318,
	},
	{
		Input:          10,
		ExpectedOutput: 23,
	},
}

func TestMultiple3And5(t *testing.T) {
	for i, e := range testCases {
		name := fmt.Sprintf("Test %d: ", i)
		t.Run(name, func(t *testing.T) {
			if res := Multiple3And5(e.Input); res != e.ExpectedOutput {
				t.Fatalf("Wanted %d, got %d", e.ExpectedOutput, res)
			}
		})
	}
}

func BenchmarkMultiple3And53(b *testing.B) {
	for i := 1; i < b.N; i++ {
		Multiple3And53(100)
		//Multiple3And5(1000000000000000000)
	}
}
func BenchmarkMultiple3And52(b *testing.B) {
	for i := 1; i < b.N; i++ {
		Multiple3And52(100)
		//Multiple3And5(1000000000000000000)
	}
}
func BenchmarkMultiple3And5(b *testing.B) {
	for i := 1; i < b.N; i++ {
		Multiple3And5(100)
		//Multiple3And5(1000000000000000000)
	}
}

