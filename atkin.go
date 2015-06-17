package primes

import (
	"github.com/joshuarubin/primes/atkin"
	"github.com/joshuarubin/primes/bitset"
)

// Atkin is a Sieve that is calculated using the sieve of atkin algorithm
type Atkin struct {
	bitset.Bitset
	pg    *atkin.Primegen
	limit uint64
}

// NewAtkin returns a new Atkin calculated for all values from 0 to the nearest
// byte greater than limit
func NewAtkin(limit uint64) Sieve {
	ret := Atkin{
		Bitset: bitset.New(limit),
		pg:     atkin.New(),
		limit:  limit,
	}

	if limit < 2 {
		return ret
	}

	for p := uint64(2); p <= ret.Max(); p = ret.pg.Next() {
		ret.Set(p)
	}

	return ret
}

// IsPrime returns if value is a prime
func (s Atkin) IsPrime(i uint64) bool {
	return s.IsSet(i)
}

// ListPrimes returns the set of all primes in the sieve
func (s Atkin) ListPrimes() []uint64 {
	return s.ListSet()
}
