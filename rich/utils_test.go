package rich

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUtils(t *testing.T) {
	Convey("pinyin", t, func() {
		Convey("py", func() {
			So(py("中国"), ShouldEqual, "zhong guo")
			So(py("中国WTo"), ShouldEqual, "zhong guo wto zhong guo")
			So(py("重瞳"), ShouldContainSubstring, "zhong tong")
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
	Convey("PC", t, func() {
		Convey("排列组合", func() {
			a := [][]string{
				{"A", "a"},
				{"b"},
				{"c", "C"},
			}
			So(PC(a), ShouldContain, "A b c")
		})
	})
}
