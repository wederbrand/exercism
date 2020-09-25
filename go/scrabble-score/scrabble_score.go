// Package scrabble calculates stuff related to Scrabble.
package scrabble

import (
	"unicode"
)

// Score calculates the score for any word in Scrabble.
func Score(input string) (result int) {
	for _, r := range input {
		switch unicode.ToUpper(r) {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			result++
		case 'D', 'G':
			result += 2
		case 'B', 'C', 'M', 'P':
			result += 3
		case 'F', 'H', 'V', 'W', 'Y':
			result += 4
		case 'K':
			result += 5
		case 'J', 'X':
			result += 8
		case 'Q', 'Z':
			result += 10
		}
	}
	return
}
