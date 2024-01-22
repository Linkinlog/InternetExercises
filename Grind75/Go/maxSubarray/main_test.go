package main

import "testing"

func TestMaxSubarray(t *testing.T) {
	t.Parallel()
	cases := []struct {
		nums     []int
		expected int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1}, 1},
		{[]int{5, 4, -1, 7, 8}, 23},
		{[]int{-2, 1}, 1},
		{[]int{-2, -1}, -1},
		{[]int{-1, -2}, -1},
		{[]int{-1, 0, -2}, 0},
		{[]int{-1, 0, 1, -2}, 1},
		{[]int{-1, 0, 1, -2, 2}, 2},
		{[]int{-1, 0, 1, -2, 2, 3}, 5},
		{[]int{-1, 0, 1, -2, 2, 3, -1}, 6},
		{[]int{-1, 0, 1, -2, 2, 3, -1, 2}, 7},
		{[]int{-1, 0, 1, -2, 2, 3, -1, 2, 3}, 8},
		{[]int{-1, 0, 1, -2, 2, 3, -1, 2, 3, -1}, 8},
		{[]int{-1, 0, 1, -2, 2, 3, -1, 2, 3, -1, 2}, 9},
		{[]int{-1, 0, 1, -2, 2, 3, -1, 2, 3, -1, 2, 3}, 10},
		{[]int{-1, 0, 1, -2, 2, 3, -1, 2, 3, -1, 2, 3, -1}, 10},
	}

	for _, c := range cases {
		got := maxSubArray(c.nums)
		if got != c.expected {
			t.Errorf("maxSubArray(%v) == %d, want %d", c.nums, got, c.expected)
		}
	}
}
