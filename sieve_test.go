package primes

import (
	"math"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSieve(t *testing.T) {
	Convey("Sieve should work", t, func() {
		max := uint64(10000)
		s := NewSieve(max)
		So(s.Len(), ShouldEqual, 1251)
		numPrimes := 0
		for i := uint64(0); i <= max; i++ {
			if s.IsPrime(i) {
				numPrimes++
				So(isPrime(i), ShouldBeTrue)
			} else {
				So(isPrime(i), ShouldBeFalse)
			}
		}
		So(numPrimes, ShouldEqual, 1229)
	})
}

// a very naïve approach to testing for primes
func isPrime(val uint64) bool {
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
