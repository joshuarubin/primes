package primes

import (
	"math"

	"github.com/joshuarubin/primes/bitset"
)

// Sundaram is a Sieve that is calculated using the sieve of sundaram algorithm
type Sundaram struct {
	bitset.Bitset
	pos uint64
}

// NewSundaram returns a new Sundaram calculated for all values from 0 to the
// nearest byte greater than n
func NewSundaram(n uint64) Sieve {
	s := bitset.New(n)

	imax := sqrt(s.Max())
	for i := uint64(1); i <= imax; i++ {
		jmax := (n - i) / (1 + 2*i)
		for j := i; j <= jmax; j++ {
			s.Set(i + j + 2*i*j)
		}
	}

	ret := Sundaram{bitset.New(n).Set(2).Set(3), 0}
	imax = (s.Max() - 1) / 2
	for i := uint64(2); i <= imax; i++ {
		if !s.IsSet(i) {
			ret.Set(2*i + 1)
		}
	}

	return &ret
}

// Next returns the next prime in the sieve
func (s *Sundaram) Next() uint64 {
	for ; s.pos <= s.Max(); s.pos++ {
		if s.IsSet(s.pos) {
			s.pos++
			return s.pos - 1
		}
	}

	return math.MaxUint64
}

// SkipTo advances the position to return primes >= pos
func (s Sundaram) SkipTo(pos uint64) {
	s.pos = pos
}
