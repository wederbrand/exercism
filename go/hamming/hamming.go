// Package hamming contains utils to play with DNA.
package hamming

import "errors"

// Distance calculates the Hamming distance between to DNA strands.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("strands must be the same length")
	}

	diffs := 0
	for i := range a {
		if a[i] != b[i] {
			diffs++
		}
	}

	return diffs, nil
}
