package primes

import (
	"math"

	"github.com/joshuarubin/primes/bitset"
)

// Eratosthenes is a Sieve that is calculated using the sieve of eratosthenes
// algorithm
type Eratosthenes struct {
	bitset.Bitset
	pos uint64
}

// NewEratosthenes returns a new Eratosthenes calculated for all values from 0
// to the nearest byte greater than n
func NewEratosthenes(n uint64) Sieve {
	// initialize the sieve with all bits set but unset 0 and 1 (as they are
	// not-prime)
	s := Eratosthenes{bitset.New(n).SetAll().Unset(0).Unset(1), 0}

	for i := uint64(2); i <= sqrt(s.Max()); i++ {
		if !s.IsSet(i) {
			continue
		}

		// i is prime, so disable all multiples of it as prime
		for next := 2 * i; next < s.Len()*bitset.Byte; next += i {
			s.Unset(next)
		}
	}

	return &s
}

// Next returns the next prime in the sieve
func (s *Eratosthenes) Next() uint64 {
	for ; s.pos <= s.Max(); s.pos++ {
		if s.IsSet(s.pos) {
			s.pos++
			return s.pos - 1
		}
	}

	return math.MaxUint64
}

// SkipTo advances the position to return primes >= pos
func (s Eratosthenes) SkipTo(pos uint64) {
	s.pos = pos
}
