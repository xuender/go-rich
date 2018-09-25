package rich

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUtils(t *testing.T) {
	Convey("pinyin", t, func() {
		Convey("py", func() {
			So(py("中国"), ShouldEqual, "zhong guo")
			So(py("中国WTo"), ShouldEqual, "zhong guo wto")
			So(py("11"), ShouldEqual, "11")
		})
	})
	Convey("Axis", t, func() {
		Convey("坐标", func() {
			So(Axis(0, 0), ShouldEqual, "A1")
			So(Axis(1, 2), ShouldEqual, "B3")
			So(Axis(3, 2), ShouldEqual, "D3")
		})
	})
}
