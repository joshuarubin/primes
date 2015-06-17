package primes

import (
	"fmt"
	"os"
	"time"

	"github.com/dustin/go-humanize"
)

// Between returns a list of all primes between a and b, inclusive
func Between(a, b uint64, algo SieveAlgo) []uint64 {
	start := time.Now()

	if b < a {
		// ensure a <= b
		a, b = b, a
	}

	ret := []uint64{}
	var s Sieve

	switch algo {
	case EratosthenesAlgo:
		s = NewEratosthenes(b)
	default:
		fmt.Fprintf(os.Stderr, "unknown sieve algorithm: %v", algo)
		return ret
	}

	for i := a; i <= b; i++ {
		if s.IsPrime(i) {
			ret = append(ret, i)
		}
	}

	fmt.Fprintf(os.Stderr, "found %s primes between %s and %s (inclusive) in %v using %s of memory using %s\n",
		humanize.Comma(int64(len(ret))),
		humanize.Comma(int64(a)),
		humanize.Comma(int64(b)),
		time.Now().Sub(start),
		humanize.Bytes(s.Len()),
		algo)

	return ret
}
