package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	gave int
	want [][]int
}{
	{3, [][]int{{1, 2, 3}, {2, 4, 6}, {3, 6, 9}}},
	{4, [][]int{{1, 2, 3, 4}, {2, 4, 6, 8}, {3, 6, 9, 12}, {4, 8, 12, 16}}},
}

func isEqual(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, e := range a {
		if e != b[i] {
			return false
		}
	}
	return true
}

func singleTest(i int, wanted [][]int) string {
	res := MultiplicationTable(i)
	for i, e := range res {
		equal := isEqual(wanted[i], e)
		if !equal {
			return fmt.Sprintf("Got %d, wanted %d", res, wanted)
		}
	}
	return ""
}

func TestMain(t *testing.T) {
	for _, test := range testCases {
		testName := fmt.Sprintf("Case '%d'", test.gave)
		t.Run(testName, func(t *testing.T) {
			if res := singleTest(test.gave, test.want); res != "" {
				t.Fatalf(res)
			}
		})
	}
}

func BenchmarkTheirs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Comparison(1)
	}
}
func BenchmarkMine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MultiplicationTable(1)
	}
}
