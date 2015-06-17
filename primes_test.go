package primes

import (
	"math"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPrimes(t *testing.T) {
	max := uint64(10000)

	Convey("Between should work", t, func() {
		Convey("for invalid algorithms", func() {
			primes := Between(0, max, SieveAlgo(999))
			So(primes, ShouldResemble, []uint64(nil))
		})

	Loop:
		for _, algo := range SieveAlgos {
			switch algo {
			case AtkinAlgo:
				Print("TODO(jrubin) enable prime test for ", algo, "\n")
				continue Loop
			}

			primes := Between(max, 0, algo)
			So(len(primes), ShouldEqual, 1229)
			for _, val := range primes {
				So(isPrime(val), ShouldBeTrue)
			}
		}
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
