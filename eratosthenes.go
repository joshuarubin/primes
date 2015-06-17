package primes

import "math"

// Eratosthenes is a bitset that sets bits at each index to 1 if the index
// number itself is a prime
type Eratosthenes []byte

// NewEratosthenes returns a new Eratosthenes calculated for all values from 0
// to the nearest multiple of 8 greater than l
func NewEratosthenes(l uint64) Sieve {
	numBlocks := uint64(math.Floor(float64(l)/8)) + 1

	// initialize the sieve with all values turned on
	s := make(Eratosthenes, int(numBlocks))
	for i := uint64(0); i < numBlocks; i++ {
		s[i] = byte(math.MaxUint8)
	}

	// set 1 as not-prime, 2 and 3 as prime to seed the sieve
	s[0]--

	for val := uint64(2); val <= sqrt(numBlocks*8); val++ {
		if !s.IsPrime(val) {
			continue
		}

		// val is prime, so disable all multiples of it as prime
		for next := 2 * val; next < numBlocks*8; next += val {
			s.setNotPrime(next)
		}
	}

	return s
}

// Len returns the size of the Eratosthenes in bytes
func (s Eratosthenes) Len() uint64 {
	return uint64(len(s))
}

func (s Eratosthenes) setNotPrime(val uint64) {
	b, mask := s.byteFor(val)
	*b = *b & flipBits(mask)
}

// IsPrime returns if value is a prime
func (s Eratosthenes) IsPrime(i uint64) bool {
	b, mask := s.byteFor(i)
	return *b&mask > 0
}

func (s Eratosthenes) byteFor(i uint64) (b *byte, mask byte) {
	return &s[i/8], byte(1) << ((i % 8) - 1)
}
