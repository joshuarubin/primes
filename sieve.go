package primes

import "math"

// Sieve is a simple interface for different sieve algorithms
type Sieve interface {
	Next() uint64
	SkipTo(uint64)
}

func sqrt(val uint64) uint64 {
	return uint64(math.Sqrt(float64(val)))
}
