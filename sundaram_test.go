package primes

import (
	"math"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSundaram(t *testing.T) {
	Convey("Sundaram algorithm should work", t, func() {
		max := uint64(10000)
		s := NewSundaram(max)
		numPrimes := 0
		for i := s.Next(); i <= max; i = s.Next() {
			numPrimes++
			So(IsPrime(i), ShouldBeTrue)
		}
		So(numPrimes, ShouldEqual, 1229)
		So(s.Next(), ShouldEqual, uint64(math.MaxUint64))
	})
}
