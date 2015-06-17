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
		So(s.IsSet(6), ShouldBeFalse)

		s.SetAll()
		So(s.IsSet(6), ShouldBeTrue)
		So(s, ShouldResemble, BitSet{math.MaxUint8, math.MaxUint8})

		So(s.IsSet(11), ShouldBeTrue)
		s.Unset(11)
		So(s.IsSet(11), ShouldBeFalse)
		So(s, ShouldResemble, BitSet{math.MaxUint8, math.MaxUint8 - (11 % 8) - 1})

		So(s.IsSet(3), ShouldBeTrue)
		s.Unset(3)
		So(s.IsSet(3), ShouldBeFalse)
		So(s, ShouldResemble, BitSet{math.MaxUint8 - (3 % 8) - 1, math.MaxUint8 - (11 % 8) - 1})

		s.Set(3)
		s.Set(11)
		So(s, ShouldResemble, BitSet{math.MaxUint8, math.MaxUint8})

		// test for out of bounds
		So(func() { s.Set(16) }, ShouldPanic)
		So(func() { s.IsSet(16) }, ShouldPanic)
		So(func() { s.Unset(16) }, ShouldPanic)
	})
}
