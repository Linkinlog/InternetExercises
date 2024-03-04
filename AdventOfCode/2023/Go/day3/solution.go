package main

import (
	"regexp"
	"strconv"
)

func part1(input []string) int {
	digits := []segment{}
	parts := []segment{}

	for r, s := range input {
		digits = append(digits, parseDigits(s, r)...)
		parts = append(parts, parseParts(s, r)...)
	}
	res := []segment{}

	for _, d := range digits {
		for _, p := range parts {
			if d.touches(p) {
				res = append(res, d)
			}
		}
	}

	sum := 0
	for _, d := range res {
		val, _ := strconv.Atoi(d.value)
		sum += val
	}

	return sum
}

func part2(input []string) int {
	digits := []segment{}
	parts := []segment{}

	for r, s := range input {
		digits = append(digits, parseDigits(s, r)...)
		parts = append(parts, parseParts(s, r)...)
	}

	sum := 0

	for _, p := range parts {
		if p.isGear(digits) {
			digiemon := p.gearDigits(digits)
			val1, _ := strconv.Atoi(digiemon[0].value)
			val2, _ := strconv.Atoi(digiemon[1].value)
			sum += (val1 * val2)
		}
	}

	return sum
}

type segment struct {
	row   int
	start int
	end   int
	value string
}

func (s segment) gearDigits(digits []segment) []segment {
	if s.value != "*" {
		return []segment{}
	}
	touching := []segment{}
	for _, d := range digits {
		if s.touches(d) {
			touching = append(touching, d)
		}
	}
	if len(touching) == 2 {
		return touching
	}
	return []segment{}
}

// :(
func (s segment) isGear(digits []segment) bool {
	if s.value != "*" {
		return false
	}
	touching := 0
	for _, d := range digits {
		if s.touches(d) {
			touching += 1
		}
	}
	return touching == 2
}

func parseDigits(s string, row int) []segment {
	digitPat := regexp.MustCompile(`\d+`)
	sDigits := digitPat.FindAllIndex([]byte(s), -1)
	digits := []segment{}

	for _, sDigit := range sDigits {
		digits = append(digits, segment{
			row:   row,
			start: sDigit[0],
			end:   sDigit[1] - 1,
			value: s[sDigit[0]:sDigit[1]],
		})
	}

	return digits
}

func parseParts(s string, row int) []segment {
	partPat := regexp.MustCompile(`[^\d\w\s.]`)
	sParts := partPat.FindAllIndex([]byte(s), -1)
	parts := []segment{}

	for _, sDigit := range sParts {
		parts = append(parts, segment{
			row:   row,
			start: sDigit[0],
			end:   sDigit[1] - 1,
			value: s[sDigit[0]:sDigit[1]],
		})
	}

	return parts
}

func (s segment) touches(x segment) bool {
	if x.start > x.end {
		return false
	}

	if s.row+1 < x.row || s.row-1 > x.row {
		return false
	}

	if x.start > s.end+1 && x.end > s.end+1 {
		return false
	}

	if x.end < s.start-1 && x.start < s.start-1 {
		return false
	}

	return true
}
