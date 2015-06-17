package primes

// Eratosthenes is a Sieve that is calculated using the sieve of eratosthenes
// algorithm
type Eratosthenes struct {
	BitSet
}

// NewEratosthenes returns a new Eratosthenes calculated for all values from 0
// to the nearest byte greater than n
func NewEratosthenes(n uint64) Sieve {
	// initialize the sieve with all bits set but unset 0 and 1 (as they are
	// not-prime)
	s := Eratosthenes{NewBitSet(n).SetAll().Unset(0).Unset(1)}

	for i := uint64(2); i <= sqrt(s.Max()); i++ {
		if !s.IsPrime(i) {
			continue
		}

		// i is prime, so disable all multiples of it as prime
		for next := 2 * i; next < s.Len()*Byte; next += i {
			s.Unset(next)
		}
	}

	return s
}

// IsPrime returns if value is a prime
func (s Eratosthenes) IsPrime(i uint64) bool {
	return s.IsSet(i)
}

// ListPrimes returns the set of all primes in the sieve
func (s Eratosthenes) ListPrimes() []uint64 {
	return s.ListSet()
}
