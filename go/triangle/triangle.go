// Package triangle has a package comment that summarizes what it's about.
package triangle

import (
	"math"
)

// Kind is actually just an int
type Kind int

const (
	// NaT is Not A Triangle
	NaT Kind = iota
	// Equ is equilateral
	Equ Kind = iota
	// Iso is isosceles
	Iso Kind = iota
	// Sca is scalene
	Sca Kind = iota
)

// KindFromSides has a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return NaT
	}

	if math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1) {
		return NaT
	}

	if math.IsInf(a, -1) || math.IsInf(b, -1) || math.IsInf(c, -1) {
		return NaT
	}

	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}

	if a+b < c || b+c < a || c+a < b {
		return NaT
	}

	if a == b && b == c {
		return Equ
	}

	if a == b || b == c || c == a {
		return Iso
	}

	return Sca
}
