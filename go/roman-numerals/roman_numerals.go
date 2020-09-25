// Package romannumerals is documented
package romannumerals

import (
	"errors"
)

// ToRomanNumeral even functions are documented
func ToRomanNumeral(input int) (string, error) {
	var t = map[int]string{
		1000: "M",
		500:  "D",
		100:  "C",
		50:   "L",
		10:   "X",
		5:    "V",
		1:    "I",
	}

	if input <= 0 || input > 3000 {
		return "", errors.New("only values between 0 and 3000 are supported")
	}

	result := ""

	thousands := input / 1000
	result += handleOne(thousands, t[1000], "", "")

	hundreds := (input % 1000) / 100
	result += handleOne(hundreds, t[100], t[500], t[1000])

	tens := (input % 100) / 10
	result += handleOne(tens, t[10], t[50], t[100])

	ones := input % 10
	result += handleOne(ones, t[1], t[5], t[10])

	return result, nil
}

func handleOne(number int, small string, half string, full string) (result string) {
	if number == 4 {
		result += small + half
		number -= 4
	}

	if number == 9 {
		result += small + full
		number -= 9
	}

	if number >= 5 {
		result += half
		number -= 5
	}

	for i := 0; i < number; i++ {
		result += small
	}

	return
}
