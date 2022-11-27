package main

import (
	"fmt"
	"testing"
)

// testCases is a struct defining all of our testcases and their types
var testCases = []struct {
	arg      [][2]int
	expected int
}{
	{[][2]int{{1, 5}}, 4},
	{[][2]int{{1, 5}, {6, 10}}, 8},
	{[][2]int{{1, 5}, {1, 5}}, 4},
	{[][2]int{{1, 4}, {7, 10}, {3, 5}}, 7},
}

// TestSumOfInterVals tests SumOfIntervals using
func TestSumOfInterVals(t *testing.T) {
	// Logic
	for _, e := range testCases {
		testName := fmt.Sprintf(`Case %d`, e.arg)
		t.Run(testName, func(t *testing.T) {
			if res := SumOfIntervals(e.arg); res != e.expected {
				t.Fatalf(`Got: %d, Wanted: %d`, res, e.expected)
			}
		})
	}
}
