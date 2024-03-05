package main

import (
	"regexp"
	"strconv"
	"strings"
)

func part1(input []string) int {
	races := parse(input)
	waysToWin := []int{}
	for _, r := range races {
		winningTimes := []int{}
		for i := 1; i <= r.time; i++ {
			if wouldTravel(i, r.time) > r.record {
				winningTimes = append(winningTimes, i)
			}
		}
		waysToWin = append(waysToWin, len(winningTimes))
	}

	sum := 0
	for _, w := range waysToWin {
		if sum == 0 {
			sum = w
			continue
		}

		sum *= w
	}

	return sum
}

func part2(input []string) int {
	r := parseMkII(input)
	waysToWin := []int{}
	winningTimes := []int{}
	for i := 1; i <= r.time; i++ {
		if wouldTravel(i, r.time) > r.record {
			winningTimes = append(winningTimes, i)
		}
	}
	waysToWin = append(waysToWin, len(winningTimes))

	sum := 0
	for _, w := range waysToWin {
		if sum == 0 {
			sum = w
			continue
		}

		sum *= w
	}

	return sum
}

type race struct {
	time   int
	record int
}

func wouldTravel(holdTime, maxTime int) int {
	return holdTime * (maxTime - holdTime)
}

func parse(input []string) []race {
	digitPat := regexp.MustCompile(`\d+`)
	time := digitPat.FindAllString(strings.Split(input[0], ":")[1], -1)
	record := digitPat.FindAllString(strings.Split(input[1], ":")[1], -1)
	if len(time) != len(record) {
		panic("input error")
	}
	races := []race{}
	for i := range time {
		timeInt, _ := strconv.Atoi(time[i])
		recordInt, _ := strconv.Atoi(record[i])
		races = append(races, race{time: timeInt, record: recordInt})
	}

	return races
}

func parseMkII(input []string) race {
	digitPat := regexp.MustCompile(`\d+`)
	times := digitPat.FindAllString(strings.Split(input[0], ":")[1], -1)
	records := digitPat.FindAllString(strings.Split(input[1], ":")[1], -1)
	if len(times) != len(records) {
		panic("input error")
	}

	time, _ := strconv.Atoi(strings.Join(times, ""))
	record, _ := strconv.Atoi(strings.Join(records, ""))

	return race{time: time, record: record}
}
