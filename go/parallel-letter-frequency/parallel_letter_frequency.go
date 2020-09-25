// Package letter is all fun and games with letters.
package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency calculates frequency, concurrently.
func ConcurrentFrequency(texts []string) FreqMap {
	result := make(FreqMap, len(texts))
	resultChannel := make(chan FreqMap)

	// send the results of each Frequency to the channel
	for _, text := range texts {
		go func(s string) {
			resultChannel <- Frequency(s)
		}(text)
	}

	// wait for exactly the same number of messages as we have texts
	for range texts {
		for letter, number := range <-resultChannel {
			result[letter] += number
		}
	}

	return result
}
