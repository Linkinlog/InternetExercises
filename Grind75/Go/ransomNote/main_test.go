package main

import "testing"

func TestCanConstruct(t *testing.T) {
	var cases = []struct {
		ransomNote string
		magazine   string
		expected   bool
	}{
		{"a", "b", false},
		{"aa", "ab", false},
		{"aa", "aab", true},
	}
	for _, c := range cases {
		actual := canConstruct(c.ransomNote, c.magazine)
		if actual != c.expected {
			t.Errorf("Input: %s, %s. Expected: %t, actual: %t", c.ransomNote, c.magazine, c.expected, actual)
		}
	}
}
