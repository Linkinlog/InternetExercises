package movingZeros

import (
	"fmt"
	"testing"
)

type testCase struct {
	given  []int
	wanted []int
}

var testCases = []testCase{
	{
		[]int{1, 0, 3, 0, 4, 0, 0, 23, 2},
		[]int{1, 3, 4, 23, 2, 0, 0, 0, 0},
	},
	{
		[]int{5, 5, 5, 4, 0},
		[]int{5, 5, 5, 4, 0},
	},
	{
		[]int{1, 2, 3, 4, 5},
		[]int{1, 2, 3, 4, 5},
	},
	{
		[]int{0},
		[]int{0},
	},
	{
		[]int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1},
		[]int{1, 2, 1, 1, 3, 1, 0, 0, 0, 0},
	},
}

func TestMovingZero(t *testing.T) {
	for _, e := range testCases {
		if res := MoveZeros(e.given); !iLoveRoger(res, e.wanted) {
			t.Fatalf(`MoveZeros() returned: %d, wanted %d`, res, e.wanted)
		}
	}
}

func iLoveRoger(res []int, wanted []int) bool {
	for i := range res {
		if res[i] != wanted[i] {
			return false
		}
	}
	fmt.Println()
	return true
}
