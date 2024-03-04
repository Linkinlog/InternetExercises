package main

import (
	"regexp"
	"strconv"
	"strings"
)

func part1(input []string) int {
	games := parseGames(input)
	sum := 0

	for _, g := range games {
		score := 0
		for i := 0; i < g.score(); i++ {
			if score == 0 {
				score = 1
				continue
			}
			score *= 2
		}

		sum += score
	}

	return sum
}

func part2(input []string) int {
	games := parseGames(input)

	for _, g := range games {
		if g.score() == 0 {
			continue // skip the losers lmao
		}
		for i := g.id + 1; i <= g.id+g.score(); i++ {
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

func (g game) winners() []int {
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

func (g game) score() int {
	return len(g.winners())
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
