package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	regex        string = `mul\((\d{1,3}),(\d{1,3})\)`
	anotherRegex string = `don't\(\).*?do\(\)`
)

func part1(input []string) any {
	regularGuy := regexp.MustCompile(regex)

	sum := 0
	for _, line := range input {
		instructions := regularGuy.FindAllStringSubmatch(line, -1)

		for _, instruction := range instructions {
			fmt.Println(instruction[1:])
			x, _ := strconv.Atoi(instruction[1])
			y, _ := strconv.Atoi(instruction[2])

			sum += x * y
		}
	}

	return sum
}

func part2(input []string) any {
	line := strings.Join(input, "")
	regularGuy := regexp.MustCompile(regex)

	sum := 0
	coolGuy := regexp.MustCompile(anotherRegex)
		for coolGuy.MatchString(line) {
			line = coolGuy.ReplaceAllString(line, "")
		}
		instructions := regularGuy.FindAllStringSubmatch(line, -1)

		for _, instruction := range instructions {
			x, _ := strconv.Atoi(instruction[1])
			y, _ := strconv.Atoi(instruction[2])

			sum += x * y
		}

	return sum
}
