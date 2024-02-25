package main

import "testing"

var givenInput string = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

var testCases = []struct {
	given         string
	part1Expected int
	part2Expected int
}{
	{
		givenInput,
		157,
		70,
	},
}

func TestHelloWorld(t *testing.T) {
	for _, e := range testCases {
		t.Run("Part1", func(t *testing.T) {
			if res := Day3Part1(e.given); res != e.part1Expected {
				t.Fatalf("Got %d, wanted %d", res, e.part1Expected)
			}
		})
		t.Run("Part2", func(t *testing.T) {
			if res := Day3Part2(e.given); res != e.part2Expected {
				t.Fatalf("Got %d, wanted %d", res, e.part2Expected)
			}
		})
	}
}
