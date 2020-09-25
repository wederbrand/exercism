// Package proverb has a package comment that summarizes what it's about.
package proverb

import (
	"fmt"
)

// Proverb has a comment documenting it.
func Proverb(rhyme []string) []string {
	var strings []string

	for i := range rhyme {
		if i == len(rhyme)-1 {
			strings = append(strings, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
		} else {
			strings = append(strings, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
		}

	}
	return strings
}
