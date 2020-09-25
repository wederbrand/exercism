// Package grains have the prince and the servant throw rice at each other. For fun.
package grains

import (
	"errors"
	"math"
)

// Square returns the number of grains on one specific square.
func Square(input int) (uint64, error) {
	if input < 1 || input > 64 {
		return 0, errors.New("only inputs between 1 and 64 are allowed")
	}
	return uint64(1 << (input - 1)), nil
}

// Total returns the total number of grains on all the 64 squares.
func Total() uint64 {
	// for performance we can calculate the number of grains on the *next* square and subtract one, like this
	// square  total   performance
	//      1      1        =  2-1
	//      2      3        =  4-1
	//      4      7        =  8-1
	//      8     15        = 16-1

	return math.MaxUint64 // 1<<(65-1)) - 1
}
