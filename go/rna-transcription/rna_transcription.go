// Package strand is of course documented
package strand

// ToRNA is also properly documented
func ToRNA(dna string) string {
	var result string

	for _, c := range dna {
		switch c {
		case 'G':
			result += "C"
		case 'C':
			result += "G"
		case 'T':
			result += "A"
		case 'A':
			result += "U"
		}
	}

	return result
}
