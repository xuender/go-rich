package rich

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestItem(t *testing.T) {
	Convey("NewItem", t, func() {
		Convey("Item", func() {
			item := NewItem()
			item.Name = "BOOK"
			So(len(item.Id), ShouldEqual, 18)
		})
	})
}
