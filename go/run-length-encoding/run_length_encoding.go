// Package encode encodes stuff
package encode

import (
	"strconv"
	"unicode"
)

// RunLengthEncode encodes the input.
func RunLengthEncode(input string) (result string) {
	lastChar := '_'
	currentCount := 0
	for _, c := range input {
		if c == lastChar {
			currentCount++
		} else {
			printLetter(currentCount, &result, lastChar)

			lastChar = c
			currentCount = 1
		}
	}

	printLetter(currentCount, &result, lastChar)

	return
}

func printLetter(currentCount int, result *string, lastChar int32) {
	if currentCount == 1 {
		*result += string(lastChar)
	}

	if currentCount > 1 {
		*result += strconv.Itoa(currentCount)
		*result += string(lastChar)
	}
}

// RunLengthDecode decodes the input.
func RunLengthDecode(input string) (result string) {
	currentCount := 0
	for _, c := range input {
		if !unicode.IsDigit(c) {
			if currentCount == 0 {
				// special case for letters that are not encoded
				currentCount = 1
			}

			for i := 0; i < currentCount; i++ {
				result += string(c)
			}

			currentCount = 0
		} else {
			atoi, _ := strconv.Atoi(string(c))
			currentCount = currentCount*10 + atoi
		}
	}

	return
}
