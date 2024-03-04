package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"whole shebang": {
			input: `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`,
			want: 35,
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
			input: `seeds: 79 14 55 13 79 14

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`,
			want: 46,
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

// seeds: 79 14 55 13
//
// seed-to-soil map:
// 50 98 2
// 52 50 48
//
// soil-to-fertilizer map:
// 0 15 37
// 37 52 2
// 39 0 15
//
// fertilizer-to-water map:
// 49 53 8
// 0 11 42
// 42 0 7
// 57 7 4
//
// water-to-light map:
// 88 18 7
// 18 25 70
//
// light-to-temperature map:
// 45 77 23
// 81 45 19
// 68 64 13
//
// temperature-to-humidity map:
// 0 69 1
// 1 0 69
//
// humidity-to-location map:
// 60 56 37
// 56 93 4
func newAlm() almanac {
	return almanac{
		seeds: []int{79, 14, 55, 13},
		seedToSoilMaps: []conversionMap{
			{destination: 50, source: 98, width: 2},
			{destination: 52, source: 50, width: 48},
		},
		soilToFertMaps: []conversionMap{
			{destination: 0, source: 15, width: 37},
			{destination: 37, source: 52, width: 2},
			{destination: 39, source: 0, width: 15},
		},
		fertToWaterMaps: []conversionMap{
			{destination: 49, source: 53, width: 8},
			{destination: 0, source: 11, width: 42},
			{destination: 42, source: 0, width: 7},
			{destination: 57, source: 7, width: 4},
		},
		waterToLightMaps: []conversionMap{
			{destination: 88, source: 18, width: 7},
			{destination: 18, source: 25, width: 70},
		},
		lightToTempMaps: []conversionMap{
			{destination: 45, source: 77, width: 23},
			{destination: 81, source: 45, width: 19},
			{destination: 68, source: 64, width: 13},
		},
		tempToHumMaps: []conversionMap{
			{destination: 0, source: 69, width: 1},
			{destination: 1, source: 0, width: 69},
		},
		HumToLocMaps: []conversionMap{
			{destination: 60, source: 56, width: 37},
			{destination: 56, source: 93, width: 4},
		},
	}
}

func TestParse(t *testing.T) {
	// t.Skip("So like theres this thing called aliens and theyll change your life")
	tests := map[string]struct {
		input []string
		want  almanac
	}{
		"whole shebang": {
			input: strings.Split(`seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`, "\n"),
			want: newAlm(),
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := parse(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseSeeds(t *testing.T) {
	tests := map[string]struct {
		input string
		want  []int
	}{
		"whole shebang": {
			input: `seeds: 79 14 55 13`,
			want:  []int{79, 14, 55, 13},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := parseSeederoonies(tt.input)
			if len(got) != len(tt.want) {
				t.Fatal("didnt work")
			}
			for i, v := range got {
				if v != tt.want[i] {
					t.Fatal("didnt work")
				}
			}
		})
	}
}

func TestParseSeedToSoil(t *testing.T) {
	tests := map[string]struct {
		input string
		want  conversionMap
	}{
		"whole shebang": {
			input: `50 98 2`,
			want: conversionMap{
				destination: 50,
				source:      98,
				width:       2,
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := parseConversionMap(tt.input)
			if got.source != tt.want.source {
				t.Fatal("source mismatch")
			}
			if got.width != tt.want.width {
				t.Fatal("width mismatch")
			}
			if got.destination != tt.want.destination {
				t.Fatal("destination mismatch")
			}
		})
	}
}

func TestConvert(t *testing.T) {
	a := newAlm()
	tests := map[string]struct {
		source int
		cMap   *[]conversionMap
		want  int
	}{
		"one": {
			source: 78,
			cMap: &a.HumToLocMaps,
			want:  82,
		},
		"two": {
			source: 43,
			cMap: &a.HumToLocMaps,
			want:  43,
		},
		"exclusive width": {
			source: 1,
			cMap: &[]conversionMap{
				{destination: 1, source: 0, width: 1},
			},
			want:  1,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := a.convert(tt.source, tt.cMap); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
