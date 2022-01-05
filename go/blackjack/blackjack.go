package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	case "ace":
		return 11
	default:
		return 0
	}
}

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
	return ParseCard(card1)+ParseCard(card2) == 21
}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
	if isBlackjack {
		//If you have a Blackjack (two cards that sum up to a value of 21),

		//and the dealer does not have an ace, a figure or a ten then you automatically win.
		if dealerScore >= 10 {
			return "S"
		} else {
			return "W"
		}

		//If the dealer does have any of those cards then you'll have to stand and wait for the reveal of the other card.
	}

	// not blackjack must mean two aces if the score is over 20
	return "P"
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {
	if handScore >= 17 {
		return "S"
	} else if handScore <= 11 {
		return "H"
	} else {
		if dealerScore >= 7 {
			return "H"
		} else {
			return "S"
		}
	}
}
