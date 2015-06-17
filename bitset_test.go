package primes

import (
	"math"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBitSet(t *testing.T) {
	Convey("BitSet should work", t, func() {
		s := NewBitSet(9)
		So(s.Len(), ShouldEqual, 2)
		So(s, ShouldResemble, BitSet{0, 0})
		So(s.IsSet(0), ShouldBeFalse)
		So(s.IsSet(6), ShouldBeFalse)
		So(s.Max(), ShouldEqual, 15)

		Convey("basic functionality", func() {
			s.SetAll()
			So(s.IsSet(0), ShouldBeTrue)
			So(s.IsSet(6), ShouldBeTrue)
			So(s, ShouldResemble, BitSet{math.MaxUint8, math.MaxUint8})

			So(s.IsSet(11), ShouldBeTrue)
			s.Unset(11)
			So(s.IsSet(11), ShouldBeFalse)
			So(s, ShouldResemble, BitSet{math.MaxUint8, math.MaxUint8 - uint8(math.Pow(float64(2), float64(11%Byte)))})

			So(s.IsSet(3), ShouldBeTrue)
			s.Unset(3)
			So(s.IsSet(3), ShouldBeFalse)
			So(s, ShouldResemble, BitSet{
				math.MaxUint8 - uint8(math.Pow(float64(2), float64(3%Byte))),
				math.MaxUint8 - uint8(math.Pow(float64(2), float64(11%Byte))),
			})

			So(len(s.ListSet()), ShouldEqual, 14)
			So(s.ListSet(), ShouldResemble, []uint64{
				0, 1, 2, 4, 5, 6, 7, 8, 9, 10, 12, 13, 14, 15,
			})

			So(s.ListUnset(), ShouldResemble, []uint64{3, 11})

			s.Set(3)
			s.Set(11)
			So(s, ShouldResemble, BitSet{math.MaxUint8, math.MaxUint8})

			// test for out of bounds
			So(func() { s.Set(16) }, ShouldPanic)
			So(func() { s.IsSet(16) }, ShouldPanic)
			So(func() { s.Unset(16) }, ShouldPanic)
		})

		Convey("flipping bites", func() {
			So(s.IsSet(0), ShouldBeFalse)
			s.Flip(0)
			So(s.IsSet(0), ShouldBeTrue)
			So(s, ShouldResemble, BitSet{1, 0})
			s.Flip(0)
			So(s.IsSet(0), ShouldBeFalse)
			So(s, ShouldResemble, BitSet{0, 0})

			So(s.IsSet(11), ShouldBeFalse)
			s.Flip(11)
			So(s.IsSet(11), ShouldBeTrue)
			So(s, ShouldResemble, BitSet{0, 8})
			s.Flip(11)
			So(s.IsSet(11), ShouldBeFalse)
			So(s, ShouldResemble, BitSet{0, 0})
		})
	})
}
