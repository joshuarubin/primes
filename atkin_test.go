package primes

import (
	"testing"

	"github.com/jbarham/primegen.go"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAtkin(t *testing.T) {
	Convey("Atkin algorithm should work", t, func() {
		max := uint64(10000)
		s := primegen.New()
		numPrimes := 0
		for i := s.Next(); i <= max; i = s.Next() {
			numPrimes++
			So(IsPrime(i), ShouldBeTrue)
		}
		So(numPrimes, ShouldEqual, 1229)
	})
}
