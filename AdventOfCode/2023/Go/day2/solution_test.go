package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	tests := map[string]struct {
		input     string
		maxSubset subset
		want      int
	}{
		"one row": {
			input: `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green`,
			maxSubset: subset{
				red:   12,
				green: 13,
				blue:  14,
			},
			want: 1,
		},
		"whole shebang": {
			input: `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`,
			maxSubset: subset{
				red:   12,
				green: 13,
				blue:  14,
			},
			want: 8,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := part1(strings.Split(tt.input, "\n"), tt.maxSubset); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := map[string]struct {
		input string
		maxSubset subset
		want  int
	}{
		"one row": {
			input: `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green`,
			maxSubset: subset{
				red:   12,
				green: 13,
				blue:  14,
			},
			want: 4 * 6 * 2,
		},
		"whole shebang": {
			input: `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`,
			maxSubset: subset{
				red:   12,
				green: 13,
				blue:  14,
			},
			want: 2286,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := part2(strings.Split(tt.input, "\n"), tt.maxSubset); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseGame(t *testing.T) {
	tests := map[string]struct {
		input string
		want  *Game
	}{
		"base": {
			input: "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red",
			want: &Game{
				id: 5,
				Subsets: []subset{
					{
						red:   6,
						blue:  1,
						green: 3,
					},
					{
						red:  1,
						blue: 2,
					},
				},
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := parseGame(tt.input)
			if got.id != tt.want.id {
				t.Errorf("parseGame().id = %v, want %v", got.id, tt.want.id)
			}

			if len(tt.want.Subsets) != len(got.Subsets) {
				t.Errorf("len(tt.want.Subsets) == %d, want %d", len(tt.want.Subsets), len(got.Subsets))
			}

			for i, s := range got.Subsets {
				if s.red != tt.want.Subsets[i].red {
					t.Errorf("parseGame().Subsets[%d].red = %v, want %v", i, s.red, tt.want.Subsets[i].red)
				}
				if s.blue != tt.want.Subsets[i].blue {
					t.Errorf("parseGame().Subsets[%d].blue = %v, want %v", i, s.blue, tt.want.Subsets[i].blue)
				}
				if s.green != tt.want.Subsets[i].green {
					t.Errorf("parseGame().Subsets[%d].green = %v, want %v", i, s.green, tt.want.Subsets[i].green)
				}
			}
		})
	}
}

func TestPossibleWith(t *testing.T) {
	tests := map[string]struct {
		input subset
		want  bool
	}{
		"base": {
			input: subset{
				red:   12,
				green: 13,
				blue:  14,
			},
			want: true,
		},
		"not enough red": {
			input: subset{
				red:   4,
				green: 13,
				blue:  14,
			},
			want: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			game := &Game{
				id: 1,
				Subsets: []subset{
					{
						blue: 5,
						red:  5,
					},
					{
						red:   5,
						green: 5,
						blue:  5,
					},
					{
						green: 5,
					},
				},
			}

			if got := game.possibleWith(tt.input); got != tt.want {
				t.Errorf("possibleWith() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowestPossible(t *testing.T) {
	tests := map[string]struct {
		subsets []subset
		want    subset
	}{
		"base": {
			subsets: []subset{
				{
					red:   12,
					green: 13,
					blue:  14,
				},
			},
			want: subset{
				red:   12,
				green: 13,
				blue:  14,
			},
		},
		"two": {
			subsets: []subset{
				{
					red:   12,
					green: 13,
					blue:  14,
				},
				{
					red:   22,
					green: 3,
					blue:  24,
				},
			},
			want: subset{
				red:   22,
				green: 13,
				blue:  24,
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			game := &Game{
				id:      1,
				Subsets: tt.subsets,
			}

			if got := game.lowestPossible(); got != tt.want {
				t.Errorf("lowestPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}
