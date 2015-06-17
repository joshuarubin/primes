package primes

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"time"

	"github.com/dustin/go-humanize"
)

// ErrUnknownSieveAlgo is returned when an invalid sieve algorithm was selected
type ErrUnknownSieveAlgo SieveAlgo

func (err ErrUnknownSieveAlgo) Error() string {
	return "unknown sieve algorithm: " + SieveAlgo(err).String()
}

func getSieve(n uint64, algo SieveAlgo) Sieve {
	switch algo {
	case EratosthenesAlgo:
		return NewEratosthenes(n)
	case SundaramAlgo:
		return NewSundaram(n)
	case AtkinAlgo:
		return NewAtkin(n)
	}

	return nil
}

// Between returns a list of all primes between a and b, inclusive
func Between(a, b uint64, algo SieveAlgo, stats bool) ([]uint64, error) {
	start := time.Now()

	if b < a { // ensure a <= b
		a, b = b, a
	}

	s := getSieve(b, algo)
	if s == nil {
		return nil, ErrUnknownSieveAlgo(algo)
	}

	ret := make([]uint64, 0, b-a+1)
	for i := a; i <= b; i++ {
		if s.IsPrime(i) {
			ret = append(ret, i)
		}
	}

	printStats(stats, a, b, algo, start, len(ret), s.Len())

	return ret, nil
}

func printStats(stats bool, a, b uint64, algo SieveAlgo, start time.Time, l int, mem uint64) {
	if !stats {
		return
	}

	fmt.Fprintf(os.Stderr, "found %s primes between %s and %s (inclusive) in %v using %s of memory using %s\n",
		humanize.Comma(int64(l)),
		humanize.Comma(int64(a)),
		humanize.Comma(int64(b)),
		time.Now().Sub(start),
		humanize.Bytes(mem),
		algo)
}

// Write efficiently writes the primes between a and b to w, newline separated
func Write(w io.Writer, a, b uint64, algo SieveAlgo, stats bool) error {
	start := time.Now()

	if b < a { // ensure a <= b
		a, b = b, a
	}

	s := getSieve(b, algo)
	if s == nil {
		return ErrUnknownSieveAlgo(algo)
	}

	n := 0
	buf := bufio.NewWriter(w)
	for i := a; i <= b; i++ {
		if s.IsPrime(i) {
			n++
			buf.Write([]byte(fmt.Sprintf("%d\n", i)))
		}
	}

	printStats(stats, a, b, algo, start, n, s.Len())

	return buf.Flush()
}

// IsPrime returns whether val isa prime
func IsPrime(val uint64) bool {
	if val == 0 || val == 1 {
		return false
	}

	if val == 2 {
		return true
	}

	for i := uint64(2); i <= uint64(math.Sqrt(float64(val))); i++ {
		if val%i == 0 {
			return false
		}
	}

	return true
}
