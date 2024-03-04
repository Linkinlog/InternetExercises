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
			input: `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`,
			want: 13,
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
			input: `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`,
			want: 30,
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

func TestParseGame(t *testing.T) {
	tests := map[string]struct {
		input string
		want  game
	}{
		"parses game": {
			input: "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			want: game{
				id:          1,
				winningNums: []int{41, 48, 83, 86, 17},
				ourNums:     []int{83, 86, 6, 31, 17, 9, 48, 53},
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := parseGame(tt.input)
			if got.id != tt.want.id {
				t.Errorf("id = %v, want %v", got.id, tt.want.id)
			}
			for i, wn := range tt.want.winningNums {
				if got.winningNums[i] != wn {
					t.Errorf("winningNums[%d] = %v, want %v", i, got.winningNums[i], wn)
				}
			}
			for i, on := range tt.want.ourNums {
				if got.ourNums[i] != on {
					t.Errorf("ourNums[%d] = %v, want %v", i, got.ourNums[i], on)
				}
			}
		})
	}
}

func TestOurWinners(t *testing.T) {
	tests := map[string]struct {
		game game
		want []int
	}{
		"returns our winners": {
			game: game{
				id:          1,
				winningNums: []int{41, 48, 83, 86, 17},
				ourNums:     []int{83, 86, 6, 31, 17, 9, 48, 53},
			},
			want: []int{83, 86, 17, 48},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.game.winners()
			if len(got) != len(tt.want) {
				t.Errorf("len(got) = %v, want %v", len(got), len(tt.want))
			}
			for i, wn := range tt.want {
				if got[i] != wn {
					t.Errorf("got[%d] = %v, want %v", i, got[i], wn)
				}
			}
		})
	}
}

func TestParseGames(t *testing.T) {
	cases := map[string]struct {
		input []string
		want  []game
	}{
		"3 games, 7 total instances": {
			input: []string{
				"Card 1: 41 48 83 86 17 | 83 87  6 31 16  9 48 53",
				"Card 2: 13 32 20 16 61 | 62 30 68 82 17 32 24 19",
				"Card 3:  1 21 53 59 44 | 69 82 63 72 16 20 14  2",
			},
			want: []game{
				{
					id:          1,
					winningNums: []int{41, 48, 83, 86, 17},
					ourNums:     []int{83, 87, 6, 31, 16, 9, 48, 53},
				},
				{
					id:          2,
					winningNums: []int{13, 32, 20, 16, 61},
					ourNums:     []int{62, 30, 68, 82, 17, 32, 24, 19},
				},
				{
					id:          3,
					winningNums: []int{1, 21, 53, 59, 44},
					ourNums:     []int{69, 82, 63, 72, 16, 20, 14, 2},
				},
			},
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			_ = parseGames(tt.input)
		})
	}
}
