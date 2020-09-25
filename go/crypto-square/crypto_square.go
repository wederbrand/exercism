package cryptosquare

import (
	"strings"
)

func Encode(input string) string {
	input = strings.ToLower(input)
	var strippedInput strings.Builder
	for _, r := range input {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			strippedInput.WriteRune(r)
		}
	}

	input = strippedInput.String()

	// calculate size of rectangle
	w := 0
	h := 0
	for w*h < len(input) {
		if w == h {
			w++
		} else {
			h++
		}
	}

	var sb strings.Builder
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			offset := i + j*w
			if offset < len(input) {
				sb.WriteByte(input[offset])
			} else {
				sb.WriteRune(' ')
			}
		}
		if i < w-1 {
			sb.WriteRune(' ')
		}
	}

	return sb.String()
}
