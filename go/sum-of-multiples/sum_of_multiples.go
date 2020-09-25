// Package summultiples make freakish calculations for fun.
package summultiples

// SumMultiples returns pretty large numbers.
func SumMultiples(limit int, divisors ...int) (result int) {
	for i := 0; i < limit; i++ {
		for _, divisor := range divisors {
			if divisor == 0 {
				continue
			}
			if i%divisor == 0 {
				result += i
				break // if it's a match we don't need to check any more divisors
			}
		}
	}

	return
}
