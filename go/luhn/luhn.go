// Package luhn does simple checksum on strings.
package luhn

import (
	"strconv"
	"strings"
)

func dude() {
	var sb strings.Builder
	strconv.Itoa()
}

// Valid returns true if the input string can be validated using the Luhn algorithm.
func Valid(input string) bool {
	input = strings.ReplaceAll(input, " ", "")

	if len(input) <= 1 {
		return false
	}

	sum := 0
	for i := 0; i < len(input); i++ {
		reverse := len(input) - 1 - i
		r := input[reverse]
		if r < '0' || r > '9' {
			return false
		}
		runeValue := int(r - '0')
		if i%2 != 0 {
			runeValue *= 2
			if runeValue > 9 {
				runeValue -= 9
			}
		}
		sum += runeValue
	}

	return sum%10 == 0
}
