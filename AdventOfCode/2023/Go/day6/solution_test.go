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
		"whole shebang": {
			input: `Time:      7  15   30
Distance:  9  40  200`,
			want: 288,
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
		"whole shebang": {
			input: `Time:      7  15   30
Distance:  9  40  200`,
			want: 71503,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := part2(strings.Split(tt.input, "\n")); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParse(t *testing.T) {
	tests := map[string]struct {
		input string
		want  []race
	}{
		"one race": {
			input: "Time: 7\nDistance: 9",
			want:  []race{{time: 7, record: 9}},
		},
		"3 races": {
			input: "Time:      7  15   30\nDistance:  9  40  200",
			want: []race{
				{time: 7, record: 9},
				{time: 15, record: 40},
				{time: 30, record: 200},
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := parse(strings.Split(tt.input, "\n"))
			if len(got) != len(tt.want) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}

			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("parse() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestParseMkII(t *testing.T) {
	tests := map[string]struct {
		input string
		want  race
	}{
		"one race": {
			input: "Time:      7  15   30\nDistance:  9  40  200",
			want: race{
				time: 71530, record: 940200,
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := parseMkII(strings.Split(tt.input, "\n"))
			if got != tt.want {
				t.Errorf("parseMkII() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWouldTravel(t *testing.T) {
	tests := map[string]struct {
		holdTime int
		maxTime  int
		want     int
	}{
		"dont hold, max 7": {
			holdTime: 0,
			maxTime:  7,
			want:     0,
		},
		"hold 1, max 7": {
			holdTime: 1,
			maxTime:  7,
			want:     6,
		},
		"hold 2, max 7": {
			holdTime: 2,
			maxTime:  7,
			want:     10,
		},
		"hold 3, max 7": {
			holdTime: 3,
			maxTime:  7,
			want:     12,
		},
		"hold 4, max 7": {
			holdTime: 4,
			maxTime:  7,
			want:     12,
		},
		"hold 5, max 7": {
			holdTime: 5,
			maxTime:  7,
			want:     10,
		},
		"hold 6, max 7": {
			holdTime: 6,
			maxTime:  7,
			want:     6,
		},
		"hold 7, max 7": {
			holdTime: 7,
			maxTime:  7,
			want:     0,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := wouldTravel(tt.holdTime, tt.maxTime); got != tt.want {
				t.Errorf("wouldTravel() = %v, want %v", got, tt.want)
			}
		})
	}
}
