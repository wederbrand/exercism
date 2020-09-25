// Package anagram has fun with anagrams.
package anagram

import (
	"sort"
	"strings"
)

// isAnagram returns true if i all the runes in a exists in b.
func isAnagram(a, b string) bool {
	a = strings.ToLower(a)
	b = strings.ToLower(b)
	if a == b {
		return false
	}

	as := []rune(a)
	bs := []rune(b)

	sort.Slice(as, func(i, j int) bool { return as[i] < as[j] })
	sort.Slice(bs, func(i, j int) bool { return bs[i] < bs[j] })

	return string(as) == string(bs)
}

// Detect returns all anagrams amongst a list of candidates.
func Detect(subject string, candidates []string) (result []string) {
	for _, candidate := range candidates {
		if isAnagram(subject, candidate) {
			result = append(result, candidate)
		}
	}

	return
}
