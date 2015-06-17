package primes

import (
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
			case SundaramAlgo, AtkinAlgo:
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
