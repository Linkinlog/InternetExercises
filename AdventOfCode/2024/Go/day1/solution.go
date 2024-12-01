package main

import (
	"fmt"
	"math"
	"slices"
)

func part1(input []string) any {
	listA := make([]int, len(input))
	listB := make([]int, len(input))

	for i, line := range input {
		var a, b int
		if _, err := fmt.Sscanf(line, "%d   %d", &a, &b); err != nil {
			panic(err)
		}

		listA[i] = a
		listB[i] = b

	}

	if len(listA) != len(listB) {
		panic("lists uneven")
	}

	slices.Sort(listA)
	slices.Sort(listB)

	sum := 0
	for i := range listA {
		sum += int(math.Abs(float64(listA[i] - listB[i])))
	}

	return sum
}

func part2(input []string) any {
	listA := make([]int, len(input))
	listB := make(map[int]int)

	for i, line := range input {
		var a, b int
		if _, err := fmt.Sscanf(line, "%d   %d", &a, &b); err != nil {
			panic(err)
		}

		listA[i] = a

		if _, ok := listB[b]; !ok {
			listB[b] = 0
		}

		listB[b] += 1

	}

	sum := 0

	for i := range listA {
		if val, ok := listB[listA[i]]; ok {
			sum += listA[i] * val
		}
	}

	return sum
}
