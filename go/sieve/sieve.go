package sieve

func Sieve(n int) []int {
	var primes []int
	var sieve [1001]bool // keep track of all found primes. Next not marked is the next prime.

	for i := 2; i <= n; i++ {
		// if prime
		if !sieve[i] {
			// add it
			primes = append(primes, i)
			// and mark it
			for j := i; j <= n; j += i {
				sieve[j] = true
			}
		}
	}

	return primes
}
