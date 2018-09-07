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
}
