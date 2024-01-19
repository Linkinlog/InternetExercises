package main

import "testing"

func TestClimbStairs(t *testing.T) {
	cases := map[string]struct {
		input int
		want  int
	}{
		"Case 1": {input: 2, want: 2},
		"Case 2": {input: 3, want: 3},
		"Case 3": {input: 4, want: 5},
	} 

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := climbStairs(tc.input)
			if got != tc.want {
				t.Fatalf("climbStairs(%d) = %d; want %d", tc.input, got, tc.want)
			}
		})
	}
}
