package primes

import "github.com/joshuarubin/primes/bitset"

// Sundaram is a Sieve that is calculated using the sieve of sundaram algorithm
type Sundaram struct {
	bitset.Bitset
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

	ret := Sundaram{bitset.New(n).Set(2).Set(3)}
	imax = (s.Max() - 1) / 2
	for i := uint64(2); i <= imax; i++ {
		if !s.IsSet(i) {
			ret.Set(2*i + 1)
		}
	}

	return ret
}

// IsPrime returns if value is a prime
func (s Sundaram) IsPrime(i uint64) bool {
	return s.IsSet(i)
}

// ListPrimes returns the set of all primes in the sieve
func (s Sundaram) ListPrimes() []uint64 {
	return s.ListSet()
}

// Len returns the size of the BitSet in bytes
func (s Sundaram) Len() uint64 {
	// twice the bitset size since 2 were required
	return s.Bitset.Len() * 2
}
