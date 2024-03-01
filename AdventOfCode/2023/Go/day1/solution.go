package main

import (
	"regexp"
	"strconv"
)

func part1(input []string) any {
	sum := 0
	for _, l := range input {
		i, err := strconv.Atoi(extractSumFromLine(l))
		if err != nil {
			return err
		}
		sum += i
	}
	return sum
}

func part2(input []string) any {
	return part1(input)
}

func mapToDigit(s string) string {
	wordDigit := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"enin":  "9",
		"thgie": "8",
		"neves": "7",
		"xis":   "6",
		"evif":  "5",
		"ruof":  "4",
		"eerht": "3",
		"owt":   "2",
		"eno":   "1",
	}

	if v, ok := wordDigit[s]; ok {
		return v
	}
	return ""
}

func extractSumFromLine(l string) string {
	num1, num2 := "", ""
	pat := regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine|enin|thgie|neves|xis|evif|ruof|eerht|owt|eno`)
	numBytes1 := pat.Find([]byte(l))
	if len(numBytes1) <= 0 {
		return ""
	}
	if mapToDigit(string(numBytes1)) != "" {
		num1 = mapToDigit(string(numBytes1))
	} else {
		num1 = string(numBytes1)
	}
	numBytes2 := pat.Find([]byte(reverse(l)))
	if len(numBytes2) <= 0 {
		return ""
	}
	if mapToDigit(string(numBytes2)) != "" {
		num2 = mapToDigit(string(numBytes2))
	} else {
		num2 = string(numBytes2)
	}

	return num1 + num2
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
