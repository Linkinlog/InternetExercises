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
		if checkLine(s.Text(), ruleMap) {
			nums := strings.Split(s.Text(), ",")
			num, _ := strconv.Atoi(nums[len(nums)/2])
			sum += num
		}
	}

	if s.Err() != nil {
		panic(s.Err())
	}

	return sum
}

func part2(input io.Reader) any {
	s := bufio.NewScanner(input)

	ruleMap := getRules(s)

	sum := 0

	failures := []string{}

	for s.Scan() {
		failure := s.Text()
		nums := strings.Split(s.Text(), ",")
		for range len(nums) * len(nums) {
			if !checkLine(failure, ruleMap) {
				for i, c := range nums {
					num, _ := strconv.Atoi(c)
					for ii := len(nums[:i]) - 1; ii >= 0; ii-- {
						for _, boundary := range ruleMap[num] {
							if !check(nums[:i], boundary) {

								// fmt.Printf("moving %s to the end\n", nums[ii])
								el := nums[ii]
								nums = append(nums[:ii], nums[ii+1:]...)
								nums = append(nums, el)
								failure = strings.Join(nums, ",")
								// fmt.Println(failure)
								failures = append(failures, failure)
							}
						}
						// fmt.Println()
					}
				}
			}
		}
	}

	for _, failure := range failures {
		if checkLine(failure, ruleMap) {
			nums := strings.Split(failure, ",")
			num, _ := strconv.Atoi(nums[len(nums)/2])
			sum += num
		}
	}

	if s.Err() != nil {
		panic(s.Err())
	}

	return sum
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

func check(nums []string, boundary int) bool {
	for _, strnum := range nums {
		num, _ := strconv.Atoi(strnum)
		if num == boundary {
			return false
		}
	}

	return true
}

func checkLine(line string, ruleMap ruleMap) bool {
	valid := true

	nums := strings.Split(line, ",")

	for i, strnum := range nums {
		num, _ := strconv.Atoi(strnum)

		for _, boundary := range ruleMap[num] {
			if !check(nums[:i], boundary) {
				valid = false
			}
		}
	}
	return valid
}
