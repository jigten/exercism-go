package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "one":
		return 1
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
	case "ten", "jack", "queen", "king":
		return 10
	case "ace":
		return 11
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	card1Val := ParseCard(card1)
	card2Val := ParseCard(card2)
	handSum := card1Val + card2Val
	dealerCardValue := ParseCard(dealerCard)

	switch {
	case card1Val == 11 && card2Val == 11:
		return "P"
	case handSum == 21 && dealerCardValue < 10:
		return "W"
	case handSum == 21 && dealerCardValue >= 10:
		return "S"
	case 17 <= handSum && handSum <= 20 || 12 <= handSum && handSum <= 16 && dealerCardValue < 7:
		return "S"
	case handSum <= 11 || 12 <= handSum && handSum <= 16 && dealerCardValue >= 7:
		return "H"
	default:
		return "S"
	}
}
