// Package wordcount counts words.
package wordcount

import (
	"regexp"
	"strings"
)

// Frequency maps words to frequency.
type Frequency map[string]int

// WordCount counts the words in the input and returns a Frequency.
func WordCount(input string) Frequency {
	// replace illegal characters with space. All but alphanumeric and ' are illegal
	illegalCharacter := regexp.MustCompile(`[^\w']`)
	input = illegalCharacter.ReplaceAllString(input, " ")

	// replace quotes after word boundary
	afterSpaceGroup := regexp.MustCompile(`[\s]+'`)
	input = afterSpaceGroup.ReplaceAllString(input, " ")

	// replace quotes before word boundary
	beforeSpaceGroup := regexp.MustCompile(`'[\s]+`)
	input = beforeSpaceGroup.ReplaceAllString(input, " ")

	// replace quotes after word boundary
	endOfText := regexp.MustCompile(`'\z`)
	input = endOfText.ReplaceAllString(input, " ")

	// replace quotes before word boundary
	beginningOfText := regexp.MustCompile(`'\A`)
	input = beginningOfText.ReplaceAllString(input, " ")

	// split the rest into fields
	input = strings.ToLower(input)
	fields := strings.Fields(input)
	var results = make(map[string]int)
	for _, field := range fields {
		results[field]++
	}

	return results
}
