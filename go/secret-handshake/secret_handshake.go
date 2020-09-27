package secret

func Handshake(code uint) []string {
	// we build the result with the highest value first so it needs to be reversed.
	reverse := true
	var result []string

	if code >= 16 {
		reverse = false
		code -= 16
	}

	if code >= 8 {
		result = append(result, "jump")
		code -= 8
	}

	if code >= 4 {
		result = append(result, "close your eyes")
		code -= 4
	}

	if code >= 2 {
		result = append(result, "double blink")
		code -= 2
	}

	if code >= 1 {
		result = append(result, "wink")
		code -= 1
	}

	if reverse {
		for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
			result[i], result[j] = result[j], result[i]
		}
	}

	return result
}
