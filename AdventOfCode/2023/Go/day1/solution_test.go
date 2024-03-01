package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"multiple lines": {
			input: `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`,
			want: 142,
		},
		"one line": {
			input: "1abc2",
			want:  12,
		},
		"only one number": {
			input: "1",
			want:  11,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := part1(strings.Split(tt.input, "\n")); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"one line": {
			input: "xtwone3four",
			want:  24,
		},
		"only one number": {
			input: "one",
			want:  11,
		},
		"two lines": {
			input: `two1nine
eightwothree`,
			want: 29 + 83,
		},
		"three lines": {
			input: `two1nine
eightwothree
abcone2threexyz`,
			want: 29 + 83 + 13,
		},
		"four lines": {
			input: `two1nine
eightwothree
abcone2threexyz
xtwone3four`,
			want: 29 + 83 + 13 + 24,
		},
		"five lines": {
			input: `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2`,
			want: 29 + 83 + 13 + 24 + 42,
		},
		"six lines": {
			input: `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234`,
			want: 29 + 83 + 13 + 24 + 42 + 14,
		},
		"multiple lines": {
			input: `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`,
			want: 281,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := part2(strings.Split(tt.input, "\n")); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
