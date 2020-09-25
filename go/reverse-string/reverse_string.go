// Package reverse reverses things
package reverse

import (
	"strings"
)

// Reverse reverses a string (ie list of runes)
func Reverse(input string) string {
	var runes []rune
	for _, r := range input {
		runes = append(runes, r)
	}

	var sb strings.Builder
	for i := len(runes) - 1; i >= 0; i-- {
		sb.WriteRune(runes[i])
	}

	return sb.String()
}
