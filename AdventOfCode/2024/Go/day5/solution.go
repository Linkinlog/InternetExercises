package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func part1(input io.Reader) any {
	s := bufio.NewScanner(input)

	ruleMap := getRules(s)

	sum := 0

	for s.Scan() {
		valid := true

		nums := strings.Split(s.Text(), ",")

		for i, strnum := range nums {
			num, _ := strconv.Atoi(strnum)

			if !check(nums[:i], ruleMap[num]) {
				valid = false
			}
		}
		if valid {
			num, _ := strconv.Atoi(nums[len(nums)/2])
			sum += num
		}
	}

	if s.Err() != nil {
		panic(s.Err())
	}

	return sum
}

func part2(input []string) any {
	_ = input
	return nil
}

type ruleMap map[int][]int

func getRules(s *bufio.Scanner) (r ruleMap) {
	r = make(map[int][]int)

	for s.Scan() {
		if s.Text() == "" {
			return r
		}

		x, y := 0, 0
		if _, err := fmt.Sscanf(s.Text(), "%d|%d", &x, &y); err != nil {
			panic(err)
		}
		r[x] = append(r[x], y)

	}

	return r
}

func check(nums []string, boundaries []int) bool {
	if len(boundaries) == 0 {
		return true
	}

	for _, strnum := range nums {
		for _, boundary := range boundaries {
			num, _ := strconv.Atoi(strnum)
			if num == boundary {
				return false
			}
		}
	}

	return true
}
