package cipher

import (
	"math"
	"regexp"
	"strings"
)

type Vigenere string

func NewCaesar() Cipher {
	return NewVigenere("d")
}

func NewShift(distance int) Cipher {
	// 1 to 25 or -1 to -25
	if math.Abs(float64(distance)) < 1 || math.Abs(float64(distance)) > 25 {
		return nil
	}
	return NewVigenere(string(itor(distance)))
}

func NewVigenere(key string) Cipher {
	onlyLowerCase := regexp.MustCompile("^[a-z]+$")
	if !onlyLowerCase.MatchString(key) {
		return nil
	}

	notOnlyTheSmallA := regexp.MustCompile("[b-z]+")
	if !notOnlyTheSmallA.MatchString(key) {
		return nil
	}

	vigenere := Vigenere(strings.ToLower(key))
	return &vigenere
}

// returns 0 for 'a' and 25 for 'z'
func rtoi(r rune) int {
	return int(r) - int('a')
}

// returns 'a' for 1 and 'z' for 25
// also aligns the input between 0 and 25
func itor(i int) (r rune) {
	i %= 26
	if i < 0 {
		i += 26
	}
	return rune(int('a') + i)
}

func (c *Vigenere) Encode(s string) string {
	s = strings.ToLower(s)
	reg := regexp.MustCompile("[^a-zA-Z]")
	s = reg.ReplaceAllString(s, "")
	var sb strings.Builder
	for i, r := range s {
		phraseOffset := i % len(*c)
		vigenere := *c
		r2 := rune(vigenere[phraseOffset])
		offset := rtoi(r2)
		r = itor(rtoi(r) + offset)
		sb.WriteRune(r)
	}

	return sb.String()
}

func (c *Vigenere) Decode(s string) string {
	s = strings.ToLower(s)
	reg := regexp.MustCompile("[^a-zA-Z]")
	s = reg.ReplaceAllString(s, "")
	var sb strings.Builder
	for i, r := range s {
		phraseOffset := i % len(*c)
		vigenere := *c
		r2 := rune(vigenere[phraseOffset])
		offset := rtoi(r2)
		r = itor(rtoi(r) - offset)
		sb.WriteRune(r)
	}

	return sb.String()
}
