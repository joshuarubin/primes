package primes

// SieveAlgo is an enum representing different kinds of sieve algorithms
type SieveAlgo int

const (
	// EratosthenesAlgo is the sieve of eratosthene algorithm
	EratosthenesAlgo SieveAlgo = iota
)

//go:generate stringer -type=SieveAlgo
