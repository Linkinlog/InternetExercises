package main

import (
	"bufio"
	"fmt"
	"strings"
)

func RockPaperScissors(list string) int {
	var score int
	var win int = 6
	var draw int = 3
	scanner := bufio.NewScanner(strings.NewReader(list))
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		x, y := RockPaperScissorsMap(fields[0]), RockPaperScissorsMap(fields[1])
		if x == y {
			score += scoresmap(y) + draw
		}
		if x == "paper" {
			switch y {
			case "scissors":
				score += scoresmap(y) + win
			case "rock":
				score += scoresmap(y)
			}
		}
		if x == "scissors" {
			switch y {
			case "rock":
				score += scoresmap(y) + win
			case "paper":
				score += scoresmap(y)

			}
		}
		if x == "rock" {
			switch y {
			case "paper":
				score += scoresmap(y) + win
			case "scissors":
				score += scoresmap(y)
			}
		}
	}
	return score
}

func Getpart2(str string) int {
	var score int
	var win int = 6
	var draw int = 3
	scanner := bufio.NewScanner(strings.NewReader(str))
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		x, y := RockPaperScissorsMap(fields[0]), part2map(fields[1])
		fmt.Print()
		if x == "paper" {
			switch y {
			case "lose":
				score += scoresmap("rock")
			case "win":
				score += scoresmap("scissors") + win
			case "draw":
				score += scoresmap(x) + draw
			}

		}
		if x == "scissors" {
			switch y {
			case "lose":
				score += scoresmap("paper")
			case "win":
				score += scoresmap("rock") + win
			case "draw":
				score += scoresmap(x) + draw
			}
		}
		if x == "rock" {
			switch y {
			case "lose":
				score += scoresmap("scissors")
			case "draw":
				score += scoresmap(x) + draw
			case "win":
				score += scoresmap("paper") + win
			}
		}
	}
	return score
}

func part2map(str string) string {
	str = strings.ToLower(str)
	part2map := map[string]string{
		"x": "lose",
		"y": "draw",
		"z": "win",
	}
	return part2map[str]
}

func scoresmap(str string) int {
	scoresmap := map[string]int{
		"rock":     1,
		"paper":    2,
		"scissors": 3,
	}

	return scoresmap[str]
}

func RockPaperScissorsMap(str string) string {
	str = strings.ToLower(str)
	rpsmap := map[string]string{
		"x": "rock",
		"y": "paper",
		"z": "scissors",
		"a": "rock",
		"b": "paper",
		"c": "scissors",
	}
	return rpsmap[str]
}

func main() {
	// 	fmt.Print(Getpart2(test1Input))
	//	fmt.Print(Getpart2(testinput()))
}
