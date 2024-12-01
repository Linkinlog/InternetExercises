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
			input: `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`,
			want: 6,
		},
		"now with even more shebang": {
			input: `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`,
			want: 2,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := part1(strings.Split(tt.input, "\n")); got != tt.want {
				t.Errorf("() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		// TODO
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := part2(strings.Split(tt.input, "\n")); got != tt.want {
				t.Errorf("() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestParseDirections(t *testing.T) {
	tests := map[string]struct {
		input string
		want  []rune
	}{
		"base": {
			input: "LR",
			want:  []rune{'L', 'R'},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := parseDirections(tt.input)
			if len(got) != len(tt.want) {
				t.Errorf("parseDirections() = %v, want %v", got, tt.want)
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("parseDirections() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestParseLookups(t *testing.T) {
	tests := map[string]struct {
		input []string
		want  map[string][]string
	}{
		"base": {
			input: []string{"AAA = (BBB, BBB)"},
			want:  map[string][]string{"AAA": {"BBB", "BBB"}},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := parseLookups(tt.input)
			if len(got) != len(tt.want) {
				t.Errorf("parseLookups() = %v, want %v", got, tt.want)
			}
			for k, v := range got {
				if len(v) != len(tt.want[k]) {
					t.Errorf("parseLookups() = %v, want %v", got, tt.want)
				}
				for i := range v {
					if v[i] != tt.want[k][i] {
						t.Errorf("parseLookups() = %v, want %v", got, tt.want)
					}
				}
			}
		})
	}
}

func TestParseInput(t *testing.T) {
	tests := map[string]struct {
		input []string
		want  []rune
		want1 map[string][]string
	}{
		"base": {
			input: []string{"RLRL", "AAA = (BBB, BBB)"},
			want:  []rune{'R', 'L', 'R', 'L'},
			want1: map[string][]string{"AAA": {"BBB", "BBB"}},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, got1 := parseInput(tt.input)
			if string(got) != string(tt.want) {
				t.Errorf("parseInput() got = %v, want %v", got, tt.want)
			}
			if len(got1) != len(tt.want1) {
				t.Errorf("parseInput() got1 = %v, want %v", got1, tt.want1)
			}
			for k, v := range got1 {
				if len(v) != len(tt.want1[k]) {
					t.Errorf("parseInput() got1 = %v, want %v", got1, tt.want1)
				}
				for i := range v {
					if v[i] != tt.want1[k][i] {
						t.Errorf("parseInput() got1 = %v, want %v", got1, tt.want1)
					}
				}
			}
		})
	}
}

func TestTravel(t *testing.T) {
	tests := map[string]struct {
		lookups map[string][]string
		dirs    []rune
		current string
		end     string
		want    int
	}{
		// LLR
		//
		// AAA = (BBB, BBB)
		// BBB = (AAA, ZZZ)
		// ZZZ = (ZZZ, ZZZ)
		"second example": {
			lookups: map[string][]string{
				"AAA": {"BBB", "BBB"},
				"BBB": {"AAA", "ZZZ"},
				"ZZZ": {"ZZZ", "ZZZ"},
			},
			dirs:    []rune{'L', 'L', 'R'},
			current: "AAA",
			end:     "ZZZ",
			want:    6,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := travel(tt.lookups, tt.dirs, tt.current, tt.end); got != tt.want {
				t.Errorf("travel() = %v, want %v", got, tt.want)
			}
		})
	}
}
