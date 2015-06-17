package primes

// Eratosthenes is a bitset that sets bits at each index to 1 if the index
// number itself is a prime
type Eratosthenes struct {
	BitSet
}

// NewEratosthenes returns a new Eratosthenes calculated for all values from 0
// to the nearest multiple of 8 greater than l
func NewEratosthenes(l uint64) Sieve {
	// initialize the sieve with all bits set but unset 1 (as it is not-prime)
	s := Eratosthenes{NewBitSet(l).SetAll().Unset(1)}

	for i := uint64(2); i <= sqrt(s.Len()*8); i++ {
		if !s.IsSet(i) {
			continue
		}

		// i is prime, so disable all multiples of it as prime
		for next := 2 * i; next < s.Len()*8; next += i {
			s.Unset(next)
		}
	}

	return s
}

// IsPrime returns if value is a prime
func (s Eratosthenes) IsPrime(i uint64) bool {
	return s.IsSet(i)
}
