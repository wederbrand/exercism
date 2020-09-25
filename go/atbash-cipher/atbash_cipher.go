package atbash

import (
	"strings"
	"unicode"
)

func Atbash(input string) string {
	translation := make(map[rune]rune)
	for i := 'a'; i <= 'z'; i++ {
		translation[i] = 'a' + 'z' - i
	}

	input = strings.ToLower(input)
	var strippedInput strings.Builder
	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			strippedInput.WriteRune(r)
		}
	}

	input = strippedInput.String()

	var result strings.Builder
	for i, r := range input {
		if i%5 == 0 {
			result.WriteRune(' ')
		}
		if unicode.IsLetter(r) {
			result.WriteRune(translation[r])
		} else {
			result.WriteRune(r)
		}
	}
	return strings.Trim(result.String(), " ")
}
