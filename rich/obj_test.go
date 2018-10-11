package rich

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestObj(t *testing.T) {
	Convey("Obj", t, func() {
		Convey("Delete", func() {
			c := Customer{}
			So(c.IsDelete(), ShouldEqual, false)
			c.Delete()
			So(c.IsDelete(), ShouldEqual, true)
		})
	})
}
