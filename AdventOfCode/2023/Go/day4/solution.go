package main

import (
	"regexp"
	"strconv"
	"strings"
)

func part1(input []string) int {
	sum := 0

	for _, s := range input {
		g := parseGame(s)
		sum += g.calcPartOneAlgo()
	}

	return sum
}

func part2(input []string) int {
	games := parseGames(input)

	for _, g := range games {
		if g.score() == 0 {
			continue
		}
		for i := g.id + 1; i <= g.id+g.score(); i++ {
			if i > len(games) {
				break
			}
			games[i-1].instances += g.instances
		}
	}

	sum := 0

	for _, g := range games {
		sum += g.instances
	}
	return sum
}

type game struct {
	id          int
	winningNums []int
	ourNums     []int
	instances   int
}

func (g game) score() int {
	return len(g.ourWinners())
}

func (g game) ourWinners() []int {
	nums := []int{}
	for _, num := range g.ourNums {
		for _, wNum := range g.winningNums {
			if num == wNum {
				nums = append(nums, num)
			}
		}
	}

	return nums
}

func (g game) calcPartOneAlgo() int {
	winners := g.ourWinners()

	if len(winners) < 1 {
		return 0
	}

	// idk but the internet says its faster
	// TODO learn bitwise
	return 1 << (len(winners) - 1)
}

func parseGames(input []string) []game {
	games := []game{}

	for _, s := range input {
		games = append(games, parseGame(s))
	}

	return games
}

func parseGame(input string) game {
	g := game{
		instances: 1,
	}
	digitPat := regexp.MustCompile(`\d+`)

	halves := strings.Split(input, ":")

	g.id, _ = strconv.Atoi(digitPat.FindString(halves[0]))

	halves = strings.Split(halves[1], "|")

	sNums := digitPat.FindAllString(halves[0], -1)
	for _, sNum := range sNums {
		num, _ := strconv.Atoi(sNum)
		g.winningNums = append(g.winningNums, num)
	}
	sNums = digitPat.FindAllString(halves[1], -1)
	for _, sNum := range sNums {
		num, _ := strconv.Atoi(sNum)
		g.ourNums = append(g.ourNums, num)
	}

	return g
}
