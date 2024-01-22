package main

import "testing"

func TestMajorityElement(t *testing.T) {
	tests := map[string]struct {
		nums []int
		want int
	}{
		"case 2": {
			nums: []int{2, 2, 1, 1, 1, 2, 2},
			want: 2,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := majorityElement(tt.nums); got != tt.want {
				t.Fatalf("got %d, wanted %d", got, tt.want)
			}
		})
	}
}
