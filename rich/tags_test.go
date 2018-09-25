package rich

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTags(t *testing.T) {
	Convey("Tags", t, func() {
		Convey("Add", func() {
			t := Tags{}
			t.Add("测试3")
			So(len(t), ShouldEqual, 1)
			t.Add("测试3")
			So(len(t), ShouldEqual, 1)
			t.Add("测试2")
			So(len(t), ShouldEqual, 2)
			So(t[0], ShouldEqual, "测试2")
		})
		Convey("Sort", func() {
			t := Tags{"3", "1", "2"}
			t.Sort()
			So(t[0], ShouldEqual, "1")
		})
		Convey("DelPy", func() {
			t := Tags{"11", "1", "22", "2"}
			t.DelPy()
			So(len(t), ShouldEqual, 2)
		})
	})
}
