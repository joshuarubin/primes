package primes

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/jbarham/primegen.go"
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
		return primegen.New()
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
	s.SkipTo(a)
	for i := s.Next(); i <= b; i = s.Next() {
		ret = append(ret, i)
	}

	printStats(stats, a, b, algo, start, len(ret))

	return ret, nil
}

func printStats(stats bool, a, b uint64, algo SieveAlgo, start time.Time, l int) {
	if !stats {
		return
	}

	fmt.Fprintf(os.Stderr, "found %s primes between %s and %s (inclusive) in %v using %s\n",
		humanize.Comma(int64(l)),
		humanize.Comma(int64(a)),
		humanize.Comma(int64(b)),
		time.Now().Sub(start),
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
	s.SkipTo(a)
	for i := s.Next(); i <= b; i = s.Next() {
		n++
		buf.Write([]byte(fmt.Sprintf("%d\n", i)))
	}

	printStats(stats, a, b, algo, start, n)

	return buf.Flush()
}

// IsPrime returns whether val isa prime
func IsPrime(val uint64) bool {
	if val < 2 {
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
