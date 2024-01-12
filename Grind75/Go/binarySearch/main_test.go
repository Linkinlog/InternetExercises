package main

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := map[string]struct {
		nums           []int
		target         int
		expectedResult int
	}{
		"first": {
			nums: []int{-1,0,3,5,9,12},
			target: 9,
			expectedResult: 4,
		},
		"second": {
			nums: []int{-1,0,3,5,9,12},
			target: 2,
			expectedResult: -1,
		},
		"third": {
			nums: []int{2,5},
			target: 5,
			expectedResult: 1,
		},
		"fourth": {
			nums: []int{5},
			target: 5,
			expectedResult: 0,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if val := search(tt.nums, tt.target); val != tt.expectedResult {
				t.Fatalf("got %d, wanted %d", val, tt.expectedResult)
			}
		})
	}
}
