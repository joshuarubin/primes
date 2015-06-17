package primes

// SieveAlgo is an enum representing different kinds of sieve algorithms
type SieveAlgo int

const (
	// EratosthenesAlgo is the sieve of eratosthene algorithm
	EratosthenesAlgo SieveAlgo = iota

	// SundaramAlgo is the sieve of sundaram algorithm
	SundaramAlgo

	// AtkinAlgo is the sieve of atkin algorithm
	AtkinAlgo
)

var (
	SieveAlgos []SieveAlgo
)

func init() {
	SieveAlgos = make([]SieveAlgo, 3)

	SieveAlgos[EratosthenesAlgo] = EratosthenesAlgo
	SieveAlgos[SundaramAlgo] = SundaramAlgo
	SieveAlgos[AtkinAlgo] = AtkinAlgo
}

//go:generate stringer -type=SieveAlgo
