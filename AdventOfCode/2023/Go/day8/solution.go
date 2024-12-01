package main

import (
	"strings"
)

func part1(input []string) int {
	dirs, lookups := parseInput(input)
	return travel(lookups, dirs, "AAA", "ZZZ")
}

func part2(input []string) int {
	return 0
}

const (
	LEFT  = 'L'
	RIGHT = 'R'
)

func parseDirections(input string) []rune {
	return []rune(input)
}

func parseLookups(input []string) map[string][]string {
	lookups := make(map[string][]string)
	for _, line := range input {
		if line == "" {
			continue
		}
		splitsville := strings.Split(line, " = ")
		noParens := strings.TrimSuffix(strings.TrimPrefix(splitsville[1], "("), ")")
		children := strings.Split(noParens, ", ")
		lookups[splitsville[0]] = children
	}
	return lookups
}

func parseInput(input []string) ([]rune, map[string][]string) {
	return parseDirections(input[0]), parseLookups(input[1:])
}

func travel(lookups map[string][]string, dirs []rune, current, end string) int {
	sum := 0
	for current != end {
		for _, dir := range dirs {
			sum += 1
			if dir == LEFT {
				current = lookups[current][0]
			} else {
				current = lookups[current][1]
			}
		}
	}
	return sum
}
