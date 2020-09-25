// Package raindrops makes funny sounds out of rain.
package raindrops

import "strconv"

// Convert converts an integer input into amazing sounds.
func Convert(input int) (output string) {
	if input%3 == 0 {
		output += "Pling"
	}
	if input%5 == 0 {
		output += "Plang"
	}
	if input%7 == 0 {
		output += "Plong"
	}

	if len(output) == 0 {
		output = strconv.Itoa(input)
	}
	return
}
