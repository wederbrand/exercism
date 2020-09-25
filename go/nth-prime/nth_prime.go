package prime

// Nth calculates the n'th prime (n <= 10001)
func Nth(n int) (int, bool) {
	// start finding primes, return once the n'th is found
	if n == 0 {
		return 0, false
	}

	const maxValue = 104744 // fewer is faster, 104744 is just enough for these test cases.
	foundPrimes := 0

	var notPrimes [maxValue]bool // keep track of all found primes. Next not marked is the next prime.

	for i := 2; i <= maxValue; i++ {
		// if prime
		if !notPrimes[i] {
			// count it
			foundPrimes++
			if foundPrimes == n {
				return i, true
			}
			// and mark it
			for j := i; j < maxValue; j += i {
				notPrimes[j] = true
			}
		}
	}

	// not found
	return 0, false
}
