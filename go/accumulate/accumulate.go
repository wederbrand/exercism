// Package accumulate has been properly documented
package accumulate

// Accumulate has documentation
func Accumulate(given []string, converter func(string) string) []string {
	result := make([]string, len(given))

	for i, word := range given {
		result[i] = converter(word)
	}

	return result
}
