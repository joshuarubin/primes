package primes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPrimes(t *testing.T) {
	max := uint64(10000)

	Convey("Between should work", t, func() {
		Convey("for invalid algorithms", func() {
			primes, err := Between(0, max, SieveAlgo(999), false)
			So(primes, ShouldResemble, []uint64(nil))
			So(err, ShouldResemble, ErrUnknownSieveAlgo(999))
			So(err.Error(), ShouldEqual, "unknown sieve algorithm: SieveAlgo(999)")
		})

	Loop:
		for _, algo := range SieveAlgos {
			switch algo {
			case AtkinAlgo:
				Print("TODO(jrubin) enable prime test for ", algo, "\n")
				continue Loop
			}

			primes, err := Between(max, 0, algo, false)
			So(err, ShouldBeNil)
			So(len(primes), ShouldEqual, 1229)
			for _, val := range primes {
				So(IsPrime(val), ShouldBeTrue)
			}
		}
	})
}
