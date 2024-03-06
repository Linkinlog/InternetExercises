package main

import (
	"fmt"
	"strconv"
	"strings"
)

func part1(input []string) any {
	g := parseInput(input)
	g.rankHands(false)

	total := 0
	// multiply each hands bet by its rank
	for i, h := range g.hands {
		total += h.bet * (i + 1)
	}

	return total
}

func part2(input []string) any {
	g := parseInput(input)
	g.rankHands(true)

	total := 0
	// multiply each hands bet by its rank
	for i, h := range g.hands {
		total += h.bet * (i + 1)
	}

	return total
}

type Game struct {
	hands []*hand
}

func (g Game) rankHands(part2 bool) {
	part := 1
	if part2 {
		part = 2
	}
	// basic sort, based on types
	for i := 0; i < len(g.hands); i++ {
		for j := i + 1; j < len(g.hands); j++ {
			if part == 1 {
				if g.hands[i].getType() > g.hands[j].getType() {
					g.hands[i], g.hands[j] = g.hands[j], g.hands[i]
				}
			} else {
				if g.hands[i].getType2() > g.hands[j].getType2() {
					g.hands[i], g.hands[j] = g.hands[j], g.hands[i]
				}
			}
		}
	}

	// more in-depth sort, based on each card
	for i := 0; i < len(g.hands); i++ {
		for j := i + 1; j < len(g.hands); j++ {
			if part == 1 {
				if g.hands[i].getType() != g.hands[j].getType() {
					continue
				}
			} else {
				if g.hands[i].getType2() != g.hands[j].getType2() {
					continue
				}
			}
			if equal(compare(g.hands[i].cards, g.hands[j].cards), g.hands[i].cards) {
				g.hands[i], g.hands[j] = g.hands[j], g.hands[i]
			}
		}
	}

}

func (g Game) String() string {
	strs := []string{}
	for i, h := range g.hands {
		kinds := []string{}
		for _, c := range h.cards {
			kinds = append(kinds, c.kind)
		}
		strs = append(strs, fmt.Sprintf("hand %d: %v %v\n", i, kinds, h.getType2()))

	}

	return strings.Join(strs, "")
}

type hand struct {
	cards []*card
	bet   int
}

type handType int

const (
	highCard handType = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

func (h hand) String() string {
	kinds := []string{}
	for _, c := range h.cards {
		kinds = append(kinds, c.kind)
	}
	return strings.Join(kinds, "")
}

// `type` is reserved (shrug emoji)
func (h hand) getType2() handType {
	kindCount := map[string]int{}
	for _, c := range h.cards {
		kindCount[c.kind]++
	}
	currentType := h.getType()

	t := currentType

	switch kindCount["J"] {
	case 4, 5:
		t = fiveOfAKind
	case 3:
		t = fourOfAKind
		if len(kindCount) == 2 {
			t = fiveOfAKind
			break
		}
	case 2:
		t = threeOfAKind
		if currentType == twoPair {
			t = fourOfAKind
			break
		}
		if currentType == onePair {
			t = threeOfAKind
			break
		}
		if currentType == fullHouse {
			t = fiveOfAKind
			break
		}
	case 1:
		if currentType == highCard {
			t = onePair
			break
		}
		if currentType == threeOfAKind {
			t = fourOfAKind
			break
		}
		if currentType == twoPair {
			t = fullHouse
			break
		}
		if currentType == onePair {
			t = threeOfAKind
			break
		}
		if currentType == fourOfAKind {
			t = fiveOfAKind
			break
		}
	}

	return t
}

// `type` is reserved (shrug emoji)
func (h hand) getType() handType {
	kindCount := map[string]int{}
	for _, c := range h.cards {
		kindCount[c.kind]++
	}

	t := highCard

	switch len(kindCount) {
	case 4:
		t = onePair
	case 3:
		t = twoPair
		for _, count := range kindCount {
			if count == 3 {
				t = threeOfAKind
			}
		}
	case 2:
		t = threeOfAKind
		for _, count := range kindCount {
			if count == 4 {
				t = fourOfAKind
			}
			if count == 2 {
				t = fullHouse
			}
		}
	case 1:
		t = fiveOfAKind
	}

	return t
}

type card struct {
	kind string
}

func compare(a, b []*card) []*card {
	for i, c := range a {
		if c.strength() > b[i].strength() {
			return a
		} else if c.strength() < b[i].strength() {
			return b
		}
	}
	return a

}

func equal(a, b []*card) bool {
	// make sure theyre the same
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i].kind != b[i].kind {
			return false
		}
	}
	return true
}

func (c card) strength() int {
	switch c.kind {
	case "A":
		return 13
	case "K":
		return 12
	case "Q":
		return 11
	case "T":
		return 10
	case "J":
		return 1
	default:
		return int(c.kind[0] - '0')
	}
}

func parseInput(input []string) Game {
	hands := []*hand{}
	for _, strGame := range input {
		splitty := strings.Split(strGame, " ")
		strHand := splitty[0]
		bet, _ := strconv.Atoi(splitty[1])

		cards := []*card{}

		for _, strCard := range strings.Split(strHand, "") {
			cards = append(cards, &card{kind: strCard})
		}
		hands = append(hands, &hand{cards: cards, bet: bet})
	}

	return Game{
		hands: hands,
	}
}
