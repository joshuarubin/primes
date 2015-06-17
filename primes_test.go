package primes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPrimes(t *testing.T) {
	algos := []SieveAlgo{
		EratosthenesAlgo,
	}

	max := uint64(10000)

	Convey("Between should work", t, func() {
		Convey("for invalid algorithms", func() {
			primes := Between(0, max, SieveAlgo(999))
			So(primes, ShouldResemble, []uint64{})
		})

		for _, algo := range algos {
			primes := Between(max, 0, algo)
			So(len(primes), ShouldEqual, 1229)
			for _, val := range primes {
				So(isPrime(val), ShouldBeTrue)
			}
		}
	})
}
