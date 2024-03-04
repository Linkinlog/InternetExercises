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
			input: `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`,
			want: 4361,
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
			input: `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`,
			want: 467835,
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

func TestTouches(t *testing.T) {
	tests := map[string]struct {
		input segment
		want  bool
	}{
		"segment touches (diagonal)": {
			input: segment{
				row:   0,
				start: 3,
				end:   3,
				value: "*",
			},
			want: true,
		},
		"segment touches (horizontal)": {
			input: segment{
				row:   1,
				start: 4,
				end:   5,
				value: "42",
			},
			want: true,
		},
		"segment touches (vertical)": {
			input: segment{
				row:   2,
				start: 3,
				end:   4,
				value: "42",
			},
			want: true,
		},
		"segment doesnt touch": {
			input: segment{
				row:   0,
				start: 2,
				end:   2,
				value: "*",
			},
			want: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			seg := segment{
				row:   1,
				start: 4,
				end:   5,
				value: "42",
			}

			if got := seg.touches(tt.input); got != tt.want {
				t.Fatalf("touches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseDigits(t *testing.T) {
	cases := map[string]struct {
		input string
		want  []segment
	}{
		"single line, no digits": {
			input: "..........",
			want:  []segment{},
		},
		"single line, one number(one digit)": {
			input: "4.........",
			want: []segment{
				{
					row:   0,
					start: 0,
					end:   0,
					value: "4",
				},
			},
		},
		"single line, one number(multiple digits)": {
			input: "467.......",
			want: []segment{
				{
					row:   0,
					start: 0,
					end:   2,
					value: "467",
				},
			},
		},
		"single line, multiple numbers(multiple digits)": {
			input: "467..114..",
			want: []segment{
				{
					row:   0,
					start: 0,
					end:   2,
					value: "467",
				},
				{
					row:   0,
					start: 5,
					end:   7,
					value: "114",
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := parseDigits(tc.input, 0)
			if len(got) != len(tc.want) {
				t.Fatalf("parseDigits() = %v, want %v", got, tc.want)
			}

			for i, want := range tc.want {
				if got[i].row != want.row {
					t.Errorf("parseDigits()[%d].row = %d, want %d", i, got[i].row, want.row)
				}

				if got[i].start != want.start {
					t.Errorf("parseDigits()[%d].start = %d, want %d", i, got[i].start, want.start)
				}

				if got[i].end != want.end {
					t.Errorf("parseDigits()[%d].end = %d, want %d", i, got[i].end, want.end)
				}

				if got[i].value != want.value {
					t.Errorf("parseDigits()[%d].value = %s, want %s", i, got[i].value, want.value)
				}
			}
		})
	}
}

func TestParseParts(t *testing.T) {
	cases := map[string]struct {
		input string
		want  []segment
	}{
		"single line, no parts": {
			input: "..........",
			want:  []segment{},
		},
		"single line, one part": {
			input: "*.........",
			want: []segment{
				{
					row:   0,
					start: 0,
					end:   0,
					value: "*",
				},
			},
		},
		"single line, multiple parts": {
			input: "*....#....",
			want: []segment{
				{
					row:   0,
					start: 0,
					end:   0,
					value: "*",
				},
				{
					row:   0,
					start: 5,
					end:   5,
					value: "#",
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := parseParts(tc.input, 0)
			if len(got) != len(tc.want) {
				t.Fatalf("parseParts() = %v, want %v", got, tc.want)
			}

			for i, want := range tc.want {
				if got[i].row != want.row {
					t.Errorf("parseParts()[%d].row = %d, want %d", i, got[i].row, want.row)
				}

				if got[i].start != want.start {
					t.Errorf("parseParts()[%d].start = %d, want %d", i, got[i].start, want.start)
				}

				if got[i].end != want.end {
					t.Errorf("parseParts()[%d].end = %d, want %d", i, got[i].end, want.end)
				}

				if got[i].value != want.value {
					t.Errorf("parseParts()[%d].value = %s, want %s", i, got[i].value, want.value)
				}
			}
		})
	}
}

func TestIsGear(t *testing.T) {
	// 467..114..
	// ...*......
	// ..35..633.
	digits := []segment{
		{
			row:   0,
			start: 0,
			end:   2,
			value: "467",
		},
		{
			row:   0,
			start: 5,
			end:   7,
			value: "114",
		},
		{
			row:   2,
			start: 2,
			end:   4,
			value: "35",
		},
		{
			row:   2,
			start: 7,
			end:   9,
			value: "633",
		},
	}

	cases := map[string]struct {
		input []segment
		want  bool
	}{
		"no parts": {
			input: []segment{
				{
					row:   0,
					start: 0,
					end:   0,
					value: "",
				},
			},
			want: false,
		},
		"gear": {
			input: []segment{
				{
					row:   1,
					start: 3,
					end:   3,
					value: "*",
				},
			},
			want: true,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if got := tc.input[0].isGear(digits); got != tc.want {
				t.Fatalf("isGear() = %v, want %v", got, tc.want)
			}
		})
	}
}
