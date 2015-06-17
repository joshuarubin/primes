package primes

import "math"

// SieveAlgo is an enum representing different kinds of sieve algorithms
type SieveAlgo int

const (
	// EratosthenesAlgo is the sieve of eratosthene algorithm
	EratosthenesAlgo SieveAlgo = iota
)

// Sieve is a simple interface for different sieve algorithms
type Sieve interface {
	IsPrime(uint64) bool
	Len() uint64
}

func sqrt(val uint64) uint64 {
	return uint64(math.Sqrt(float64(val)))
}

func flipBits(b byte) byte {
	return math.MaxUint8 &^ b
}
