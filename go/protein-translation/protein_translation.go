// Package protein has been documented, I don't dare otherwise
package protein

import (
	"errors"
)

// ErrStop stop processing the RNA strand
var ErrStop = errors.New("stop")

// ErrInvalidBase you did a boohoo
var ErrInvalidBase = errors.New("invalid base")

// FromCodon is from codon
func FromCodon(input string) (string, error) {
	for {
		switch input {
		case "AUG":
			return "Methionine", nil
		case "UUU", "UUC":
			return "Phenylalanine", nil
		case "UUA", "UUG":
			return "Leucine", nil
		case "UCU", "UCC", "UCA", "UCG":
			return "Serine", nil
		case "UAU", "UAC":
			return "Tyrosine", nil
		case "UGU", "UGC":
			return "Cysteine", nil
		case "UGG":
			return "Tryptophan", nil
		case "UAA", "UAG", "UGA":
			return "", ErrStop
		default:
			return "", ErrInvalidBase
		}
	}
}

// FromRNA is from RNA
func FromRNA(input string) ([]string, error) {
	var result []string

	for i := 0; i < len(input); i += 3 {
		codon := input[i : i+3]
		fromCodon, err := FromCodon(codon)
		if err != nil {
			if err == ErrInvalidBase {
				return result, err
			}
			if err == ErrStop {
				return result, nil
			}
		}

		result = append(result, fromCodon)
	}

	return result, nil
}
