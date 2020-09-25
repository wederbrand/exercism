// Package pythagorean cleverly calculates things with triangles.
package pythagorean

import (
	"math"
)

// Triplet is really just the sides of a triangle.
type Triplet [3]int

// Range returns a list of all Pythagorean triplets with sides in the
// range min to max inclusive.
func Range(min, max int) (result []Triplet) {
	for a := min; a <= max-2; a++ {
		for b := a + 1; b <= max-1; b++ {
			c := math.Sqrt(math.Pow(float64(a), 2) + math.Pow(float64(b), 2))
			intPart, decimalPart := math.Modf(c)
			if int(intPart) > b && int(intPart) <= max && decimalPart == 0 {
				result = append(result, Triplet{a, b, int(intPart)})
			}
		}
	}

	return
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c
// (the perimeter) is equal to p.
func Sum(p int) (result []Triplet) {
	for _, triplet := range Range(1, p/2) {
		if triplet[0]+triplet[1]+triplet[2] == p {
			result = append(result, triplet)
		}
	}

	return
}
