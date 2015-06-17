package primes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSieveAlgo(t *testing.T) {
	Convey("SieveAlgos should be correct", t, func() {
		So(len(SieveAlgos), ShouldEqual, 3)
		for i, algo := range SieveAlgos {
			So(SieveAlgos[i], ShouldEqual, i)
			So(algo, ShouldEqual, i)
			So(algo.String(), ShouldEqual, SieveAlgo(i).String())
		}

		So(EratosthenesAlgo, ShouldEqual, 0)
		So(SundaramAlgo, ShouldEqual, 1)
		So(AtkinAlgo, ShouldEqual, 2)

		So(EratosthenesAlgo.String(), ShouldEqual, "EratosthenesAlgo")
		So(SundaramAlgo.String(), ShouldEqual, "SundaramAlgo")
		So(AtkinAlgo.String(), ShouldEqual, "AtkinAlgo")
	})
}
