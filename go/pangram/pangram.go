// Package pangram would hardly work without docs
package pangram

import (
	"strings"
)

// IsPangram should apparently be documented
func IsPangram(input string) bool {
	input = strings.ToLower(input)
	for _, r := range "abcdefghijklmnopqrstuvwxyz" {
		if !strings.ContainsRune(input, r) {
			return false
		}
	}
	return true
}
