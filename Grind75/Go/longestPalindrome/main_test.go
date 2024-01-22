package main

import "testing"

func TestLongestPalindrome(t *testing.T) {
	cases := []struct {
		s    string
		want int
	}{
		{"abccccdd", 7},
		{"a", 1},
		{"bb", 2},
	}
	for _, c := range cases {
		got := longestPalindrome(c.s)
		if got != c.want {
			t.Errorf("longestPalindrome(%q) == %d, want %d", c.s, got, c.want)
		}
	}
}
