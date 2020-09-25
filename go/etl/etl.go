// Package etl has documentation
package etl

import (
	"strings"
)

// Transform would never be happy without some documentation
func Transform(given map[int][]string) map[string]int {
	var expectation = make(map[string]int)
	for value, letters := range given {
		for _, letter := range letters {
			expectation[strings.ToLower(letter)] = value
		}
	}

	return expectation
}
