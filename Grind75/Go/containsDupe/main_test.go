package main

import "testing"

func TestContainsDuplicate(t *testing.T) {
	cases := []struct {
		nums     []int
		expected bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}

	for _, c := range cases {
		result := containsDuplicate(c.nums)
		if result != c.expected {
			t.Errorf("containsDuplicate(%v) == %t, expected %t", c.nums, result, c.expected)
		}
	}
}
