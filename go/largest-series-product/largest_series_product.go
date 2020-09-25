package lsproduct

import (
	"errors"
)

func LargestSeriesProduct(digits string, span int) (int, error) {
	if span == 0 {
		return 1, nil
	}

	if span > len(digits) {
		return 0, errors.New("span must be smaller than string length")
	}

	if span < 0 {
		return 0, errors.New("span must be greater than zero")
	}

	maxProduct := -1
	for start := 0; start <= len(digits)-span; start++ {
		product := 1
		for i := start; i < span+start; i++ {
			positionValue := int(digits[i] - '0')
			if positionValue < 0 || positionValue > 9 {
				return 0, errors.New("digits input must only contain digits")
			}
			product *= positionValue
		}
		if product > maxProduct {
			maxProduct = product
		}
	}

	return maxProduct, nil
}
