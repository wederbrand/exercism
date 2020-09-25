package isbn

import (
	"strings"
)

func IsValidISBN(input string) bool {
	// get rid of dashes
	input = strings.Replace(input, "-", "", -1)
	if len(input) != 10 {
		return false
	}

	sum := 0
	for i, r := range input {
		multiplier := 10 - i
		value := int(r - '0')
		if r == 'X' {
			if i < 9 {
				// X is only allowed last
				return false
			}
			value = 10
		}
		if value < 0 || value > 10 {
			return false
		}

		// all good
		sum += value * multiplier
	}

	return sum%11 == 0
}
