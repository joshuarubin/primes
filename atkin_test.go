package primes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAtkin(t *testing.T) {
	Convey("Atkin algorithm should work", t, func() {
		max := uint64(10000)
		s := NewAtkin(max)
		So(s.Len(), ShouldEqual, 1251)
		numPrimes := 0
		for i := uint64(0); i <= max; i++ {
			if s.IsPrime(i) {
				numPrimes++
				So(IsPrime(i), ShouldBeTrue)
			} else {
				So(IsPrime(i), ShouldBeFalse)
			}
		}
		So(numPrimes, ShouldEqual, 1229)
		ps := s.ListPrimes()
		So(len(ps), ShouldEqual, 1230)
		for _, p := range ps {
			So(IsPrime(p), ShouldBeTrue)
		}
	})
}
