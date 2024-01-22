package main

import "testing"

func TestAddBinary(t *testing.T) {
	cases := []struct {
		a    string
		b    string
		want string
	}{
		{"0", "0", "0"},
		{"1", "0", "1"},
		{"0", "1", "1"},
		{"1", "1", "10"},
		{"11", "1", "100"},
		{"1010", "1011", "10101"},
		{"1111", "1111", "11110"},
	}
	for _, c := range cases {
		got := addBinary(c.a, c.b)
		if got != c.want {
			t.Errorf("addBinary(%q, %q) == %q, want %q", c.a, c.b, got, c.want)
		}
	}
}

func TestGetDigit(t *testing.T) {
	cases := []struct {
		numS     string
		index    int
		expected int
	}{
		{"01", 1, 1},
		{"0", 0, 0},
		{"1", 0, 1},
		{"0", 1, 0},
	}

	for _, c := range cases {
		got := getDigit(c.numS, c.index)
		if got != c.expected {
			t.Errorf("getDigit(%s, %d) == %d, wanted %d", c.numS, c.index, got, c.expected)
		}
	}
}
