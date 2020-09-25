// Package isogram is well documented
package isogram

import (
	"unicode"
)

// IsIsogram might be even better documented
func IsIsogram(input string) bool {
	seen := make(map[rune]bool)

	for _, r := range input {
		u := unicode.ToUpper(r)

		if u == ' ' || u == '-' {
			continue
		}

		if seen[u] {
			return false
		}

		seen[u] = true
	}

	return true
}
