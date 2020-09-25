package diffiehellman

import (
	"math/big"
	"math/rand"
	"time"
)

func PrivateKey(p *big.Int) *big.Int {
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	// rand will return 0 up to, but not including our limit.
	// We need to go higher than 1 so we'll add 2 later.
	// Hence, removing it now to keep the limit low enough
	lim := new(big.Int).Sub(p, big.NewInt(2))
	result := new(big.Int).Rand(seed, lim)
	return result.Add(result, big.NewInt(2))
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	// A = g**a mod p
	return new(big.Int).Exp(big.NewInt(g), private, p)
}

func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey(p)
	public = PublicKey(private, p, g)
	return
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	// s = B**a mod p
	return new(big.Int).Exp(public2, private1, p)
}
