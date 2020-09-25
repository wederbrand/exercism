package dna

import (
	"errors"
	"unicode"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

// Counts is a method on the DNA type. A method is a function with a special receiver argument.
func (d DNA) Counts() (Histogram, error) {
	var h = Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}

	for _, r := range d {
		u := unicode.ToUpper(r)
		switch u {
		case 'A', 'C', 'G', 'T':
			h[u]++
		default:
			return nil, errors.New("illegal character in strand")
		}
	}

	return h, nil
}
