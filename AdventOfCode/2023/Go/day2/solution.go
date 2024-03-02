package main

import (
	"regexp"
	"strconv"
	"strings"
)

func part1(input []string, maxy subset) int {
	sum := 0

	for _, l := range input {
		g := parseGame(l)
		if g.possibleWith(maxy) {
			sum += g.id
		}
	}

	return sum
}

func part2(input []string, maxy subset) int {
	sum := 0

	for _, l := range input {
		g := parseGame(l)
		sum += g.lowestPossible().power()
	}

	return sum
}

type subset struct {
	red   int
	green int
	blue  int
}

func (s subset) power() int {
	return s.red * s.green * s.blue
}

type Game struct {
	id      int
	Subsets []subset
}

func parseGame(s string) *Game {
	digitPattern := `\d+`
	reg := regexp.MustCompile(digitPattern)
	game := &Game{}

	segments := strings.Split(s, ":")
	if len(segments) < 2 {
		return game
	}
	game.id, _ = strconv.Atoi(reg.FindString(segments[0]))

	subsets := strings.Split(segments[1], ";")
	for _, sub := range subsets {
		s := subset{}
		colors := strings.Split(sub, ",")
		for _, color := range colors {
			color = strings.ToLower(color)
			if strings.Contains(color, "red") {
				s.red, _ = strconv.Atoi(reg.FindString(color))
			}
			if strings.Contains(color, "green") {
				s.green, _ = strconv.Atoi(reg.FindString(color))
			}
			if strings.Contains(color, "blue") {
				s.blue, _ = strconv.Atoi(reg.FindString(color))
			}
		}
		game.Subsets = append(game.Subsets, s)
	}
	return game
}

func (g *Game) possibleWith(maxSub subset) bool {
	for _, sub := range g.Subsets {
		if maxSub.red < sub.red || maxSub.blue < sub.blue || maxSub.green < sub.green {
			return false
		}
	}
	return true
}

func (g *Game) lowestPossible() subset {
	if len(g.Subsets) < 1 {
		return subset{}
	}
	s := subset{
		red:   g.Subsets[0].red,
		green: g.Subsets[0].green,
		blue:  g.Subsets[0].blue,
	}

	for _, sub := range g.Subsets {
		if sub.blue > s.blue {
			s.blue = sub.blue
		}
		if sub.green > s.green {
			s.green = sub.green
		}
		if sub.red > s.red {
			s.red = sub.red
		}
	}

	return s
}
