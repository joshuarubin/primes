package primes

import (
	"bytes"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPrimes(t *testing.T) {
	max := uint64(10000)

	Convey("Primes should be calculated correctly with different algorithms", t, func() {
		Convey("IsPrime should be correct", func() {
			So(IsPrime(0), ShouldBeFalse)
			So(IsPrime(1), ShouldBeFalse)
			So(IsPrime(2), ShouldBeTrue)
			So(IsPrime(3), ShouldBeTrue)
			So(IsPrime(4), ShouldBeFalse)
		})

		Convey("for invalid algorithms", func() {
			primes, err := Between(0, max, SieveAlgo(999), false)
			So(primes, ShouldResemble, []uint64(nil))
			So(err, ShouldResemble, ErrUnknownSieveAlgo(999))
			So(err.Error(), ShouldEqual, "unknown sieve algorithm: SieveAlgo(999)")

			var b bytes.Buffer
			err = Write(&b, 0, max, SieveAlgo(999), false)
			So(err, ShouldResemble, ErrUnknownSieveAlgo(999))
			So(err.Error(), ShouldEqual, "unknown sieve algorithm: SieveAlgo(999)")
			So(string(b.Bytes()), ShouldEqual, "")
		})

		Convey("Between should work", func() {
			for i, algo := range SieveAlgos {
				primes, err := Between(max, 0, algo, i%2 == 0)
				So(err, ShouldBeNil)
				So(len(primes), ShouldEqual, 1229)
				for _, val := range primes {
					So(IsPrime(val), ShouldBeTrue)
				}
			}
		})

		Convey("Write shuold work", func() {
			for i, algo := range SieveAlgos {
				var b bytes.Buffer
				err := Write(&b, 10, 0, algo, i%2 == 0)
				So(err, ShouldBeNil)
				So(string(b.Bytes()), ShouldEqual, "2\n3\n5\n7\n")
			}
		})
	})
}
