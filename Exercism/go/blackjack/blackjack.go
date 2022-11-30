package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	cards := map[string]int{
		"ace":   11,
		"king":  10,
		"queen": 10,
		"jack":  10,
		"ten":   10,
		"nine":  9,
		"eight": 8,
		"seven": 7,
		"six":   6,
		"five":  5,
		"four":  4,
		"three": 3,
		"two":   2,
	}
	return cards[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	// Convert card values to their int values for comparisons
	card1Val, card2Val := ParseCard(card1), ParseCard(card2)

	dealerDangerArray := [2]int{
		11,
		10,
	}
	// If both cards are an Ace, we split
	if card1Val == card2Val && card1Val == 11 {
		return "P"
	}

	// If we have a blackjack, we win
	if card1Val+card2Val == 21 {
		// Make a temporary var to avoid edge cases while looping
		// Also helps ensure we are not passing a W when we shouldn't
		res := "W"
		for _, e := range dealerDangerArray {
			// Unless the dealer has a ten/face/ace
			if ParseCard(dealerCard) == e {
				res = "S"
				return res
			}
		}
		return res
	}

	// If our cards total between 17-20 we stand
	if val := card1Val + card2Val; val >= 17 && val <= 20 {
		return "S"
	} else if val >= 12 && val <= 16 && ParseCard(dealerCard) >= 7 {
		// If our cards total between 12-16 and the dealers card is over 7, we hit
		return "H"
	} else if val >= 12 && val <= 16 {
		// If the previous is true except the dealers card, we stand
		return "S"
	}

	// Anything else will ba a card <=11 so we hit
	return "H"
}
