package primes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPrimes(t *testing.T) {
	Convey("Between should work", t, func() {
		max := uint64(10000)
		primes := Between(max, 0, EratosthenesAlgo)
		So(len(primes), ShouldEqual, 1229)
		for _, val := range primes {
			So(isPrime(val), ShouldBeTrue)
		}
	})
}
