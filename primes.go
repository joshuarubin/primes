package primes

import (
	"fmt"
	"os"
	"time"

	"github.com/dustin/go-humanize"
)

// Between returns a list of all primes between a and b, inclusive
func Between(a, b uint64) []uint64 {
	start := time.Now()

	ret := []uint64{}

	if b < a {
		// ensure a <= b
		a, b = b, a
	}

	s := NewSieve(b)

	for i := a; i <= b; i++ {
		if s.IsPrime(i) {
			ret = append(ret, i)
		}
	}

	fmt.Fprintf(os.Stderr, "found %s primes between %d and %d (inclusive) in %v using %s of memory\n",
		humanize.Comma(int64(len(ret))),
		a, b,
		time.Now().Sub(start),
		humanize.Bytes(s.Len()))

	return ret
}
