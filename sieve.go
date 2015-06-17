package primes

import "math"

// Sieve is a simple interface for different sieve algorithms
type Sieve interface {
	IsPrime(uint64) bool
	Len() uint64
	ListPrimes() []uint64
}

func sqrt(val uint64) uint64 {
	return uint64(math.Sqrt(float64(val)))
}
