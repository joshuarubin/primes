package primes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSieveAlgo(t *testing.T) {
	Convey("SieveAlgos should be correct", t, func() {
		So(EratosthenesAlgo, ShouldEqual, 0)
		So(EratosthenesAlgo.String(), ShouldEqual, "EratosthenesAlgo")
	})
}
