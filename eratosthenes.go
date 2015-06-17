package primes

// Eratosthenes is a bitset that sets bits at each index to 1 if the index
// number itself is a prime
type Eratosthenes struct {
	BitSet
}

// NewEratosthenes returns a new Eratosthenes calculated for all values from 0
// to the nearest multiple of 8 greater than l
func NewEratosthenes(l uint64) Sieve {
	s := Eratosthenes{
		// initialize the sieve with all values turned on
		BitSet: NewBitSet(l).SetAll(),
	}

	// set 1 as not-prime
	s.Unset(1)

	for val := uint64(2); val <= sqrt(s.Len()*8); val++ {
		if !s.IsSet(val) {
			continue
		}

		// val is prime, so disable all multiples of it as prime
		for next := 2 * val; next < s.Len()*8; next += val {
			s.Unset(next)
		}
	}

	return s
}

// IsPrime returns if value is a prime
func (s Eratosthenes) IsPrime(i uint64) bool {
	return s.IsSet(i)
}
