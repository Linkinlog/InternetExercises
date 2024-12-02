package main

import (
	"strconv"
	"strings"
)

/**
	Rules:
		The levels are either all increasing or all decreasing.
		Any two adjacent levels differ by at least one and at most three.
		If the levels are invalid, we may try to remove one level and check if the remaining levels are valid. (Part 2)
**/

func part1(input []string) any {
	sum := 0

	for _, line := range input {
		levels := extractLevels(line)

		if !isValid(levels) {
			continue
		}

		sum = sum + 1
	}

	return sum
}

func part2(input []string) any {
	sum := 0

	for _, line := range input {
		levels := extractLevels(line)

		if isValid(levels) {
			sum = sum + 1
			continue
		}

		for i := 0; i < len(levels); i++ {
			temp := append(append([]int{}, levels[:i]...), levels[i+1:]...)

			if isValid(temp) {
				sum = sum + 1
				break
			}
		}
	}

	return sum
}

func extractLevels(line string) []int {
	levelSplit := strings.Split(line, " ")

	levels := []int{}
	for _, level := range levelSplit {
		i, err := strconv.Atoi(level)
		if err != nil {
			panic(err)
		}

		levels = append(levels, i)
	}

	return levels
}

func isValid(levels []int) bool {
	if !(isIncreasing(levels) || isDecreasing(levels)) {
		return false
	}

	if !withinRange(levels, 1, 3) {
		return false
	}

	return true
}

func isIncreasing(levels []int) bool {
	previousChar := levels[0]
	if previousChar > levels[len(levels)-1] {
		return false
	}

	for i := 1; i < len(levels); i++ {
		currentChar := levels[i]
		if previousChar > currentChar {
			return false
		}

		previousChar = currentChar
	}

	return true
}

func isDecreasing(levels []int) bool {
	previousChar := levels[0]
	if previousChar < levels[len(levels)-1] {
		return false
	}

	for i := 1; i < len(levels); i++ {
		currentChar := levels[i]
		if previousChar < currentChar {
			return false
		}

		previousChar = currentChar
	}

	return true
}

func withinRange(levels []int, rangeStart, rangeEnd int) bool {
	previousChar := levels[0]
	for i := 1; i < len(levels); i++ {
		currentChar := levels[i]

		diff := int(previousChar) - int(currentChar)
		if diff < 0 {
			diff = -diff
		}

		if diff < rangeStart || diff > rangeEnd {
			return false
		}

		previousChar = currentChar
	}

	return true
}

