package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    switch card {
        case "ace": return 11
    	case "two": return 2
    	case "three": return 3
    	case "four": return 4
    	case "five": return 5
    	case "six": return 6
    	case "seven": return 7
    	case "eight": return 8
    	case "nine": return 9
    	case "ten": return 10
    	case "jack": return 10
    	case "queen": return 10
    	case "king": return 10
    	default: return 0
    }
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    sumOfCards := ParseCard(card1) + ParseCard(card2)
    
    switch {
        case sumOfCards == 22: return "P"
		case sumOfCards == 21 && ParseCard(dealerCard) < 10: return "W"
    	case sumOfCards >= 17 && sumOfCards <= 21: return "S"
		case sumOfCards >= 12 && sumOfCards <= 17  && ParseCard(dealerCard) < 7: return "S"
    	default: return "H"
    }
}
